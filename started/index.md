---
layout: page
title: Getting Started

order: 10

redirect_from:
- /tour/introduction/
- /tour/introduction/hello
- /tour/introduction/learning
- /tour/introduction/samples

install:
    tab1:
        name: windows
        title: Windows
        source: windows
        active: 1
    tab2:
        name: macos
        title: macOS X
        source: macos
    tab3:
        name: linux
        title: Linux
        source: linux
    tab4:
        name: rpi
        title: Raspberry Pi
        source: rpi
    tab5:
        name: bsd
        title: BSD
        source: bsd
    tab6:
        name: android
        title: Android
        source: android
    tab7:
        name: ios
        title: iOS
        source: ios
    tab8:
        name: termux
        title: Termux (create Android apk without PC - on Android)
        source: termux

---

Using the Fyne toolkit to build cross platform applications is very simple but does require some tools to be installed before you can begin. If your computer is set up for development with Go then the following steps may not be required, but we advise reading the tips for your operating system just in case. If later steps in this tutorial fail then you should re-visit the prerequisites below.

## Prerequisites

Fyne requires 3 basic elements to be present, the Go tools (at least version 1.12), a C compiler (to connect with system graphics drivers) and a system graphics driver. The instructions vary depending on your operating system, choose the appropriate tab below for installation instructions.

Note that these steps are just required for development - your Fyne applications will not require any setup or dependency installation for end users!

{% include tabs.html bodyclass="fullborder" tabs=page.install id="install" %}

<div id="install__windows" class="hidden">
<div style="text-align: left" markdown="1">

