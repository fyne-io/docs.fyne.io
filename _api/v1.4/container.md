---
layout: page
tags: [api]
title: Fyne API "fyne.Container"
---

# fyne.Container
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


<div class="deprecated">
Deprecated: Use NewContainerWithoutLayout to create a container that uses manual layout.</div>

#### func  NewContainerWithLayout

```go
func NewContainerWithLayout(layout Layout, objects ...CanvasObject) *Container
```
NewContainerWithLayout returns a new Container instance holding the specified CanvasObjects which will be laid out according to the specified Layout.

#### func  NewContainerWithoutLayout

```go
func NewContainerWithoutLayout(objects ...CanvasObject) *Container
```
NewContainerWithoutLayout returns a new Container instance holding the specified CanvasObjects that are manually arranged.

#### func (*Container) Add

```go
func (c *Container) Add(add CanvasObject)
```
Add appends the specified object to the items this container manages.


<div class="since">Since: <code>
1.4</code></div>

#### func (*Container) AddObject

```go
func (c *Container) AddObject(o CanvasObject)
```
AddObject adds another CanvasObject to the set this Container holds.


<div class="deprecated">
Deprecated: Use replacement Add() function</div>

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

#### func (*Container) Remove

```go
func (c *Container) Remove(rem CanvasObject)
```
Remove updates the contents of this container to no longer include the specified object.

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
