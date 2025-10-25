---
tags: [api]
title: fyne.TextTruncation
slug: texttruncation

aliases:
- /api/v2/fyne/texttruncation.html
- /api/v2.0//texttruncation
- /api/v2.0//texttruncation.html
- /api/v2.1//texttruncation
- /api/v2.1//texttruncation.html
- /api/v2.2//texttruncation
- /api/v2.2//texttruncation.html
- /api/v2.3//texttruncation
- /api/v2.3//texttruncation.html
- /api/v2.4//texttruncation
- /api/v2.4//texttruncation.html
- /api/v2.5//texttruncation
- /api/v2.5//texttruncation.html
- /api/v2.6//texttruncation
- /api/v2.6//texttruncation.html
- /api/v2.7//texttruncation
- /api/v2.7//texttruncation.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type TextTruncation

```go
type TextTruncation int
```

TextTruncation controls how a `Label` or `Entry` will truncate its text. The default value is `TextTruncateOff` which will not truncate.


<div class="since">Since: <code>
2.4</code></div>

```go
const (
	// TextTruncateOff means no truncation will be applied, it is the default.
	// This means that the minimum size of a text block will be the space required to display it fully.
	TextTruncateOff TextTruncation = iota
	// TextTruncateClip will truncate text when it reaches the end of the available space.
	TextTruncateClip
	// TextTruncateEllipsis is like regular truncation except that an ellipses (…) will be inserted
	// wherever text has been shortened to fit.
	//
	// Since: 2.4
	TextTruncateEllipsis
)
```
