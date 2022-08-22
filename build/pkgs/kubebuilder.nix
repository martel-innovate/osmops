#
# Fix kubebuilder package. It's broken in the nixkpkgs version we're using.
# The problem is just a silly source hash mismatch:
#
# trying https://github.com/kubernetes-sigs/kubebuilder/archive/v3.1.0.tar.gz
#   % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
#                                  Dload  Upload   Total   Spent    Left  Speed
# 100   135  100   135    0     0    602      0 --:--:-- --:--:-- --:--:--   602
# 100 1816k    0 1816k    0     0  2959k      0 --:--:-- --:--:-- --:--:-- 2959k
# unpacking source archive /private/tmp/nix-build-source.drv-0/v3.1.0.tar.gz
# hash mismatch in fixed-output derivation '/nix/store/l2mvac03b398x7jnhbqdf9051k4rsini-source':
#   wanted: sha256:1726j2b5jyvllvnk60g6px3g2jyyphd9pc4vgid45mis9b60sh8a
#   got:    sha256:0bl5ff2cplal6hg75800crhyviamk1ws85sq60h4zg21hzf21y68
# cannot build derivation '/nix/store/if764s9fl71ihg60sifgr2a9ffp8qb24-kubebuilder-3.1.0.drv': 1 dependencies couldn't be built
# error: build of '/nix/store/if764s9fl71ihg60sifgr2a9ffp8qb24-kubebuilder-3.1.0.drv' failed
#
# I tried fixing it with an overlay, but that didn't work:
#
#   fixBrokenPkgsOverlay = self: super: {
#     kubebuilder = super.kubebuilder.overrideAttrs (oldAttrs: rec {
#       version = "3.1.0";
#       src = super.fetchFromGitHub {
#         owner = "kubernetes-sigs";
#         repo = "kubebuilder";
#         rev = "v${version}";
#         sha256 = "0bl5ff2cplal6hg75800crhyviamk1ws85sq60h4zg21hzf21y68";
#       };
#     });
#   };
#
# I think the problem could be that `overrideAttrs` works with `mkDerivation`,
# but the kubebuilder package uses `buildGoModule`? In fact, with the above
# overlay I get the exact same error as if the `src` attribute hasn't been
# overridden.
#
{ lib
, buildGoModule
, fetchFromGitHub
, installShellFiles
, makeWrapper
, git
, go
}:

buildGoModule rec {
  pname = "kubebuilder";
  version = "3.1.0";

  src = fetchFromGitHub {
    owner = "kubernetes-sigs";
    repo = "kubebuilder";
    rev = "v${version}";
    sha256 = "0bl5ff2cplal6hg75800crhyviamk1ws85sq60h4zg21hzf21y68";
  };
  vendorSha256 = "0zxyd950ksjswja64rfri5v2yaalfg6qmq8215ildgrcavl9974n";

  subPackages = ["cmd" "pkg/..."];

  preBuild = ''
    export buildFlagsArray+=("-ldflags=-X main.kubeBuilderVersion=v${version} \
        -X main.goos=$GOOS \
        -X main.goarch=$GOARCH \
        -X main.gitCommit=v${version} \
        -X main.buildDate=v${version}")
  '';

  doCheck = true;

  postInstall = ''
    mv $out/bin/cmd $out/bin/kubebuilder
    wrapProgram $out/bin/kubebuilder \
      --prefix PATH : ${lib.makeBinPath [ go ]}
  '';

  allowGoReference = true;
  nativeBuildInputs = [ makeWrapper git ];

  meta = with lib; {
    homepage = "https://github.com/kubernetes-sigs/kubebuilder";
    description = "SDK for building Kubernetes APIs using CRDs";
    license = licenses.asl20;
    maintainers = with maintainers; [ cmars ];
  };
}
