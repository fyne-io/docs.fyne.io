---
layout: page
tags: [api]
title: Fyne API "widget.SelectEntry"
package: fyne.io/fyne/v2/widget
---

# widget.SelectEntry
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type SelectEntry

```go
type SelectEntry struct {
	Entry
}
```

SelectEntry is an input field which supports selecting from a fixed set of options.

#### func  NewSelectEntry

```go
func NewSelectEntry(options []string) *SelectEntry
```
NewSelectEntry creates a SelectEntry.

#### func (*SelectEntry) CreateRenderer

```go
func (e *SelectEntry) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer returns a new renderer for this select entry.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*SelectEntry) Disable

```go
func (e *SelectEntry) Disable()
```
Disable this widget so that it cannot be interacted with, updating any style appropriately.


<div class="implements">Implements: <code>
fyne.DisableableWidget</code></div>

#### func (*SelectEntry) Enable

```go
func (e *SelectEntry) Enable()
```
Enable this widget, updating any style or features appropriately.


<div class="implements">Implements: <code>
fyne.DisableableWidget</code></div>

#### func (*SelectEntry) MinSize

```go
func (e *SelectEntry) MinSize() fyne.Size
```
MinSize returns the minimal size of the select entry.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*SelectEntry) Move

```go
func (e *SelectEntry) Move(pos fyne.Position)
```
Move changes the relative position of the select entry.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*SelectEntry) Resize

```go
func (e *SelectEntry) Resize(size fyne.Size)
```
Resize changes the size of the select entry.


<div class="implements">Implements: <code>
fyne.Widget</code></div>

#### func (*SelectEntry) SetOptions

```go
func (e *SelectEntry) SetOptions(options []string)
```
SetOptions sets the options the user might select from.
