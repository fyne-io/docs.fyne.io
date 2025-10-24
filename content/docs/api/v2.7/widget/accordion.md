---
tags: [api]
title: widget.Accordion
slug: accordion

aliases:
- /api/v2.7/widget/accordion

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type Accordion struct {
	BaseWidget
	Items     []*AccordionItem
	MultiOpen bool
}
```

Accordion displays a list of AccordionItems. Each item is represented by a button that reveals a detailed view when tapped.

###

```go
func NewAccordion(items ...*AccordionItem) *Accordion
```
NewAccordion creates a new accordion widget.

###

```go
func (a *Accordion) Append(item *AccordionItem)
```
Append adds the given item to this Accordion.

###

```go
func (a *Accordion) Close(index int)
```
Close collapses the item at the given index.

###

```go
func (a *Accordion) CloseAll()
```
CloseAll collapses all items.

###

```go
func (a *Accordion) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

###

```go
func (a *Accordion) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below.

###

```go
func (a *Accordion) Open(index int)
```
Open expands the item at the given index.

###

```go
func (a *Accordion) OpenAll()
```
OpenAll expands all items, note that your Accordion should have [MultiOpen] set to `true` for this to operate as expected. For single-open accordions it will open only the first item.

###

```go
func (a *Accordion) Prepend(item *AccordionItem)
```
Prepend adds the given item to the beginning of this Accordion.


<div class="since">Since: <code>
2.6</code></div>

###

```go
func (a *Accordion) Remove(item *AccordionItem)
```
Remove deletes the given item from this Accordion.

###

```go
func (a *Accordion) RemoveIndex(index int)
```
RemoveIndex deletes the item at the given index from this Accordion.
