---
layout: page
tags: [api]
title: Fyne API "binding.DataListener"
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
