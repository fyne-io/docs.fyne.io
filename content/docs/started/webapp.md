---
layout: page
title: Run in a Browser
slug: webapp

weight: 1110
---

Fyne applications can also be run over the web using a standard web browser!
A web app created with Fyne will bundle a WebAssembly binary which will run in most modern browsers

To prepare your app to be used over the web we use the "fyne" cli app again, which has a
"serve" command for quick testing

```
go install fyne.io/tools/cmd/fyne@latest
fyne serve
```

You will see, after a few moments, that a web server has been started on port :8080.
Just open your web browser to https://localhost:8080 and you can use your app!

### Packaging for web distribution

The `fyne serve` command is great for local testing, but just like other platforms you'll want
to be able to distribute your app as well. To prepare the files for upload just use the
`fyne package` command like with regular [Packaging](/started/packaging).

```
fyne package -os web
```

### Demo

You can see a Fyne app in action to test on any of your devices by visiting [demo.fyne.io](https://demo.fyne.io/).

### Limitations

As of release v2.5.0 the web driver is fully supported but is not yet 100% complete, so your app may not be able to use
the following features:

* File open/save dialog
* Storage of Documents

These issues are being worked on and will be resolved in a future release.
