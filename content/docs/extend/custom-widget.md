---
title: Writing a Custom Widget
slug: custom-widget

weight: 8020

aliases:
- /extend/custom-widget.html 
- /develop/custom-widget.html
- /tutorial/write-custom-widget
---

The standard widgets included with Fyne are designed to support standard user interactions and requirements.
As a GUI often has to provide custom functionality it may be necessary to write a custom widget.
This article outlines how.

A widget is split into two areas - each implementing a standard interface - the `fyne.Widget` and the `fyne.WidgetRenderer`. The widget defines behaviour and state, with the renderer being used to define how it should be drawn to screen.

### fyne.Widget

A widget in Fyne is simply a stateful canvas object that has its rendering
definition separated from the main logic. As you can see from the `fyne.Widget`
interface there is not much that must be implemented.

```go
type Widget interface {
	CanvasObject

	CreateRenderer() WidgetRenderer
}

```

As a widget needs to be used like any other canvas item we inherit from the same interface.
To save writing all the functions required we can make use of the `widget.BaseWidget` type which handles the basics.

Each widget definition will contain much more than the interface requires.
It is standard in a Fyne widget to export the fields which define behaviour
(just like the primitives defined in the `canvas` package).

For example, look at the `widget.Button` type:

```go
type Button struct {
	BaseWidget
	Text  string
	Style ButtonStyle
	Icon  fyne.Resource

	OnTapped func()
}
```

You can see how each of these items store state about the widget behaviour but nothing about how it is rendered.

### fyne.WidgetRenderer

The widget renderer is responsible for managing a list of `fyne.CanvasObject` primitives that come together to create the design of our widget. It is much like
a `fyne.Container` with a custom layout and some additional theme handling.

Every widget must provide a renderer, but it is entirely possible to re-use a renderer from another widget - especially if your widget is a lightweight wrapper around another standard control.

```go
type WidgetRenderer interface {
	Layout(Size)
	MinSize() Size

	Refresh()
	Objects() []CanvasObject
	Destroy()
}
```

As you can see the `Layout(Size)` and `MinSize()` functions are similar to the
`fyne.Layout` interface, without the `[]fyne.CanvasObject` parameter - this is because a widget does need to be laid out but it controls which objects will be included.

The `Refresh()` method is triggered when the widget this renderer draws has changed or if the theme is altered.
In either situation we may need to make adjustments to how it looks.
Lastly the `Destroy()` method is called when this renderer is no longer needed so it should clear any resources that would otherwise leak.

Compare again with the button widget - it's `fyne.WidgetRenderer` implementation is based on the following type:

```go
type buttonRenderer struct {
	icon   *canvas.Image
	label  *canvas.Text
	shadow *fyne.CanvasObject

	objects []fyne.CanvasObject
	button  *Button
}
```

As you can see it has fields to cache the actual image, text and shadow canvas
objects for drawing.
It keeps track of the slice of objects required by `fyne.WidgetRenderer` as a convenience.

Lastly it keeps a reference to the `widget.Button` for all state information.
In the `Refresh()` method it will update the graphical state based on any changes
in the underlying `widget.Button` type.

### Bring it together

A basic widget will extend the `widget.BaseWidget` type and declare any state
that the widget holds.
The `CreateRenderer()` function must exist and return a new `fyne.WidgetRenderer` instance. The widget and driver code in Fyne will ensure that this
is cached accordingly - this method may be called many times (for example if a
widget is hidden and then shown). If `CreateRenderer()` is called again you
should return a new renderer instance as the old one may have been destroyed.

Take care not to keep any important state in your renderer - animation tickers
are well suited to that location but user state would not be. A widget that is
hidden may have it's renderer destroyed and if it is shown again the new renderer
must be able to reflect the same widget state.

#### Using SimpleRenderer for custom widgets

Custom widgets built from a single `CanvasObject`, for example a container wrapping
multiple builtin widgets, can be implemented using `SimpleRenderer`.
Below example is a custom widget that can be used as item in a list view,
showing a title on the left hand side that will be truncated when too long, and
a comment on the right hand side. The constructor would be called from the
`CreateItem` function of the list, and the title and comment changed in the
`UpdateItem` function:

```go
type MyListItemWidget struct {
	widget.BaseWidget
	Title   *widget.Label
	Comment *widget.Label
}

func NewMyListItemWidget(title, comment string) *MyListItemWidget {
	item := &MyListItemWidget{
		Title:   widget.NewLabel(title),
		Comment: widget.NewLabel(comment),
	}
	item.Title.Truncation = fyne.TextTruncateEllipsis
	item.ExtendBaseWidget(item)

	return item
}

func (item *MyListItemWidget) CreateRenderer() fyne.WidgetRenderer {
	c := container.NewBorder(nil, nil, nil, item.Comment, item.Title)
	return widget.NewSimpleRenderer(c)
}
```
