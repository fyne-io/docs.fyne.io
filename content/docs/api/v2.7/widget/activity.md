---
tags: [api]
title: widget.Activity
slug: activity

aliases:
- /api/v2.7/widget/activity

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type Activity struct {
	BaseWidget
}
```

Activity is used to indicate that something is happening that should be waited for, or is in the background (depending on usage).


<div class="since">Since: <code>
2.5</code></div>

###

```go
func NewActivity() *Activity
```
NewActivity returns a widget for indicating activity


<div class="since">Since: <code>
2.5</code></div>

###

```go
func (a *Activity) CreateRenderer() fyne.WidgetRenderer
```

###

```go
func (a *Activity) MinSize() fyne.Size
```

###

```go
func (a *Activity) Start()
```
Start the activity indicator animation

###

```go
func (a *Activity) Stop()
```
Stop the activity indicator animation
