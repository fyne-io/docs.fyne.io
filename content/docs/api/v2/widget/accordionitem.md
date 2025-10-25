---
tags: [api]
title: widget.AccordionItem
slug: accordionitem

aliases:
- /api/v2/widget/accordionitem.html
- /api/v2.0/widget/accordionitem
- /api/v2.0/widget/accordionitem.html
- /api/v2.1/widget/accordionitem
- /api/v2.1/widget/accordionitem.html
- /api/v2.2/widget/accordionitem
- /api/v2.2/widget/accordionitem.html
- /api/v2.3/widget/accordionitem
- /api/v2.3/widget/accordionitem.html
- /api/v2.4/widget/accordionitem
- /api/v2.4/widget/accordionitem.html
- /api/v2.5/widget/accordionitem
- /api/v2.5/widget/accordionitem.html
- /api/v2.6/widget/accordionitem
- /api/v2.6/widget/accordionitem.html
- /api/v2.7/widget/accordionitem
- /api/v2.7/widget/accordionitem.html

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type AccordionItem

```go
type AccordionItem struct {
	Title  string
	Detail fyne.CanvasObject
	Open   bool
}
```

AccordionItem represents a single item in an Acc rdion.

#### func  NewAccordionItem

```go
func NewAccordionItem(title string, detail fyne.CanvasObject) *AccordionItem
```
NewAccordionItem creates a new item for an Accordion.
