---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
---
```go
import "fyne.io/fyne/widget"
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

AccordionItem represents a single item in an AccordionContainer.

#### func  NewAccordionItem

```go
func NewAccordionItem(title string, detail fyne.CanvasObject) *AccordionItem
```
NewAccordionItem creates a new item for an Accordion.
