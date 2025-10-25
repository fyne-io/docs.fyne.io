---
tags: [api]
title: binding.DataListener
slug: datalistener

aliases:
- /api/v2/data/binding/datalistener.html
- /api/v2.0/data/binding/datalistener
- /api/v2.0/data/binding/datalistener.html
- /api/v2.1/data/binding/datalistener
- /api/v2.1/data/binding/datalistener.html
- /api/v2.2/data/binding/datalistener
- /api/v2.2/data/binding/datalistener.html
- /api/v2.3/data/binding/datalistener
- /api/v2.3/data/binding/datalistener.html
- /api/v2.4/data/binding/datalistener
- /api/v2.4/data/binding/datalistener.html
- /api/v2.5/data/binding/datalistener
- /api/v2.5/data/binding/datalistener.html
- /api/v2.6/data/binding/datalistener
- /api/v2.6/data/binding/datalistener.html
- /api/v2.7/data/binding/datalistener
- /api/v2.7/data/binding/datalistener.html

package: fyne.io/fyne/v2/data/binding
---


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
