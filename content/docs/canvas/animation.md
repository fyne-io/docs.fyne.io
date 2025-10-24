---
title: Animation

weight: 3080

since: 2.0
---

Fyne includes an animation framework that allows you to smoothly transition canvas properties from
one value to another over time. An animation can contain any code which means that any types of object
attributes can be managed, however there are builtin animations for size, position and color.

Animations are normally created using the builtin helpers of the canvas package, such as `NewSizeAnimation`,
and calling `Start()` on the created animation. You can set animations to repeat or auto reverse, as we will see below.

Let us look first at a colour animation which gradually changes the fill colour of a `Rectangle`.
In the following code sample we set an rectangle to be set as the content of a window, as we have done in earlier code
samples. The big difference is the animation that we start just before showing the window.
The animation is created using `NewColorRGBAAnimation` which will transition the colour channels from the defined
`red` state through to `blue` and it will take 2 seconds (the specified duration) to do so.

```go
package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	obj := canvas.NewRectangle(color.Black)
	obj.Resize(fyne.NewSize(50, 50))
	w.SetContent(container.NewWithoutLayout(obj))

	red := color.NRGBA{R:0xff, A:0xff}
	blue := color.NRGBA{B:0xff, A:0xff}
	canvas.NewColorRGBAAnimation(red, blue, time.Second*2, func(c color.Color) {
		obj.FillColor = c
		canvas.Refresh(obj)
	}).Start()

	w.Resize(fyne.NewSize(250, 50))
	w.SetPadded(false)
	w.ShowAndRun()
}
```

It is also possible to animate multiple properties at the same time. If you look carefully you will see that we
added the rectangle into a container without layout - this means that we can manually move or resize the object.
Let's add a new position animation that will move the `Rectangle` across the window, and automatically reverse as well.

```go
move := canvas.NewPositionAnimation(fyne.NewPos(0, 0), fyne.NewPos(200, 0), time.Second, obj.Move)
move.AutoReverse = true
move.Start()
```

Because the `Move()` function of `CanvasObject` expects a `fyne.Position` argument, and so does the position
animation callback, we can simply pass the method name instead of creating a new function
If you add the code above just under the first animation you will see that the object moves across the window
at the same time as it changes colour!
