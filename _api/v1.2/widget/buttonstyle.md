---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type ButtonStyle

```go
type ButtonStyle int
```

ButtonStyle determines the behaviour and rendering of a button.

```go
const (
	// DefaultButton is the standard button style
	DefaultButton ButtonStyle = iota
	// PrimaryButton that should be more prominent to the user
	PrimaryButton
)
```
