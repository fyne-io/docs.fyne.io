---
layout: page
title: Mobile Packaging

redirect-from:
- /started/packaging-mobile
---

Your Fyne app code will work out of the box as mobile apps, just as it did for desktop.
However it is a little more complex to package the code for distribution.
This page will explore the process to do just that to get your app on iOS and Android.

Firstly you will need some more development tools installed for mobile packaging
to complete. For Android builds you must have the Android SDK and NDK
installed with appropriate environment set up so that the tools (such as `adb`)
can be found on the command line. To build iOS apps you will need Xcode installed
on your macOS computer as well as the command line tools optional package.

Once you have a working development environment the packaging is simple.
To build an application for Android and iOS the following commands will do
everything for you. Be sure to have a unique application identifier as it is
unwise to change these after your first release.

```
fyne package -os android -app-id com.example.myapp -icon mobileIcon.png
fyne package -os ios -app-id com.example.myapp -icon mobileIcon.png
```

After these commands have completed (which may take some time on first
compilation) you will see two new files in your directory, `myapp.apk` and
`myapp.app`. You will see that the latter has the same name as a darwin
application bundle - don't get them confused as they will not work on the
other platform.

To install the android app on your phone or a simulator simply call:

```
adb install myapp.apk
```

For iOS to install on device open Xcode and choose the "Devices and
Simulators" menu item within the "Window" menu. Then find your phone and drag
the `myapp.app` icon onto your app list. 

If you want to install on a simulator make sure to package your application using `iossimulator` vs `ios` 

```
fyne package -os iossimulator -app-id com.example.myapp -icon mobileIcon.png
```

Afterwards you can use the command line tools as follows:

```
xcrun simctl install booted myapp.app
```

