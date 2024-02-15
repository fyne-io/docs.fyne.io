---
layout: page
tags: [api]
title: Fyne API "fyne.Validatable"
---

# fyne.Validatable
---
```go
import "fyne.io/fyne"
```

## Usage

#### type Validatable

```go
type Validatable interface {
	Validate() error

	// SetOnValidationChanged is used to set the callback that will be triggered when the validation state changes.
	// The function might be overwritten by a parent that cares about child validation (e.g. widget.Form).
	SetOnValidationChanged(func(error))
}
```

Validatable is an interface for specifying if a widget is validatable.


<div class="since">Since: <code>
1.4</code></div>
