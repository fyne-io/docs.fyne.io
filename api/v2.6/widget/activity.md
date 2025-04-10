---
layout: page
tags: [api]
title: Fyne API "widget.Activity"
package: fyne.io/fyne/v2/widget
---

# widget.Activity
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Activity

```go
type Activity struct {
	BaseWidget
}
```

Activity is used to indicate that something is happening that should be waited for, or is in the background (depending on usage).


<div class="since">Since: <code>
2.5</code></div>

#### func  NewActivity

```go
func NewActivity() *Activity
```
NewActivity returns a widget for indicating activity


<div class="since">Since: <code>
2.5</code></div>

#### func (*Activity) CreateRenderer

```go
func (a *Activity) CreateRenderer() fyne.WidgetRenderer
```

#### func (*Activity) MinSize

```go
func (a *Activity) MinSize() fyne.Size
```

#### func (*Activity) Start

```go
func (a *Activity) Start()
```
Start the activity indicator animation

#### func (*Activity) Stop

```go
func (a *Activity) Stop()
```
Stop the activity indicator animation
