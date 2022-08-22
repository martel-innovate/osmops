#
# Adapted from https://github.com/c0c0n3/hasnix
#
# TODO: rather go with one of the project templates below?
# - https://github.com/nix-dot-dev/getting-started-nix-template
# - https://github.com/vlktomas/nix-examples
#

{
  pkgs ? import <nixpkgs> {}
}:

let
  inherit (pkgs) fetchFromGitHub;
  version = import ./config/version.nix;
  project = import ./config/project.nix;

in rec {

  nixpin = fetchFromGitHub version.nixpkgsGitHub;

  fixBrokenPkgsOverlay = self: super: {
    kubebuilder = super.callPackage ./pkgs/kubebuilder.nix { };
  };

  pinnedPkgs = import nixpin {
    overlays = [ fixBrokenPkgsOverlay ];
  };

  devTools = with pinnedPkgs; {
    # stuff listed in the source-watcher tute
    inherit go kubebuilder kind kubectl kustomize fluxcd;
    # TODO ideally we should include docker too...

    # VS code go extension deps
    inherit gopls delve gopkgs go-outline gomodifytags impl gotests;
    inherit go-tools; # = staticcheck
    # missing from nixpkgs: goplay; leaving this out

    # Only needed to connect to the Malaga demo cluster.
    inherit openvpn;
  };

  devShell = pinnedPkgs.mkShell {
    buildInputs = builtins.attrValues devTools;
  };

}
