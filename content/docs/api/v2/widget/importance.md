---
tags: [api]
title: widget.Importance
slug: importance

aliases:
- /api/v2/widget/importance.html
- /api/v2.0/widget/importance
- /api/v2.0/widget/importance.html
- /api/v2.1/widget/importance
- /api/v2.1/widget/importance.html
- /api/v2.2/widget/importance
- /api/v2.2/widget/importance.html
- /api/v2.3/widget/importance
- /api/v2.3/widget/importance.html
- /api/v2.4/widget/importance
- /api/v2.4/widget/importance.html
- /api/v2.5/widget/importance
- /api/v2.5/widget/importance.html
- /api/v2.6/widget/importance
- /api/v2.6/widget/importance.html
- /api/v2.7/widget/importance
- /api/v2.7/widget/importance.html

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Importance

```go
type Importance int
```

Importance represents how prominent the widget should appear


<div class="since">Since: <code>
2.4</code></div>

```go
const (
	// MediumImportance applies a standard appearance.
	MediumImportance Importance = iota
	// HighImportance applies a prominent appearance.
	HighImportance
	// LowImportance applies a subtle appearance.
	LowImportance

	// DangerImportance applies an error theme to the widget.
	//
	// Since: 2.3
	DangerImportance
	// WarningImportance applies a warning theme to the widget.
	//
	// Since: 2.3
	WarningImportance

	// SuccessImportance applies a success theme to the widget.
	//
	// Since: 2.4
	SuccessImportance
)
```
