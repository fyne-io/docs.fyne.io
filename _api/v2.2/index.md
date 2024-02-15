---
layout: page
tags: [api]
title: Fyne API "fyne"
---

# fyne
---
```go
import "fyne.io/fyne/v2"
```

Package fyne describes the objects and components available to any Fyne app. These can all be created, manipulated and tested without rendering (for speed). Your main package should use the app package to create an application with a default driver that will render your UI.

A simple application may look like this:

```go
    package main

    import "fyne.io/fyne/v2/app"
    import "fyne.io/fyne/v2/container"
    import "fyne.io/fyne/v2/widget"

    func main() {
    	a := app.New()
    	w := a.NewWindow("Hello")

    	hello := widget.NewLabel("Hello Fyne!")
    	w.SetContent(container.NewVBox(
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
const AnimationRepeatForever = -1
```
AnimationRepeatForever is an AnimationCount value that indicates it should not stop looping.


<div class="since">Since: <code>
2.0</code></div>

```go
const KeyModifierShortcutDefault = KeyModifierSuper
```
KeyModifierShortcutDefault is the default key modifier for shortcuts (Control or Command).


<div class="since">Since: <code>
2.2</code></div>

```go
const KeyModifierShortcutDefault = KeyModifierControl
```
KeyModifierShortcutDefault is the default key modifier for shortcuts (Control or Command).


<div class="since">Since: <code>
2.2</code></div>

```go
var (
	// AnimationEaseInOut is the default easing, it starts slowly, accelerates to the middle and slows to the end.
	//
	// Since: 2.0
	AnimationEaseInOut = animationEaseInOut
	// AnimationEaseIn starts slowly and accelerates to the end.
	//
	// Since: 2.0
	AnimationEaseIn = animationEaseIn
	// AnimationEaseOut starts at speed and slows to the end.
	//
	// Since: 2.0
	AnimationEaseOut = animationEaseOut
	// AnimationLinear is a linear mapping for animations that progress uniformly through their duration.
	//
	// Since: 2.0
	AnimationLinear = animationLinear
)
```

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
func Max(x, y float32) float32
```
Max returns the larger of the passed values.

#### func  Min

```go
func Min(x, y float32) float32
```
Min returns the smaller of the passed values.

#### func  SetCurrentApp

```go
func SetCurrentApp(current App)
```
SetCurrentApp is an internal function to set the app instance currently running.

#### types

 * [Animation](animation.html)
 * [AnimationCurve](animationcurve.html)
 * [App](app.html)
 * [AppMetadata](appmetadata.html)
 * [BuildType](buildtype.html)
 * [Canvas](canvas.html)
 * [CanvasObject](canvasobject.html)
 * [Clipboard](clipboard.html)
 * [Container](container.html)
 * [Delta](delta.html)
 * [Device](device.html)
 * [DeviceOrientation](deviceorientation.html)
 * [Disableable](disableable.html)
 * [DoubleTappable](doubletappable.html)
 * [DragEvent](dragevent.html)
 * [Draggable](draggable.html)
 * [Driver](driver.html)
 * [Focusable](focusable.html)
 * [HardwareKey](hardwarekey.html)
 * [KeyEvent](keyevent.html)
 * [KeyModifier](keymodifier.html)
 * [KeyName](keyname.html)
 * [KeyboardShortcut](keyboardshortcut.html)
 * [Layout](layout.html)
 * [LegacyTheme](legacytheme.html)
 * [Lifecycle](lifecycle.html)
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
 * [Tabbable](tabbable.html)
 * [Tappable](tappable.html)
 * [TextAlign](textalign.html)
 * [TextStyle](textstyle.html)
 * [TextWrap](textwrap.html)
 * [Theme](theme.html)
 * [ThemeColorName](themecolorname.html)
 * [ThemeIconName](themeiconname.html)
 * [ThemeSizeName](themesizename.html)
 * [ThemeVariant](themevariant.html)
 * [URI](uri.html)
 * [URIReadCloser](urireadcloser.html)
 * [URIWriteCloser](uriwritecloser.html)
 * [Validatable](validatable.html)
 * [Vector2](vector2.html)
 * [Widget](widget.html)
 * [WidgetRenderer](widgetrenderer.html)
 * [Window](window.html)
