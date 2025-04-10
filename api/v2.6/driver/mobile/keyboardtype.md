---
layout: page
tags: [api]
title: Fyne API "mobile.KeyboardType"
package: fyne.io/fyne/v2/driver/mobile
---

# mobile.KeyboardType
---
```go
import "fyne.io/fyne/v2/driver/mobile"
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
	// PasswordKeyboard is used to ensure that text is not leaked to 3rd party keyboard providers
	PasswordKeyboard
)
```
