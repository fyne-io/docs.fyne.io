---
layout: page
tags: [api]
title: Fyne API "widget.FormItem"
---

# widget.FormItem
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type FormItem

```go
type FormItem struct {
	Text   string
	Widget fyne.CanvasObject
}
```

FormItem provides the details for a row in a form

#### func  NewFormItem

```go
func NewFormItem(text string, widget fyne.CanvasObject) *FormItem
```
NewFormItem creates a new form item with the specified label text and input widget
