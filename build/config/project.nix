#
# Project build info.
#
rec
{
  # Absolute path to the repo's root dir.
  root = ../../.;

  # The name of this project. Taken to be the name of the repo dir; sort of
  # customary for online repos, but change it if you don't like it :-)
  # Project derivations like local Haskell packages get added to the
  # Nix packages in a set having this name so you can reference them
  # easily e.g. `pkgs.my-project.haskell.my-pkg-2`.
  # Have a look at `pkgset.nix` to see what winds up in `pkgs.my-project`.
  name = baseNameOf (toString root);

  # Absolute path to the directory containing the local source packages
  # implementing the project's components.
  componentsDir = root + "/components";
}
