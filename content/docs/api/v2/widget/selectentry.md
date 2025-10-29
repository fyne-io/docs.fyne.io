---
tags: [api]
title: widget.SelectEntry
slug: selectentry

aliases:
- /api/widget/selectentry
- /api/widget/selectentry.html
- /api/v2.0/widget/selectentry
- /api/v2.0/widget/selectentry.html
- /api/v2.1/widget/selectentry
- /api/v2.1/widget/selectentry.html
- /api/v2.2/widget/selectentry
- /api/v2.2/widget/selectentry.html
- /api/v2.3/widget/selectentry
- /api/v2.3/widget/selectentry.html
- /api/v2.4/widget/selectentry
- /api/v2.4/widget/selectentry.html
- /api/v2.5/widget/selectentry
- /api/v2.5/widget/selectentry.html
- /api/v2.6/widget/selectentry
- /api/v2.6/widget/selectentry.html
- /api/v2.7/widget/selectentry
- /api/v2.7/widget/selectentry.html

package: fyne.io/fyne/v2/widget
---


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

#### func (*SelectEntry) Disable

```go
func (e *SelectEntry) Disable()
```
Disable this widget so that it cannot be interacted with, updating any style appropriately.

#### func (*SelectEntry) Enable

```go
func (e *SelectEntry) Enable()
```
Enable this widget, updating any style or features appropriately.

#### func (*SelectEntry) MinSize

```go
func (e *SelectEntry) MinSize() fyne.Size
```
MinSize returns the minimal size of the select entry.

#### func (*SelectEntry) Move

```go
func (e *SelectEntry) Move(pos fyne.Position)
```
Move changes the relative position of the select entry.

#### func (*SelectEntry) Resize

```go
func (e *SelectEntry) Resize(size fyne.Size)
```
Resize changes the size of the select entry.

#### func (*SelectEntry) SetOptions

```go
func (e *SelectEntry) SetOptions(options []string)
```
SetOptions sets the options the user might select from.
