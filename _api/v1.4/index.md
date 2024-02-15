---
layout: page
tags: [api]
title: Fyne API "fyne"
---

# fyne
---
```go
import "fyne.io/fyne"
```

Package fyne describes the objects and components available to any Fyne app. These can all be created, manipulated and tested without rendering (for speed). Your main package should use the app package to create an application with a default driver that will render your UI.

A simple application may look like this:

```go
    package main

    import "fyne.io/fyne/app"
    import "fyne.io/fyne/widget"

    func main() {
    	a := app.New()
    	w := a.NewWindow("Hello")

    	hello := widget.NewLabel("Hello Fyne!")
    	w.SetContent(widget.NewVBox(
    		hello,
    		widget.NewButton("Hi!", func() {
    			hello.SetText("Welcome :)")
    		}),
    	))

    	w.ShowAndRun()
    }
```

## Usage

```go
const SettingsScaleAuto = float32(-1.0)
```
SettingsScaleAuto is a specific scale value that indicates a canvas should scale according to the DPI of the window that contains it.


<div class="deprecated">
Deprecated: Automatic scaling is now handled in the drivers and is not a user setting.</div>

#### func  IsHorizontal

```go
func IsHorizontal(orient DeviceOrientation) bool
```
IsHorizontal is a helper utility that determines if a passed orientation is horizontal

#### func  IsVertical

```go
func IsVertical(orient DeviceOrientation) bool
```
IsVertical is a helper utility that determines if a passed orientation is vertical

#### func  LogError

```go
func LogError(reason string, err error)
```
LogError reports an error to the command line with the specified err cause, if not nil. The function also reports basic information about the code location.

#### func  Max

```go
func Max(x, y int) int
```
Max returns the larger of the passed values.

#### func  Min

```go
func Min(x, y int) int
```
Min returns the smaller of the passed values.

#### func  SetCurrentApp

```go
func SetCurrentApp(current App)
```
SetCurrentApp is an internal function to set the app instance currently running.

#### types

 * [App](app.html)
 * [BuildType](buildtype.html)
 * [Canvas](canvas.html)
 * [CanvasObject](canvasobject.html)
 * [Clipboard](clipboard.html)
 * [Container](container.html)
 * [Device](device.html)
 * [DeviceOrientation](deviceorientation.html)
 * [Disableable](disableable.html)
 * [DoubleTappable](doubletappable.html)
 * [DragEvent](dragevent.html)
 * [Draggable](draggable.html)
 * [Driver](driver.html)
 * [Focusable](focusable.html)
 * [KeyEvent](keyevent.html)
 * [KeyName](keyname.html)
 * [Layout](layout.html)
 * [ListableURI](listableuri.html)
 * [MainMenu](mainmenu.html)
 * [Menu](menu.html)
 * [MenuItem](menuitem.html)
 * [Notification](notification.html)
 * [OverlayStack](overlaystack.html)
 * [PointEvent](pointevent.html)
 * [Position](position.html)
 * [Preferences](preferences.html)
 * [Resource](resource.html)
 * [ScrollEvent](scrollevent.html)
 * [Scrollable](scrollable.html)
 * [SecondaryTappable](secondarytappable.html)
 * [Settings](settings.html)
 * [Shortcut](shortcut.html)
 * [ShortcutCopy](shortcutcopy.html)
 * [ShortcutCut](shortcutcut.html)
 * [ShortcutHandler](shortcuthandler.html)
 * [ShortcutPaste](shortcutpaste.html)
 * [ShortcutSelectAll](shortcutselectall.html)
 * [Shortcutable](shortcutable.html)
 * [Size](size.html)
 * [StaticResource](staticresource.html)
 * [Storage](storage.html)
 * [StringValidator](stringvalidator.html)
 * [Tappable](tappable.html)
 * [TextAlign](textalign.html)
 * [TextStyle](textstyle.html)
 * [TextWrap](textwrap.html)
 * [Theme](theme.html)
 * [URI](uri.html)
 * [URIReadCloser](urireadcloser.html)
 * [URIWriteCloser](uriwritecloser.html)
 * [Validatable](validatable.html)
 * [Widget](widget.html)
 * [WidgetRenderer](widgetrenderer.html)
 * [Window](window.html)
