---
layout: page
tags: [api]
title: Fyne API "container"
package: fyne.io/fyne/v2/container
---

# container
---
```go
import "fyne.io/fyne/v2/container"
```

Package container provides containers that are used to lay out and organise applications.

## Usage

```go
const (
	// ScrollBoth supports horizontal and vertical scrolling.
	ScrollBoth ScrollDirection = widget.ScrollBoth
	// ScrollHorizontalOnly specifies the scrolling should only happen left to right.
	ScrollHorizontalOnly = widget.ScrollHorizontalOnly
	// ScrollVerticalOnly specifies the scrolling should only happen top to bottom.
	ScrollVerticalOnly = widget.ScrollVerticalOnly
	// ScrollNone turns off scrolling for this container.
	//
	// Since: 2.1
	ScrollNone = widget.ScrollNone
)
```
Constants for valid values of ScrollDirection.

#### func  New

```go
func New(layout fyne.Layout, objects ...fyne.CanvasObject) *fyne.Container
```
New returns a new Container instance holding the specified CanvasObjects which will be laid out according to the specified Layout.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewAdaptiveGrid

```go
func NewAdaptiveGrid(rowcols int, objects ...fyne.CanvasObject) *fyne.Container
```
NewAdaptiveGrid creates a new container with the specified objects and using the grid layout. When in a horizontal arrangement the rowcols parameter will specify the column count, when in vertical it will specify the rows. On mobile this will dynamically refresh when device is rotated.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewBorder

```go
func NewBorder(top, bottom, left, right fyne.CanvasObject, objects ...fyne.CanvasObject) *fyne.Container
```
NewBorder creates a new container with the specified objects and using the border layout. The top, bottom, left and right parameters specify the items that should be placed around edges, the remaining elements will be in the center. Nil can be used to an edge if it should not be filled.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewCenter

```go
func NewCenter(objects ...fyne.CanvasObject) *fyne.Container
```
NewCenter creates a new container with the specified objects centered in the available space.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewGridWithColumns

```go
func NewGridWithColumns(cols int, objects ...fyne.CanvasObject) *fyne.Container
```
NewGridWithColumns creates a new container with the specified objects and using the grid layout with a specified number of columns. The number of rows will depend on how many children are in the container.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewGridWithRows

```go
func NewGridWithRows(rows int, objects ...fyne.CanvasObject) *fyne.Container
```
NewGridWithRows creates a new container with the specified objects and using the grid layout with a specified number of rows. The number of columns will depend on how many children are in the container.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewGridWrap

```go
func NewGridWrap(size fyne.Size, objects ...fyne.CanvasObject) *fyne.Container
```
NewGridWrap creates a new container with the specified objects and using the gridwrap layout. Every element will be resized to the size parameter and the content will arrange along a row and flow to a new row if the elements don't fit.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewHBox

```go
func NewHBox(objects ...fyne.CanvasObject) *fyne.Container
```
NewHBox creates a new container with the specified objects and using the HBox layout. The objects will be placed in the container from left to right and always displayed at their horizontal MinSize. Use a different layout if the objects are intended to be larger then their horizontal MinSize.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewMax

```go
func NewMax(objects ...fyne.CanvasObject) *fyne.Container
```
NewMax creates a new container with the specified objects filling the available space.


<div class="since">Since: <code>
1.4</code></div>


<div class="deprecated">
Deprecated: Use container.NewStack() instead.</div>

#### func  NewPadded

```go
func NewPadded(objects ...fyne.CanvasObject) *fyne.Container
```
NewPadded creates a new container with the specified objects inset by standard padding size.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewStack

```go
func NewStack(objects ...fyne.CanvasObject) *fyne.Container
```
NewStack returns a new container that stacks objects on top of each other. Objects at the end of the container will be stacked on top of objects before. Having only a single object has no impact as CanvasObjects will fill the available space even without a Stack.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewVBox

```go
func NewVBox(objects ...fyne.CanvasObject) *fyne.Container
```
NewVBox creates a new container with the specified objects and using the VBox layout. The objects will be stacked in the container from top to bottom and always displayed at their vertical MinSize. Use a different layout if the objects are intended to be larger then their vertical MinSize.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewWithoutLayout

```go
func NewWithoutLayout(objects ...fyne.CanvasObject) *fyne.Container
```
NewWithoutLayout returns a new Container instance holding the specified CanvasObjects that are manually arranged.


<div class="since">Since: <code>
2.0</code></div>

#### types

 * [AppTabs](apptabs.html)
 * [DocTabs](doctabs.html)
 * [Scroll](scroll.html)
 * [ScrollDirection](scrolldirection.html)
 * [Split](split.html)
 * [TabItem](tabitem.html)
 * [TabLocation](tablocation.html)
