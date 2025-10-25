---
tags: [api]
title: widget.Accordion
slug: accordion

aliases:
- /api/v2/widget/accordion.html
- /api/v2.0/widget/accordion
- /api/v2.0/widget/accordion.html
- /api/v2.1/widget/accordion
- /api/v2.1/widget/accordion.html
- /api/v2.2/widget/accordion
- /api/v2.2/widget/accordion.html
- /api/v2.3/widget/accordion
- /api/v2.3/widget/accordion.html
- /api/v2.4/widget/accordion
- /api/v2.4/widget/accordion.html
- /api/v2.5/widget/accordion
- /api/v2.5/widget/accordion.html
- /api/v2.6/widget/accordion
- /api/v2.6/widget/accordion.html
- /api/v2.7/widget/accordion
- /api/v2.7/widget/accordion.html

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Accordion

```go
type Accordion struct {
	BaseWidget
	Items     []*AccordionItem
	MultiOpen bool
}
```

Accordion displays a list of AccordionItems. Each item is represented by a button that reveals a detailed view when tapped.

#### func  NewAccordion

```go
func NewAccordion(items ...*AccordionItem) *Accordion
```
NewAccordion creates a new accordion widget.

#### func (*Accordion) Append

```go
func (a *Accordion) Append(item *AccordionItem)
```
Append adds the given item to this Accordion.

#### func (*Accordion) Close

```go
func (a *Accordion) Close(index int)
```
Close collapses the item at the given index.

#### func (*Accordion) CloseAll

```go
func (a *Accordion) CloseAll()
```
CloseAll collapses all items.

#### func (*Accordion) CreateRenderer

```go
func (a *Accordion) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*Accordion) MinSize

```go
func (a *Accordion) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below.

#### func (*Accordion) Open

```go
func (a *Accordion) Open(index int)
```
Open expands the item at the given index.

#### func (*Accordion) OpenAll

```go
func (a *Accordion) OpenAll()
```
OpenAll expands all items, note that your Accordion should have [MultiOpen] set to `true` for this to operate as expected. For single-open accordions it will open only the first item.

#### func (*Accordion) Prepend

```go
func (a *Accordion) Prepend(item *AccordionItem)
```
Prepend adds the given item to the beginning of this Accordion.


<div class="since">Since: <code>
2.6</code></div>

#### func (*Accordion) Remove

```go
func (a *Accordion) Remove(item *AccordionItem)
```
Remove deletes the given item from this Accordion.

#### func (*Accordion) RemoveIndex

```go
func (a *Accordion) RemoveIndex(index int)
```
RemoveIndex deletes the item at the given index from this Accordion.
