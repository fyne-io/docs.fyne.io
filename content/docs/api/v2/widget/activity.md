---
tags: [api]
title: widget.Activity
slug: activity

aliases:
- /api/v2/widget/activity.html
- /api/v2.0/widget/activity
- /api/v2.0/widget/activity.html
- /api/v2.1/widget/activity
- /api/v2.1/widget/activity.html
- /api/v2.2/widget/activity
- /api/v2.2/widget/activity.html
- /api/v2.3/widget/activity
- /api/v2.3/widget/activity.html
- /api/v2.4/widget/activity
- /api/v2.4/widget/activity.html
- /api/v2.5/widget/activity
- /api/v2.5/widget/activity.html
- /api/v2.6/widget/activity
- /api/v2.6/widget/activity.html
- /api/v2.7/widget/activity
- /api/v2.7/widget/activity.html

package: fyne.io/fyne/v2/widget
---


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
