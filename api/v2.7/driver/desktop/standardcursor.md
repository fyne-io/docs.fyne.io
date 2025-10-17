---
layout: page
tags: [api]
title: Fyne API "desktop.StandardCursor"
package: fyne.io/fyne/v2/driver/desktop
---

# desktop.StandardCursor
---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type StandardCursor

```go
type StandardCursor int
```

StandardCursor represents a standard Fyne cursor. These values were previously of type `fyne.Cursor`.


<div class="since">Since: <code>
2.0</code></div>

```go
const (
	// DefaultCursor is the default cursor typically an arrow
	DefaultCursor StandardCursor = iota
	// TextCursor is the cursor often used to indicate text selection
	TextCursor
	// CrosshairCursor is the cursor often used to indicate bitmaps
	CrosshairCursor
	// PointerCursor is the cursor often used to indicate a link
	PointerCursor
	// HResizeCursor is the cursor often used to indicate horizontal resize
	HResizeCursor
	// VResizeCursor is the cursor often used to indicate vertical resize
	VResizeCursor
	// HiddenCursor will cause the cursor to not be shown
	HiddenCursor
)
```

#### func (StandardCursor) Image

```go
func (d StandardCursor) Image() (image.Image, int, int)
```
Image is not used for any of the StandardCursor types.


<div class="since">Since: <code>
2.0</code></div>
