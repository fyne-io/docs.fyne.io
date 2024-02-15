---
layout: page
tags: [api]
title: Fyne API "widget.Card"
---

# widget.Card
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Card

```go
type Card struct {
	BaseWidget
	Title, Subtitle string
	Image           *canvas.Image
	Content         fyne.CanvasObject
}
```

Card widget groups title, subtitle with content and a header image


<div class="since">Since: <code>
1.4</code></div>

#### func  NewCard

```go
func NewCard(title, subtitle string, content fyne.CanvasObject) *Card
```
NewCard creates a new card widget with the specified title, subtitle and content (all optional).


<div class="since">Since: <code>
1.4</code></div>

#### func (*Card) CreateRenderer

```go
func (c *Card) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*Card) MinSize

```go
func (c *Card) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*Card) SetContent

```go
func (c *Card) SetContent(obj fyne.CanvasObject)
```
SetContent changes the body of this card to have the specified content.

#### func (*Card) SetImage

```go
func (c *Card) SetImage(img *canvas.Image)
```
SetImage changes the image displayed above the title for this card.

#### func (*Card) SetSubTitle

```go
func (c *Card) SetSubTitle(text string)
```
SetSubTitle updates the secondary title for this card.

#### func (*Card) SetTitle

```go
func (c *Card) SetTitle(text string)
```
SetTitle updates the main title for this card.
