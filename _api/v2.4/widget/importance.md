---
layout: page
tags: [api]
title: Fyne API "widget.Importance"
package: fyne.io/fyne/v2/widget
---

# widget.Importance
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
	// Since 2.3
	DangerImportance
	// WarningImportance applies a warning theme to the widget.
	//
	// Since 2.3
	WarningImportance

	// SuccessImportance applies a success theme to the widget.
	//
	// Since 2.4
	SuccessImportance
)
```
