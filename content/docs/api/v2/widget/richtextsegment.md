---
tags: [api]
title: widget.RichTextSegment
slug: richtextsegment

aliases:
- /api/widget/richtextsegment
- /api/widget/richtextsegment.html
- /api/v2.0/widget/richtextsegment
- /api/v2.0/widget/richtextsegment.html
- /api/v2.1/widget/richtextsegment
- /api/v2.1/widget/richtextsegment.html
- /api/v2.2/widget/richtextsegment
- /api/v2.2/widget/richtextsegment.html
- /api/v2.3/widget/richtextsegment
- /api/v2.3/widget/richtextsegment.html
- /api/v2.4/widget/richtextsegment
- /api/v2.4/widget/richtextsegment.html
- /api/v2.5/widget/richtextsegment
- /api/v2.5/widget/richtextsegment.html
- /api/v2.6/widget/richtextsegment
- /api/v2.6/widget/richtextsegment.html
- /api/v2.7/widget/richtextsegment
- /api/v2.7/widget/richtextsegment.html

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type RichTextSegment

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
