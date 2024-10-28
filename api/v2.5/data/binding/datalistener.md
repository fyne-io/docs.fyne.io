---
layout: page
tags: [api]
title: Fyne API "binding.DataListener"
package: fyne.io/fyne/v2/data/binding
---

# binding.DataListener
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type DataListener

```go
type DataListener interface {
	DataChanged()
}
```

DataListener is any object that can register for changes in a bindable DataItem. See NewDataListener to define a new listener using just an inline function.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewDataListener

```go
func NewDataListener(fn func()) DataListener
```
NewDataListener is a helper function that creates a new listener type from a simple callback function.


<div class="since">Since: <code>
2.0</code></div>

### Example

```go
package addlistenerdemo
import (
	"time"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

func ListenerCanConnectAStartButtonWithABooleanState() {
	isRunning := binding.NewBool()

	handleTap := func() { isRunning.Set(true) }
	startButton := widget.NewButton("Start", handleTap)

	isRunning.AddListener(binding.NewDataListener(
		func() {
			isStarted, _ := isRunning.Get()
			if isStarted {
				startButton.Disable()
			} else {
				startButton.Enable()
			}
		},
	))

	isRunning.Set(false)
	time.Sleep(1 * time.Second)
	isRunning.Set(true)
	time.Sleep(1 * time.Second)
}
```
