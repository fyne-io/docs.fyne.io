---
layout: page
tags: [api]
title: Fyne API "fyne.Driver"
package: fyne.io/fyne/v2
---

# fyne.Driver
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Driver

```go
type Driver interface {
	// CreateWindow creates a new UI Window.
	CreateWindow(string) Window
	// AllWindows returns a slice containing all app windows.
	AllWindows() []Window

	// RenderedTextSize returns the size required to render the given string of specified
	// font size and style. It also returns the height to text baseline, measured from the top.
	// If the source is specified it will be used, otherwise the current theme will be asked for the font.
	RenderedTextSize(text string, fontSize float32, style TextStyle, source Resource) (size Size, baseline float32)

	// CanvasForObject returns the canvas that is associated with a given [CanvasObject].
	CanvasForObject(CanvasObject) Canvas
	// AbsolutePositionForObject returns the position of a given [CanvasObject] relative to the top/left of a canvas.
	AbsolutePositionForObject(CanvasObject) Position

	// Device returns the device that the application is currently running on.
	Device() Device
	// Run starts the main event loop of the driver.
	Run()
	// Quit closes the driver and open windows, then exit the application.
	// On some some operating systems this does nothing, for example iOS and Android.
	Quit()

	// StartAnimation registers a new animation with this driver and requests it be started.
	StartAnimation(*Animation)
	// StopAnimation stops an animation and unregisters from this driver.
	StopAnimation(*Animation)

	// DoubleTapDelay returns the maximum duration where a second tap after a first one
	// will be considered a [DoubleTap] instead of two distinct [Tap] events.
	//
	// Since: 2.5
	DoubleTapDelay() time.Duration

	// SetDisableScreenBlanking allows an app to ask the device not to sleep/lock/blank displays
	//
	// Since: 2.5
	SetDisableScreenBlanking(bool)
}
```

Driver defines an abstract concept of a Fyne render driver. Any implementation must provide at least these methods.
