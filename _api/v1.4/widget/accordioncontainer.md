---
layout: page
tags: [api]
title: Fyne API "widget.AccordionContainer"
---

# widget.AccordionContainer
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type AccordionContainer

```go
type AccordionContainer = Accordion
```

AccordionContainer displays a list of AccordionItems. Each item is represented by a button that reveals a detailed view when tapped.


<div class="deprecated">
Deprecated: This has been renamed to Accordion</div>

#### func  NewAccordionContainer

```go
func NewAccordionContainer(items ...*AccordionItem) *AccordionContainer
```
NewAccordionContainer creates a new accordion widget.


<div class="deprecated">
Deprecated: Use NewAccordion instead</div>
