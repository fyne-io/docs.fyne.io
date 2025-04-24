---
layout: page
title: Run Fyne Demo

---

If you want to see the Fyne toolkit in action before you start to code your own application,
you can see our [demo app](https://github.com/fyne-io/demo/tree/main).

### Running

If you want to, you can run the demo directly using the following command (requires Go 1.16 or later):

    go run github.com/fyne-io/demo@latest

For earlier versions of Go, you need to use the following command instead:

    go run github.com/fyne-io/demo

By browsing the different tabs of the app you can see all the features of Fyne toolkit.

### Installing

It is possible to install the app as a graphical application like all others on your computer.
We have the helpful `fyne` tool to do this for you.
First you will need to install the tool:

	go install fyne.io/tools/cmd/fyne@latest

After that you can simply package and install the demo app:

	fyne install github.com/fyne-io/demo

(This may require admin privileges on your system, and that in turn may 
require some path changes to include the `fyne` and `go` executables.)

After this step you can find "Fyne Demo" in your app launcher.

### Exploring the code

If you are interested in any of the features you should  check out the
[source code](https://github.com/fyne-io/demo/tree/main)
or join one of the [community channels](https://fyne.io#contact).
