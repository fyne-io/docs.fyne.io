---
tags: [api]
title: widget.FormItem
slug: formitem

aliases:
- /api/v2.7/widget/formitem

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type FormItem struct {
	Text   string
	Widget fyne.CanvasObject

	// Since: 2.0
	HintText string
}
```

FormItem provides the details for a row in a form

###

```go
func NewFormItem(text string, widget fyne.CanvasObject) *FormItem
```
NewFormItem creates a new form item with the specified label text and input widget
