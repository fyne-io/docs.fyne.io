---
layout: page
title: Distributing to App Stores	
slug: distribution

weight: 1130
---

Packaging a graphical app as described in the [Packaging](/started/packaging)
page provides a file or bundle that could be directly shared or distributed.
However signing and uploading to app stores and market places is an additional
step that requires platform-specific configuration, which we will cover in this
page.

In each of these steps we will use a new tool that is part of the fyne command
line utilities. The `fyne release` step handles the signing and preparation
for each store, but the parameters vary per-platform, which we see below.

### macOS App Store (since fyne 1.4.2)

Prerequisites:

* Apple mac running macOS and Xcode
* Apple Developer account
* Mac App Store application certificate
* Mac App Store installer certificate
* Apple Transporter app from App Store

1. 
Set up your app / version ready for a build to be uploaded at
[AppStore Connect](https://appstoreconnect.apple.com).

2.
Bundle the completed app for release:

```
$ fyne release -appID com.example.myapp -appVersion 1.0 -appBuild 1 -category games
```

3.
Drag the `.pkg` onto Transporter and tap "Deliver".

4.
Go to back to the AppStore Connect website, choose your build for the release and submit for review.

### Google Play Store (Android)

Prerequisites:

* Google Play Console account
* distribution keystore (creation instructions in
[android docs](https://developer.android.com/studio/publish/app-signing))

1.
Set up your app / version ready for build to be uploaded at
[Google Play Console](https://play.google.com/apps/publish). Turn off "Play app signing" option as we manage it ourselves.

2.
Bundle the completed app for release:

```
$ fyne release -os android -appID com.example.myapp -appVersion 1.0 -appBuild 1
```

3.
Drag the `.apk` file into the build drop zone on the app version page in Play Console

4.
Start rollout of new version.

### iOS App Store (since fyne 1.4.1)

Prerequisites:

* Apple mac running macOS and Xcode
* Apple Developer account
* iOS App Store distribution certificate
* Apple Transporter app from App Store

1.
Set up your app / version ready for a build to be uploaded at
[AppStore Connect](https://appstoreconnect.apple.com).

2.
Bundle the completed app for release:

```
$ fyne release -os ios -appID com.example.myapp -appVersion 1.0 -appBuild 1
```

3.
Drag the `.ipa` onto Transporter and tap "Deliver".

4.
Go to back to the AppStore Connect website, choose your build for the release and submit for review.
