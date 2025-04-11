---
layout: page
tags: [api]
title: Fyne API "fyne.Container"
package: fyne.io/fyne/v2
---

# fyne.Container
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Container

```go
type Container struct {
	Hidden bool // Is this Container hidden

	Layout Layout // The Layout algorithm for arranging child [CanvasObject]s

	Objects []CanvasObject // The set of [CanvasObject]s this container holds
}
```

Container is a [CanvasObject] that contains a collection of child objects. The layout of the children is set by the specified Layout.

#### func  NewContainer

```go
func NewContainer(objects ...CanvasObject) *Container
```
NewContainer returns a new [Container] instance holding the specified [CanvasObject]s.


<div class="deprecated">
Deprecated: Use [fyne.io/fyne/v2/container.NewWithoutLayout] to create a container that uses manual layout.</div>

#### func  NewContainerWithLayout

```go
func NewContainerWithLayout(layout Layout, objects ...CanvasObject) *Container
```
NewContainerWithLayout returns a new [Container] instance holding the specified [CanvasObject]s which will be laid out according to the specified Layout.


<div class="deprecated">
Deprecated: Use [fyne.io/fyne/v2/container.New] instead.</div>

#### func  NewContainerWithoutLayout

```go
func NewContainerWithoutLayout(objects ...CanvasObject) *Container
```
NewContainerWithoutLayout returns a new [Container] instance holding the specified [CanvasObject]s that are manually arranged.


<div class="deprecated">
Deprecated: Use [fyne.io/fyne/v2/container.NewWithoutLayout] instead.</div>

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
AddObject adds another [CanvasObject] to the set this Container holds.


<div class="deprecated">
Deprecated: Use [Container.Add] instead.</div>

#### func (*Container) Hide

```go
func (c *Container) Hide()
```
Hide sets this container, and all its children, to be not visible.

#### func (*Container) MinSize

```go
func (c *Container) MinSize() Size
```
MinSize calculates the minimum size of c. This is delegated to the [Container.Layout], if specified, otherwise it will be calculated.

#### func (*Container) Move

```go
func (c *Container) Move(pos Position)
```
Move the container (and all its children) to a new position, relative to its parent.

#### func (*Container) Position

```go
func (c *Container) Position() Position
```
Position gets the current position of c relative to its parent.

#### func (*Container) Refresh

```go
func (c *Container) Refresh()
```
Refresh causes this object to be redrawn in it's current state

#### func (*Container) Remove

```go
func (c *Container) Remove(rem CanvasObject)
```
Remove updates the contents of this container to no longer include the specified object. This method is not intended to be used inside a loop, to remove all the elements. It is much more efficient to call [Container.RemoveAll) instead.

#### func (*Container) RemoveAll

```go
func (c *Container) RemoveAll()
```
RemoveAll updates the contents of this container to no longer include any objects.


<div class="since">Since: <code>
2.2</code></div>

#### func (*Container) Resize

```go
func (c *Container) Resize(size Size)
```
Resize sets a new size for c.

#### func (*Container) Show

```go
func (c *Container) Show()
```
Show sets this container, and all its children, to be visible.

#### func (*Container) Size

```go
func (c *Container) Size() Size
```
Size returns the current size c.

#### func (*Container) Visible

```go
func (c *Container) Visible() bool
```
Visible returns true if the container is currently visible, false otherwise.
