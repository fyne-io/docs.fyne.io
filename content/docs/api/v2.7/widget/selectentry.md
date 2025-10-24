---
tags: [api]
title: widget.SelectEntry
slug: selectentry

aliases:
- /api/v2.7/widget/selectentry

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type SelectEntry struct {
	Entry
}
```

SelectEntry is an input field which supports selecting from a fixed set of options.

###

```go
func NewSelectEntry(options []string) *SelectEntry
```
NewSelectEntry creates a SelectEntry.

###

```go
func (e *SelectEntry) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer returns a new renderer for this select entry.

###

```go
func (e *SelectEntry) Disable()
```
Disable this widget so that it cannot be interacted with, updating any style appropriately.

###

```go
func (e *SelectEntry) Enable()
```
Enable this widget, updating any style or features appropriately.

###

```go
func (e *SelectEntry) MinSize() fyne.Size
```
MinSize returns the minimal size of the select entry.

###

```go
func (e *SelectEntry) Move(pos fyne.Position)
```
Move changes the relative position of the select entry.

###

```go
func (e *SelectEntry) Resize(size fyne.Size)
```
Resize changes the size of the select entry.

###

```go
func (e *SelectEntry) SetOptions(options []string)
```
SetOptions sets the options the user might select from.
