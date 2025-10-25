---
tags: [api]
title: mobile.KeyboardType
slug: keyboardtype

aliases:
- /api/v2/driver/mobile/keyboardtype.html
- /api/v2.0/driver/mobile/keyboardtype
- /api/v2.0/driver/mobile/keyboardtype.html
- /api/v2.1/driver/mobile/keyboardtype
- /api/v2.1/driver/mobile/keyboardtype.html
- /api/v2.2/driver/mobile/keyboardtype
- /api/v2.2/driver/mobile/keyboardtype.html
- /api/v2.3/driver/mobile/keyboardtype
- /api/v2.3/driver/mobile/keyboardtype.html
- /api/v2.4/driver/mobile/keyboardtype
- /api/v2.4/driver/mobile/keyboardtype.html
- /api/v2.5/driver/mobile/keyboardtype
- /api/v2.5/driver/mobile/keyboardtype.html
- /api/v2.6/driver/mobile/keyboardtype
- /api/v2.6/driver/mobile/keyboardtype.html
- /api/v2.7/driver/mobile/keyboardtype
- /api/v2.7/driver/mobile/keyboardtype.html

package: fyne.io/fyne/v2/driver/mobile
---


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
