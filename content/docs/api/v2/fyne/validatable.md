---
tags: [api]
title: fyne.Validatable
slug: validatable

aliases:
- /api//validatable
- /api//validatable.html
- /api/v2.0//validatable
- /api/v2.0//validatable.html
- /api/v2.1//validatable
- /api/v2.1//validatable.html
- /api/v2.2//validatable
- /api/v2.2//validatable.html
- /api/v2.3//validatable
- /api/v2.3//validatable.html
- /api/v2.4//validatable
- /api/v2.4//validatable.html
- /api/v2.5//validatable
- /api/v2.5//validatable.html
- /api/v2.6//validatable
- /api/v2.6//validatable.html
- /api/v2.7//validatable
- /api/v2.7//validatable.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
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
