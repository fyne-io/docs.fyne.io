---
layout: page
title: Compile Options

redirect_from:
- /started/compiling
---

## Build tags

Fyne will typically configure your application appropriately for the target platform by selecting the driver and configuration. The following build tags are supported and can help in your development. For example if you wish to simulate a mobile application whilst running on a desktop computer you could use the following command:

	go run -tags mobile main.go


| Tag      | Description               |
|----------|---------------------------|
| `debug`  | Show debug information, including visual layout to help understand your app. |
| `gles`   | Force use of embedded OpenGL (GLES) instead of full OpenGL. This is normally controlled by the target device and not normally needed. |
| `hints`  | Display developer hints for improvements or optimisations. Running with `hints` will log when your application does not follow material design or other recommendations. |
| `mobile` | This tag runs an application in a simulated mobile window. Useful when you want to preview your app on a mobile platform without compiling and installing to the device. |
| `no_animations` | Disable the non-essential animations in standard widgets and containers. |
| `no_emoji` | Don't include the embedded emoji font. This will disable emoji in  your app but will make the binary smaller. |
| `no_metadata` | Disable runtime lookup of metadata even when not packaging in release mode. |
| `no_native_menus` | This flag is specifically for macOS and indicates that the application should not use the macOS native menus. Instead menus will be displayed inside the application window. Most useful for testing an application on macOS to simulate the behavior on Windows or Linux. |
