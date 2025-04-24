---
layout: page
title: Packaging for Desktop

order: 50
---

Packaging a graphical app for distribution can be complex.
Graphical applications typically have icons and metadata associated
with them as well as specific formats required to integrate with each
environment. Windows executables need embedded icons, macOS apps are bundles and
with Linux there are various metadata files that should get installed. What a hassle!

Thankfully the "fyne" app has a "package" command that can handle this automatically. Just specifying the target OS and any required metadata (such as icon) will generate the appropriate package. The icon conversion will be done automatically for .icns or .ico so just provide a .png file :). All you need is to have the application already built for the target platform...

```
go install fyne.io/tools/cmd/fyne@latest
fyne package -os darwin -icon myapp.png
```
If you're using an older version of Go (<1.16), you should install fyne using `go get`

```
go get fyne.io/tools/cmd/fyne
fyne package -os darwin -icon myapp.png
```

Will create myapp.app, a complete bundle structure for distribution to macOS users. You could then build the linux and Windows versions too...

```
fyne package -os linux -icon myapp.png
fyne package -os windows -icon myapp.png
```

These commands will create:

  * myapp.tar.gz that contains a folder structure starting at usr/local/ that a Linux user could expand to the root of their system.
  * myapp.exe (after the second build, which is required for a windows package) will have the icon and app metadata embedded.

If you just want to install the desktop app on your computer then you can make
use of the helpful install subcommand. For example to install your current
application system wide you could simply execute the following:

```
fyne install -icon myapp.png
```

All of these commands also support a default icon file of `Icon.png` so that you
can avoid typing the parameter for each execution. Since Fyne 2.1 there is also a
[metadata file](/started/metadata) where you can set default options for your project.

Of course you can still run your applications using the standard Go
tools if you prefer.

{% include youtube.html id="zIV3Pv5QbeM" %}
