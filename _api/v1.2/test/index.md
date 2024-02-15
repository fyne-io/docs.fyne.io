---
layout: page
tags: [api]
title: Fyne API test
---

# test
---
```go
import "fyne.io/fyne/test"
```

Package test provides utility drivers for running UI tests without rendering

## Usage

#### func  Canvas

```go
func Canvas() fyne.Canvas
```
Canvas returns a reusable in-memory canvas used for testing

#### func  NewApp

```go
func NewApp() fyne.App
```
NewApp returns a new dummy app used for testing. It loads a test driver which creates a virtual window in memory for testing.

#### func  NewClipboard

```go
func NewClipboard() fyne.Clipboard
```
NewClipboard returns a single use in-memory clipboard used for testing

#### func  NewDriver

```go
func NewDriver() fyne.Driver
```
NewDriver sets up and registers a new dummy driver for test purpose

#### func  NewDriverWithPainter

```go
func NewDriverWithPainter(painter SoftwarePainter) fyne.Driver
```
NewDriverWithPainter creates a new dummy driver that will pass the given painter to all canvases created

#### func  NewWindow

```go
func NewWindow(content fyne.CanvasObject) fyne.Window
```
NewWindow creates and registers a new window for test purposes

#### func  Tap

```go
func Tap(obj fyne.Tappable)
```
Tap simulates a left mouse click on the specified object.

#### func  TapAt

```go
func TapAt(obj fyne.Tappable, pos fyne.Position)
```
TapAt simulates a left mouse click on the passed object at a specified place within it.

#### func  TapSecondary

```go
func TapSecondary(obj fyne.Tappable)
```
TapSecondary simulates a right mouse click on the specified object.

#### func  TapSecondaryAt

```go
func TapSecondaryAt(obj fyne.Tappable, pos fyne.Position)
```
TapSecondaryAt simulates a right mouse click on the passed object at a specified place within it.

#### func  Type

```go
func Type(obj fyne.Focusable, chars string)
```
Type performs a series of key events to simulate typing of a value into the specified object. The focusable object will be focused before typing begins. The chars parameter will be input one rune at a time to the focused object.

#### func  TypeOnCanvas

```go
func TypeOnCanvas(c fyne.Canvas, chars string)
```
TypeOnCanvas is like the Type function but it passes the key events to the canvas object rather than a focusable widget.

#### types

 * [SoftwarePainter](softwarepainter.html)
 * [WindowlessCanvas](windowlesscanvas.html)
