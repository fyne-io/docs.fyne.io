---
layout: page
tags: [api]
title: Fyne API "layout"
package: fyne.io/fyne/v2/layout
---

# layout
---
```go
import "fyne.io/fyne/v2/layout"
```

Package layout defines the various layouts available to Fyne apps.

## Usage

#### func  NewAdaptiveGridLayout

```go
func NewAdaptiveGridLayout(rowcols int) fyne.Layout
```
NewAdaptiveGridLayout returns a new grid layout which uses columns when horizontal but rows when vertical.

#### func  NewBorderLayout

```go
func NewBorderLayout(top, bottom, left, right fyne.CanvasObject) fyne.Layout
```
NewBorderLayout creates a new BorderLayout instance with top, bottom, left and right objects set. All other items in the container will fill the centre space

#### func  NewCenterLayout

```go
func NewCenterLayout() fyne.Layout
```
NewCenterLayout creates a new CenterLayout instance

#### func  NewFormLayout

```go
func NewFormLayout() fyne.Layout
```
NewFormLayout returns a new FormLayout instance

#### func  NewGridLayout

```go
func NewGridLayout(cols int) fyne.Layout
```
NewGridLayout returns a grid layout arranged in a specified number of columns. The number of rows will depend on how many children are in the container that uses this layout.

#### func  NewGridLayoutWithColumns

```go
func NewGridLayoutWithColumns(cols int) fyne.Layout
```
NewGridLayoutWithColumns returns a new grid layout that specifies a column count and wrap to new rows when needed.

#### func  NewGridLayoutWithRows

```go
func NewGridLayoutWithRows(rows int) fyne.Layout
```
NewGridLayoutWithRows returns a new grid layout that specifies a row count that creates new rows as required.

#### func  NewGridWrapLayout

```go
func NewGridWrapLayout(size fyne.Size) fyne.Layout
```
NewGridWrapLayout returns a new GridWrapLayout instance

#### func  NewHBoxLayout

```go
func NewHBoxLayout() fyne.Layout
```
NewHBoxLayout returns a horizontal box layout for stacking a number of child canvas objects or widgets left to right. The objects are always displayed at their horizontal MinSize. Use a different layout if the objects are intended to be larger then their horizontal MinSize.

#### func  NewMaxLayout

```go
func NewMaxLayout() fyne.Layout
```
NewMaxLayout creates a new MaxLayout instance


<div class="deprecated">
Deprecated: Use layout.NewStackLayout() instead.</div>

#### func  NewPaddedLayout

```go
func NewPaddedLayout() fyne.Layout
```
NewPaddedLayout creates a new PaddedLayout instance


<div class="since">Since: <code>
1.4</code></div>

#### func  NewSpacer

```go
func NewSpacer() fyne.CanvasObject
```
NewSpacer returns a spacer object which can fill vertical and horizontal space. This is primarily used with a box layout.

#### func  NewStackLayout

```go
func NewStackLayout() fyne.Layout
```
NewStackLayout returns a new StackLayout instance. Objects are stacked on top of each other with later objects on top of those before. Having only a single object has no impact as CanvasObjects will fill the available space even without a Stack.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewVBoxLayout

```go
func NewVBoxLayout() fyne.Layout
```
NewVBoxLayout returns a vertical box layout for stacking a number of child canvas objects or widgets top to bottom. The objects are always displayed at their vertical MinSize. Use a different layout if the objects are intended to be larger then their vertical MinSize.

#### types

 * [Spacer](spacer.html)
 * [SpacerObject](spacerobject.html)
