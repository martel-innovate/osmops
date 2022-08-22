#
# Version of the Nix infrastructure our project uses.
# We pin absolutely everything to make sure no matter what everybody
# gets the exact same build/dev environment, Docker images, etc.
# Reproducibility is king. (TODO need Flakes?)
#
{
  # Arguments to `fetchFromGitHub`, detailing the Nixpkgs source we
  # want to use.
  nixpkgsGitHub = {
    owner = "NixOS";
    repo = "nixpkgs";
    rev = "f6ccdfcd2ac4f2e259d20e378737dcbd0ca7debe";                 # (1)
    sha256 = "1d2lk7a0l166pvgy0xfdlhxgja986hgn39szn9d1fqamyhxzvbaz";  # (2)
  };

}
# NOTE
# 1. Nixpkgs commit.
# Git hash of the Nixpkgs commit to fetch. We'll pin our infrastructure
# to the Nix definitions as they were at that commit. Normally it should
# be the latest commit known to work with our project. To get the ID of
# the latest commit on the `nixpkgs-unstable` branch, run
#
#   $ git ls-remote https://github.com/nixos/nixpkgs nixpkgs-unstable
#
# A quick way to get a description of a commit is to use the GitHub API
# to GET a JSON object describing the commit associated to the commit
# hash in the URL---short hashes work too. E.g.
#
#     https://api.github.com/repos/nixos/nixpkgs/commits/5dbf5f9
#
# 2. Nixpkgs commit SHA256.
# To figure out the SHA256 of the commit in (1), you could initially
# set it to a made-up one and just let the Nix build bomb out, it'll
# tell you what's the actual SHA256 to use. Or you could run e.g.
#
#   $ nix run -f '<nixpkgs>' nix-prefetch-github -c nix-prefetch-github \
#             --rev 5dbf5f90d97c0af9efd36ecfdb8648e74ce39532 NixOS nixpkgs
#
