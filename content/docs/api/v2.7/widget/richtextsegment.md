---
tags: [api]
title: widget.RichTextSegment
slug: richtextsegment

aliases:
- /api/v2.7/widget/richtextsegment

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type RichTextSegment interface {
	Inline() bool
	Textual() string
	Update(fyne.CanvasObject)
	Visual() fyne.CanvasObject

	Select(pos1, pos2 fyne.Position)
	SelectedText() string
	Unselect()
}
```

RichTextSegment describes any element that can be rendered in a RichText widget.


<div class="since">Since: <code>
2.1</code></div>
