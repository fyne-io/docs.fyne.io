---
tags: [api]
title: widget.Card
slug: card

aliases:
- /api/v2.7/widget/card

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

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

###

```go
func NewCard(title, subtitle string, content fyne.CanvasObject) *Card
```
NewCard creates a new card widget with the specified title, subtitle and content (all optional).


<div class="since">Since: <code>
1.4</code></div>

###

```go
func (c *Card) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

###

```go
func (c *Card) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

###

```go
func (c *Card) SetContent(obj fyne.CanvasObject)
```
SetContent changes the body of this card to have the specified content.

###

```go
func (c *Card) SetImage(img *canvas.Image)
```
SetImage changes the image displayed above the title for this card.

###

```go
func (c *Card) SetSubTitle(text string)
```
SetSubTitle updates the secondary title for this card.

###

```go
func (c *Card) SetTitle(text string)
```
SetTitle updates the main title for this card.
