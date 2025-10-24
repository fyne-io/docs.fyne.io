---
tags: [api]
title: widget.AccordionItem
slug: accordionitem

aliases:
- /api/v2.7/widget/accordionitem

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type AccordionItem struct {
	Title  string
	Detail fyne.CanvasObject
	Open   bool
}
```

AccordionItem represents a single item in an Acc rdion.

###

```go
func NewAccordionItem(title string, detail fyne.CanvasObject) *AccordionItem
```
NewAccordionItem creates a new item for an Accordion.
