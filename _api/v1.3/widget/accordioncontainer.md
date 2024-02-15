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

#### type AccordionContainer

```go
type AccordionContainer struct {
	BaseWidget
	Items     []*AccordionItem
	MultiOpen bool
}
```

AccordionContainer displays a list of AccordionItems. Each item is represented by a button that reveals a detailed view when tapped.

#### func  NewAccordionContainer

```go
func NewAccordionContainer(items ...*AccordionItem) *AccordionContainer
```
NewAccordionContainer creates a new accordion widget.

#### func (*AccordionContainer) Append

```go
func (a *AccordionContainer) Append(item *AccordionItem)
```
Append adds the given item to this AccordionContainer.

#### func (*AccordionContainer) Close

```go
func (a *AccordionContainer) Close(index int)
```
Close collapses the item at the given index.

#### func (*AccordionContainer) CloseAll

```go
func (a *AccordionContainer) CloseAll()
```
CloseAll collapses all items.

#### func (*AccordionContainer) CreateRenderer

```go
func (a *AccordionContainer) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*AccordionContainer) MinSize

```go
func (a *AccordionContainer) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below.

#### func (*AccordionContainer) Open

```go
func (a *AccordionContainer) Open(index int)
```
Open expands the item at the given index.

#### func (*AccordionContainer) OpenAll

```go
func (a *AccordionContainer) OpenAll()
```
OpenAll expands all items.

#### func (*AccordionContainer) Remove

```go
func (a *AccordionContainer) Remove(item *AccordionItem)
```
Remove deletes the given item from this AccordionContainer.

#### func (*AccordionContainer) RemoveIndex

```go
func (a *AccordionContainer) RemoveIndex(index int)
```
RemoveIndex deletes the item at the given index from this AccordionContainer.
