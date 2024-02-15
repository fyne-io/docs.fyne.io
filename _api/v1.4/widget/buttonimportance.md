---
layout: page
tags: [api]
title: Fyne API "widget.ButtonImportance"
---

# widget.ButtonImportance
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type ButtonImportance

```go
type ButtonImportance int
```

ButtonImportance represents how prominent the button should appear


<div class="since">Since: <code>
1.4</code></div>

```go
const (
	// MediumImportance applies a standard appearance.
	MediumImportance ButtonImportance = iota
	// HighImportance applies a prominent appearance.
	HighImportance
	// LowImportance applies a subtle appearance.
	LowImportance
)
```
