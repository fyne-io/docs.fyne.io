---
layout: page
title: App Metadata
slug: metadata

weight: 1120
---

Since release v2.1.0 of the `fyne` command we support a metadata file that allows you to store
information about your application in the repository.
This file is optional, but can help to avoid having to remember specific build parameters for
each package and release command.

## Basic configuration

The file should be named `FyneApp.toml` in the directory where you run the `fyne` command
(this is normally the `main` package). The contents of the file are as follows:

```toml
Website = "https://example.com"

[Details]
Icon = "Icon.png"
Name = "My App"
ID = "com.example.app"
Version = "1.0.0"
Build = 1
```

The top portion of the file is metadata that will be used if you upload
your app to the [Apps](https://apps.fyne.io) listing page, so it is optional.

The `Details` table contains data about your application that are used
in the release process by other app stores and operating systems.

The `fyne` tool will use this file if it is found, many mandatory command parameters are not required
if the metadata is present. You can still override these values by using command line parameters.

## Linux & BSD configuration

For Linux and BSD builds there is an optional table called `LinuxAndBSD`. This table contains additional parameters for a "desktop entry" configuration file of the Fyne app. All parameters are optional, but when present they will be used by the `fyne` tool (in addition to parameters from the `Details` table).

The contents of this section is as follows (with example data):

```toml
[LinuxAndBSD]
  GenericName = "Web Browser"
  Categories = ["Network"]
  Comment = "View sites on the Internet"
  Keywords = ["browser", "web"]
  ExecParams = "-x 42"
```

Hint: For instructions on how to define these parameters correctly, please see the [Desktop Entry Specification](https://specifications.freedesktop.org/desktop-entry-spec/desktop-entry-spec-latest.html) from freedesktop.org.
