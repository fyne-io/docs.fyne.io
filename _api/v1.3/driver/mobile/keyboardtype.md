---
layout: page
tags: [api]
title: Fyne API mobile
---

# mobile
---
```go
import "fyne.io/fyne/driver/mobile"
```

## Usage

#### type KeyboardType

```go
type KeyboardType int32
```

KeyboardType represents a type of virtual keyboard

```go
const (
	// DefaultKeyboard is the keyboard with default input style and "return" return key
	DefaultKeyboard KeyboardType = iota
	// SingleLineKeyboard is the keyboard with default input style and "Done" return key
	SingleLineKeyboard
	// NumberKeyboard is the keyboard with number input style and "Done" return key
	NumberKeyboard
)
```
