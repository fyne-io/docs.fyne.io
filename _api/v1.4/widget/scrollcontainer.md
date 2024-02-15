---
layout: page
tags: [api]
title: Fyne API "widget.ScrollContainer"
---

# widget.ScrollContainer
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type ScrollContainer

```go
type ScrollContainer struct {
	BaseWidget

	Direction ScrollDirection
	Content   fyne.CanvasObject
	Offset    fyne.Position
}
```

ScrollContainer defines a container that is smaller than the Content. The Offset is used to determine the position of the child widgets within the container.


<div class="deprecated">
Deprecated: use container.Scroll instead.</div>

#### func  NewHScrollContainer

```go
func NewHScrollContainer(content fyne.CanvasObject) *ScrollContainer
```
NewHScrollContainer create a scrollable parent wrapping the specified content. Note that this may cause the MinSize.Width to be smaller than that of the passed object.


<div class="deprecated">
Deprecated: use container.NewHScroll instead.</div>

#### func  NewScrollContainer

```go
func NewScrollContainer(content fyne.CanvasObject) *ScrollContainer
```
NewScrollContainer creates a scrollable parent wrapping the specified content. Note that this may cause the MinSize to be smaller than that of the passed object.


<div class="deprecated">
Deprecated: use container.NewScroll instead.</div>

#### func  NewVScrollContainer

```go
func NewVScrollContainer(content fyne.CanvasObject) *ScrollContainer
```
NewVScrollContainer create a scrollable parent wrapping the specified content. Note that this may cause the MinSize.Height to be smaller than that of the passed object.


<div class="deprecated">
Deprecated: use container.NewVScroll instead.</div>

#### func (*ScrollContainer) CreateRenderer

```go
func (s *ScrollContainer) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*ScrollContainer) DragEnd

```go
func (s *ScrollContainer) DragEnd()
```
DragEnd will stop scrolling on mobile has stopped

#### func (*ScrollContainer) Dragged

```go
func (s *ScrollContainer) Dragged(e *fyne.DragEvent)
```
Dragged will scroll on any drag - bar or otherwise - for mobile

#### func (*ScrollContainer) MinSize

```go
func (s *ScrollContainer) MinSize() fyne.Size
```
MinSize returns the smallest size this widget can shrink to

#### func (*ScrollContainer) Refresh

```go
func (s *ScrollContainer) Refresh()
```
Refresh causes this widget to be redrawn in it's current state

#### func (*ScrollContainer) Resize

```go
func (s *ScrollContainer) Resize(size fyne.Size)
```
Resize sets a new size for the scroll container.

#### func (*ScrollContainer) ScrollToBottom

```go
func (s *ScrollContainer) ScrollToBottom()
```
ScrollToBottom will scroll content to container bottom - to show latest info which end user just added

#### func (*ScrollContainer) ScrollToTop

```go
func (s *ScrollContainer) ScrollToTop()
```
ScrollToTop will scroll content to container top

#### func (*ScrollContainer) Scrolled

```go
func (s *ScrollContainer) Scrolled(ev *fyne.ScrollEvent)
```
Scrolled is called when an input device triggers a scroll event

#### func (*ScrollContainer) SetMinSize

```go
func (s *ScrollContainer) SetMinSize(size fyne.Size)
```
SetMinSize specifies a minimum size for this scroll container. If the specified size is larger than the content size then scrolling will not be enabled This can be helpful to appear larger than default if the layout is collapsing this widget.
