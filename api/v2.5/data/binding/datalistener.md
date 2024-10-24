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
package fynedemo

import (
	"testing"

	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

type ViewModel struct {
	IsStarted binding.Bool
}

func (vm *ViewModel) Start() {
	vm.IsStarted.Set(true)
}

func TestStartButtonConnectedToViewModelByAListener(t *testing.T) {

	vm := ViewModel{
		IsStarted: binding.NewBool(),
	}
	vm.IsStarted.Set(false)

	startButton := widget.NewButton("Start", vm.Start)

	vm.IsStarted.AddListener(binding.NewDataListener(
		func() {

			isStarted, _ := vm.IsStarted.Get()
			if isStarted {
				startButton.Disable()
			} else {
				startButton.Enable()
			}

			if startButton.Disabled() && !isStarted {
				t.Errorf("StartButton.Disabled(): `%v` want `%v`", startButton.Disabled(), !startButton.Disabled())
			}
		},
	))

	test.Tap(startButton)
}
```