1. Download Go from the [download page](https://golang.org/dl/) and follow instructions
2. Install one of the available C compilers for windows, the following are tested with Go and Fyne:
    * MSYS2 with MingW-w64 - [msys2.org](https://www.msys2.org/)
    * TDM-GCC - [tdm-gcc.tdragon.net](https://jmeubank.github.io/tdm-gcc/download/)
    * Cygwin - [cygwin.com](https://www.cygwin.com/)
3. In Windows your graphics driver will already be installed, but it is recommended to ensure they are up to date.

The steps for installing with MSYS2 (recommended) are as follows:
* Install MSYS2 from [msys2.org](https://www.msys2.org/)
* Once installed do not use the MSYS terminal that opens
* Open "MSYS2 MinGW 64-bit" from the start menu
* Execute the following commands (if asked for install options be sure to choose "all"):

        $ pacman -Syu
        $ pacman -S git mingw-w64-x86_64-toolchain

* You will need to add /c/Program\ Files/Go/bin and ~/Go/bin to your $PATH, for MSYS2 you can paste the following command into your terminal:

        $ echo "export PATH=\$PATH:/c/Program\ Files/Go/bin:~/Go/bin" >> ~/.bashrc

* For the compiler to work on other terminals you will need to set up the windows %PATH% variable to find these tools. Go to the "Edit the system environment variables" control panel, tap "Advanced" and add "C:\msys64\mingw64\bin" to the Path list.

</div>
</div>

<div id="install__macos" class="hidden">
<div style="text-align: left" markdown="1">

1. Download Go from the [download page](https://golang.org/dl/) and follow instructions
2. Install Xcode from the [Mac App Store](https://apps.apple.com/us/app/xcode/id497799835?mt=12)
3. Set up the Xcode command line tools by opening a Terminal window and typing the following:
    `xcode-select --install`
4. In macOS the graphics drivers will already be installed.

</div>
</div>

<div id="install__linux" class="hidden">
<div style="text-align: left" markdown="1">

* You will need to install Go, gcc and the graphics library header files using your package manager, one of the following commands will probably work.
* **Debian / Ubuntu:**
`sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev`
* **Fedora:**
`sudo dnf install golang golang-misc gcc libXcursor-devel libXrandr-devel mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel`
* **Arch Linux:**
`sudo pacman -S go xorg-server-devel libxcursor libxrandr libxinerama libxi`
* **Solus:**
`sudo eopkg it -c system.devel golang mesalib-devel libxrandr-devel libxcursor-devel libxi-devel libxinerama-devel`
* **openSUSE:**
`sudo zypper install go gcc libXcursor-devel libXrandr-devel Mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel`
* **Void Linux:**
`sudo xbps-install -S go base-devel xorg-server-devel libXrandr-devel libXcursor-devel libXinerama-devel libXxf86vm-devel`
* **Alpine Linux**
`sudo apk add go gcc libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev linux-headers mesa-dev`
* **NixOS**
`nix-shell -p libGL pkg-config xorg.libX11.dev xorg.libXcursor xorg.libXi xorg.libXinerama xorg.libXrandr xorg.libXxf86vm`

</div>
</div>

<div id="install__rpi" class="hidden">
<div style="text-align: left" markdown="1">

* You will need to install Go, gcc and the graphics library header files using the package manager.
* `sudo apt-get install golang gcc libegl1-mesa-dev xorg-dev`

</div>
</div>

<div id="install__bsd" class="hidden">
<div style="text-align: left" markdown="1">

* You will need to install Go, gcc and the graphics library header files using the package manager.
* **FreeBSD:**
`sudo pkg install go gcc xorg pkgconf`
* **OpenBSD:**
`sudo pkg_add go`
* **NetBSD:**
`sudo pkgin install go pkgconf`
</div>
</div>

<div id="install__android" class="hidden">
<div style="text-align: left" markdown="1">

* To develop apps for Android you will first need to install the tools for your current computer (Windows, macOS or Linux)
* Once complete you will need to install the Android SDK and Android NDK - the recommended approach is to install [Android Studio](https://developer.android.com/studio/index.html) and then go to **Tools > SDK Manager** and from **SDK Tools** install the **NDK (Side by side)** package.
* Alternatively you can download the [Standalone Android NDK](https://github.com/android/ndk/wiki#supported-downloads) which is a more lean approach. Extract the folder and point the `ANDROID_NDK_HOME` environment variable to it.

</div>
</div>

<div id="install__ios" class="hidden">
<div style="text-align: left" markdown="1">

* To develop apps for iOS you will need access to an Apple Mac computer, configured according to the **macOS** tab above.
* You will also need to create an [Apple Developer account](https://developer.apple.com) and sign up to the developer program (costs apply) to obtain the necessary certificate to run your app on any devices.

</div>
</div>


<div id="install__termux" class="hidden">
<div style="text-align: left" markdown="1">

Compiling Fyne apps on Android you will need Android 9 or above

* Install [fdroid](https://f-droid.org/) and from there [termux](https://f-droid.org/packages/com.termux/)
* Open Termux and install Go and Git:
`pkg install golang git`
* Install NDK and SDK to termux from [https://github.com/Lzhiyong/termux-ndk](https://github.com/Lzhiyong/termux-ndk) and set enviroment variables `ANDROID_HOME` and `ANDROID_NDK_HOME`

(you can get more help from  [video](https://youtu.be/uGtVjf4_Ivo))


</div>
</div>
<script type="text/javascript">

    function clickTab(tab) {
        document.querySelector('li.tabcontrol[data-name="'+tab+'"]').click();
    }

    $(document).ready(function(){
        var ua = navigator.userAgent || navigator.vendor || window.opera;
        if (/android/i.test(ua)) {
            clickTab("android");
            return;
        } else if (/iPad|iPhone|iPod/.test(ua) && !window.MSStream) {
            clickTab("ios");
            return;
        }

        var os = window.navigator.platform;
        if (os == "win32") {
            clickTab("windows");
        } else if (os == "MacIntel") {
            clickTab("macos");
        } else if (os == "Linux i686" || os == "Linux x86_64") {
            clickTab("linux");
        } else if (os == "Linux armv7l") {
            clickTab("rpi");
        } else if (os == "FreeBSD i386" || os == "FreeBSD amd64" || os == "OpenBSD i386" || os == "OpenBSD amd64" || os == "NetBSD i386" || os == "NetBSD amd64") {
        	clickTab("bsd");
        }
    });
</script> 

## Downloading

Since Go 1.16 you will need to set up the module before you can use the package.

Run the following command and replace `MODULE_NAME` with your preferred module name (this should be called in a new folder specific for your application).

    $ cd myapp
    $ go mod init MODULE_NAME

You now need to download the Fyne module and helper tool. This will be done using the following commands: 

    $ go get fyne.io/fyne/v2@latest
    $ go install fyne.io/fyne/v2/cmd/fyne@latest

If you are unsure of how Go modules work, consider reading [Tutorial: Create a Go module](https://golang.org/doc/tutorial/create-module).

## Check your installation

Before coding an app or running an example you can check your install using the
[Fyne Setup](https://geoffrey-artefacts.fynelabs.com/github/andydotxyz/fyne-io/setup/latest/) tool.
Simply download the right app for your computer from the link and run it,
you should see something like the following screen:

<p style="text-align: center; margin: auto">
<img src="/images/started/setup.png" style="width: 320px" />
</p>

If there are any problems with your installation see the [troubleshooting](/faq/troubleshoot) section for hints.

## Run the demo

If you want to see the Fyne toolkit in action before you start to code your own application,
you can see our [demo app](https://github.com/fyne-io/fyne/tree/master/cmd/fyne_demo) running on your computer by executing:

    $ go run fyne.io/fyne/v2/cmd/fyne_demo@latest

Please note that the first run has to compile some C-code and can thus take longer than usual. Subsequent builds reuse the cache and will be much faster.

### Older Go version

To run the demo on an older version of Go, simply execute the follwing command instead:

    $ go run fyne.io/fyne/v2/cmd/fyne_demo

### Installing

If you want to, you can also install the demo using the following command (requires Go 1.16 or later):

    $ go install fyne.io/fyne/v2/cmd/fyne_demo@latest

For earlier versions of Go, you need to use the following command instead:

    $ go get fyne.io/fyne/v2/cmd/fyne_demo

If your `GOBIN` environment has been added to path (should be by default on macOS and Windows), you can then run the demo:

    $ fyne_demo

And that's all there is to it! Now you can write your own Fyne application in your IDE of choice. If you want to see some Fyne code in action then you can read [your first application](/started/firstapp).
