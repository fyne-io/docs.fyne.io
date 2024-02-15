---
layout: page
tags: [api]
title: Fyne API fyne
---

# fyne
---
```go
import "fyne.io/fyne"
```

## Usage

#### type Container

```go
type Container struct {
	Hidden bool // Is this Container hidden

	Layout  Layout         // The Layout algorithm for arranging child CanvasObjects
	Objects []CanvasObject // The set of CanvasObjects this container holds
}
```

Container is a CanvasObject that contains a collection of child objects. The layout of the children is set by the specified Layout.

#### func  NewContainer

```go
func NewContainer(objects ...CanvasObject) *Container
```
NewContainer returns a new Container instance holding the specified CanvasObjects.

#### func  NewContainerWithLayout

```go
func NewContainerWithLayout(layout Layout, objects ...CanvasObject) *Container
```
NewContainerWithLayout returns a new Container instance holding the specified CanvasObjects which will be laid out according to the specified Layout.

#### func (*Container) AddObject

```go
func (c *Container) AddObject(o CanvasObject)
```
AddObject adds another CanvasObject to the set this Container holds.

#### func (*Container) Hide

```go
func (c *Container) Hide()
```
Hide sets this container, and all its children, to be not visible.

#### func (*Container) MinSize

```go
func (c *Container) MinSize() Size
```
MinSize calculates the minimum size of a Container. This is delegated to the Layout, if specified, otherwise it will mimic MaxLayout.

#### func (*Container) Move

```go
func (c *Container) Move(pos Position)
```
Move the container (and all its children) to a new position, relative to its parent.

#### func (*Container) Position

```go
func (c *Container) Position() Position
```
Position gets the current position of this Container, relative to its parent.

#### func (*Container) Refresh

```go
func (c *Container) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Container) Resize

```go
func (c *Container) Resize(size Size)
```
Resize sets a new size for the Container.

#### func (*Container) Show

```go
func (c *Container) Show()
```
Show sets this container, and all its children, to be visible.

#### func (*Container) Size

```go
func (c *Container) Size() Size
```
Size returns the current size of this container.

#### func (*Container) Visible

```go
func (c *Container) Visible() bool
```
Visible returns true if the container is currently visible, false otherwise.
