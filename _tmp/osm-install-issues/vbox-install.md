```bash

Last login: Thu Sep  2 12:45:53 on ttys005
mactel:source-watcher andrea$ brew cask uninstall virtualbox
Error: Unknown command: cask
mactel:source-watcher andrea$ brew cask --help
Error: Unknown command: cask
mactel:source-watcher andrea$ brew --help
Example usage:
  brew search [TEXT|/REGEX/]
  brew info [FORMULA...]
  brew install FORMULA...
  brew update
  brew upgrade [FORMULA...]
  brew uninstall FORMULA...
  brew list [FORMULA...]

Troubleshooting:
  brew config
  brew doctor
  brew install --verbose --debug FORMULA

Contributing:
  brew create [URL [--no-fetch]]
  brew edit [FORMULA...]

Further help:
  brew commands
  brew help [COMMAND]
  man brew
  https://docs.brew.sh
mactel:source-watcher andrea$ brew cask
Error: Unknown command: cask
mactel:source-watcher andrea$ brew uninstall virtualbox
==> Uninstalling Cask virtualbox
==> Running uninstall script VirtualBox_Uninstall.tool
Password:

Welcome to the VirtualBox uninstaller script.

Executing: /usr/bin/kmutil showloaded --list-only --bundle-identifier org.virtualbox.kext.VBoxUSB
No variant specified, falling back to release
Executing: /usr/bin/kmutil showloaded --list-only --bundle-identifier org.virtualbox.kext.VBoxNetFlt
No variant specified, falling back to release
Executing: /usr/bin/kmutil showloaded --list-only --bundle-identifier org.virtualbox.kext.VBoxNetAdp
No variant specified, falling back to release
Executing: /usr/bin/kmutil showloaded --list-only --bundle-identifier org.virtualbox.kext.VBoxDrv
No variant specified, falling back to release
The following files and directories (bundles) will be removed:
    /Users/andrea/Library/LaunchAgents/org.virtualbox.vboxwebsrv.plist
    /usr/local/bin/VirtualBox
    /usr/local/bin/VBoxManage
    /usr/local/bin/VBoxVRDP
    /usr/local/bin/VBoxHeadless
    /usr/local/bin/vboxwebsrv
    /usr/local/bin/VBoxBugReport
    /usr/local/bin/VBoxBalloonCtrl
    /usr/local/bin/VBoxAutostart
    /usr/local/bin/VBoxDTrace
    /usr/local/bin/vbox-img
    /Library/LaunchDaemons/org.virtualbox.startup.plist
    /Library/Python/2.7/site-packages/vboxapi/VirtualBox_constants.py
    /Library/Python/2.7/site-packages/vboxapi/VirtualBox_constants.pyc
    /Library/Python/2.7/site-packages/vboxapi/__init__.py
    /Library/Python/2.7/site-packages/vboxapi/__init__.pyc
    /Library/Python/2.7/site-packages/vboxapi-1.0-py2.7.egg-info
    /Library/Application Support/VirtualBox/LaunchDaemons/
    /Library/Application Support/VirtualBox/VBoxDrv.kext/
    /Library/Application Support/VirtualBox/VBoxUSB.kext/
    /Library/Application Support/VirtualBox/VBoxNetFlt.kext/
    /Library/Application Support/VirtualBox/VBoxNetAdp.kext/
    /Applications/VirtualBox.app/
    /Library/Python/2.7/site-packages/vboxapi/

And the traces of following packages will be removed:
    org.virtualbox.pkg.vboxkexts
    org.virtualbox.pkg.virtualbox
    org.virtualbox.pkg.virtualboxcli

The uninstallation processes requires administrative privileges
because some of the installed files cannot be removed by a normal
user. You may be prompted for your password now...

Successfully unloaded VirtualBox kernel extensions.
Forgot package 'org.virtualbox.pkg.vboxkexts' on '/'.
Forgot package 'org.virtualbox.pkg.virtualbox' on '/'.
Forgot package 'org.virtualbox.pkg.virtualboxcli' on '/'.
Done.
==> Uninstalling packages:
==> Purging files for version 6.0.0,127566 of Cask virtualbox
mactel:source-watcher andrea$ brew install virtualbox
Error:
  homebrew-core is a shallow clone.
  homebrew-cask is a shallow clone.
To `brew update`, first run:
  git -C /usr/local/Homebrew/Library/Taps/homebrew/homebrew-core fetch --unshallow
  git -C /usr/local/Homebrew/Library/Taps/homebrew/homebrew-cask fetch --unshallow
These commands may take a few minutes to run due to the large size of the repositories.
This restriction has been made on GitHub's request because updating shallow
clones is an extremely expensive operation due to the tree layout and traffic of
Homebrew/homebrew-core and Homebrew/homebrew-cask. We don't do this for you
automatically to avoid repeatedly performing an expensive unshallow operation in
CI systems (which should instead be fixed to not use shallow clones). Sorry for
the inconvenience!
==> Caveats
virtualbox requires a kernel extension to work.
If the installation fails, retry after you enable it in:
  System Preferences → Security & Privacy → General

For more information, refer to vendor documentation or this Apple Technical Note:
  https://developer.apple.com/library/content/technotes/tn2459/_index.html

==> Downloading https://download.virtualbox.org/virtualbox/6.1.18/VirtualBox-6.1.18-142142-OSX.dmg
######################################################################## 100.0%
==> Installing Cask virtualbox
==> Running installer for virtualbox; your password may be necessary.
Package installers may write to any location; options such as `--appdir` are ignored.
installer: Package name is Oracle VM VirtualBox
installer: choices changes file '/var/folders/tt/prkpxkn1001cg362nwffx4kc0000gn/T/choices20210902-67677-lnqbru.xml' applied
installer: Upgrading at base path /
installer: The upgrade was successful.
==> Changing ownership of paths required by virtualbox; your password may be necessary.
🍺  virtualbox was successfully installed!
mactel:source-watcher andrea$
  [Restored 2 Sep 2021 at 19:32:53]
Last login: Thu Sep  2 19:32:53 on ttys008
mactel:source-watcher andrea$ brew install virtualbox virtualbox-extension-pack
Error:
  homebrew-core is a shallow clone.
  homebrew-cask is a shallow clone.
To `brew update`, first run:
  git -C /usr/local/Homebrew/Library/Taps/homebrew/homebrew-core fetch --unshallow
  git -C /usr/local/Homebrew/Library/Taps/homebrew/homebrew-cask fetch --unshallow
These commands may take a few minutes to run due to the large size of the repositories.
This restriction has been made on GitHub's request because updating shallow
clones is an extremely expensive operation due to the tree layout and traffic of
Homebrew/homebrew-core and Homebrew/homebrew-cask. We don't do this for you
automatically to avoid repeatedly performing an expensive unshallow operation in
CI systems (which should instead be fixed to not use shallow clones). Sorry for
the inconvenience!
Warning: Cask 'virtualbox' is already installed.

To re-install virtualbox, run:
  brew reinstall virtualbox
Warning: Cask 'virtualbox-extension-pack' is already installed.

To re-install virtualbox-extension-pack, run:
  brew reinstall virtualbox-extension-pack
mactel:source-watcher andrea$ brew uninstall virtualbox-extension-pack
==> Uninstalling Cask virtualbox-extension-pack
Password:
0%...10%...20%...30%...40%...50%...60%...70%...80%...90%...100%
==> Purging files for version 6.0.0 of Cask virtualbox-extension-pack
mactel:source-watcher andrea$ brew install virtualbox-extension-pack
Error:
  homebrew-core is a shallow clone.
  homebrew-cask is a shallow clone.
To `brew update`, first run:
  git -C /usr/local/Homebrew/Library/Taps/homebrew/homebrew-core fetch --unshallow
  git -C /usr/local/Homebrew/Library/Taps/homebrew/homebrew-cask fetch --unshallow
These commands may take a few minutes to run due to the large size of the repositories.
This restriction has been made on GitHub's request because updating shallow
clones is an extremely expensive operation due to the tree layout and traffic of
Homebrew/homebrew-core and Homebrew/homebrew-cask. We don't do this for you
automatically to avoid repeatedly performing an expensive unshallow operation in
CI systems (which should instead be fixed to not use shallow clones). Sorry for
the inconvenience!
==> Caveats
Installing virtualbox-extension-pack means you have AGREED to the license at:
  https://www.virtualbox.org/wiki/VirtualBox_PUEL

==> Downloading https://download.virtualbox.org/virtualbox/6.1.18/Oracle_VM_VirtualBox_Extension_Pack-6.1.18.vbox-extpack
######################################################################## 100.0%
All formula dependencies satisfied.
==> Installing Cask virtualbox-extension-pack
0%...10%...20%...30%...40%...50%...60%...70%...80%...90%...100%
🍺  virtualbox-extension-pack was successfully installed!
mactel:source-watcher andrea$

```