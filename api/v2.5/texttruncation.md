---
layout: page
tags: [api]
title: Fyne API "fyne.TextTruncation"
package: fyne.io/fyne/v2
---

# fyne.TextTruncation
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
