---
tags: [api]
title: fyne.Animation
slug: animation

aliases:
- /api/v2/fyne/animation.html
- /api/v2.0//animation
- /api/v2.0//animation.html
- /api/v2.1//animation
- /api/v2.1//animation.html
- /api/v2.2//animation
- /api/v2.2//animation.html
- /api/v2.3//animation
- /api/v2.3//animation.html
- /api/v2.4//animation
- /api/v2.4//animation.html
- /api/v2.5//animation
- /api/v2.5//animation.html
- /api/v2.6//animation
- /api/v2.6//animation.html
- /api/v2.7//animation
- /api/v2.7//animation.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Animation

```go
type Animation struct {
	AutoReverse bool
	Curve       AnimationCurve
	Duration    time.Duration
	RepeatCount int
	Tick        func(float32)
}
```

Animation represents an animated element within a Fyne canvas. These animations may control individual objects or entire scenes.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewAnimation

```go
func NewAnimation(d time.Duration, fn func(float32)) *Animation
```
NewAnimation creates a very basic animation where the callback function will be called for every rendered frame between [time.Now] and the specified duration. The callback values start at 0.0 and will be 1.0 when the animation completes.


<div class="since">Since: <code>
2.0</code></div>

#### func (*Animation) Start

```go
func (a *Animation) Start()
```
Start registers the animation with the application run-loop and starts its execution.

#### func (*Animation) Stop

```go
func (a *Animation) Stop()
```
Stop will end this animation and remove it from the run-loop.
