---
tags: [api]
title: container.Navigation
slug: navigation

aliases:
- /api/v2.0/container/navigation
- /api/v2.1/container/navigation
- /api/v2.2/container/navigation
- /api/v2.3/container/navigation
- /api/v2.4/container/navigation
- /api/v2.5/container/navigation
- /api/v2.6/container/navigation
- /api/v2.7/container/navigation

package: fyne.io/fyne/v2/container
---


---
```go
import "fyne.io/fyne/v2/container"
```

## Usage

#### type Navigation

```go
type Navigation struct {
	widget.BaseWidget

	Root      fyne.CanvasObject
	Title     string
	OnBack    func()
	OnForward func()
}
```

Navigation container is used to provide your application with a control bar and an area for content objects. Objects can be any CanvasObject, and only the most recent one will be visible.


<div class="since">Since: <code>
2.7</code></div>

#### func  NewNavigation

```go
func NewNavigation(root fyne.CanvasObject) *Navigation
```
NewNavigation creates a new navigation container with a given root object.


<div class="since">Since: <code>
2.7</code></div>

#### func  NewNavigationWithTitle

```go
func NewNavigationWithTitle(root fyne.CanvasObject, s string) *Navigation
```
NewNavigationWithTitle creates a new navigation container with a given root object and a default title.


<div class="since">Since: <code>
2.7</code></div>

#### func (*Navigation) Back

```go
func (nav *Navigation) Back() fyne.CanvasObject
```
Back returns the top level CanvasObject, adjusts the title accordingly, and disabled the back button when no more objects are left to go back to.


<div class="since">Since: <code>
2.7</code></div>

#### func (*Navigation) CreateRenderer

```go
func (nav *Navigation) CreateRenderer() fyne.WidgetRenderer
```

#### func (*Navigation) Forward

```go
func (nav *Navigation) Forward() fyne.CanvasObject
```
Forward shows the next object in the stack again.


<div class="since">Since: <code>
2.7</code></div>

#### func (*Navigation) Push

```go
func (nav *Navigation) Push(obj fyne.CanvasObject)
```
Push puts the given object on top of the navigation stack and hides the object below.


<div class="since">Since: <code>
2.7</code></div>

#### func (*Navigation) PushWithTitle

```go
func (nav *Navigation) PushWithTitle(obj fyne.CanvasObject, s string)
```
PushWithTitle puts the given CanvasObject on top, hides the object below, and uses the given title as label text.


<div class="since">Since: <code>
2.7</code></div>

#### func (*Navigation) SetCurrentTitle

```go
func (nav *Navigation) SetCurrentTitle(s string)
```
SetCurrentTitle changes the navigation title for the current level.


<div class="since">Since: <code>
2.7</code></div>

#### func (*Navigation) SetTitle

```go
func (nav *Navigation) SetTitle(s string)
```
SetTitle changes the root navigation title shown by default.


<div class="since">Since: <code>
2.7</code></div>
