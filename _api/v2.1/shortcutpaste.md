---
layout: page
tags: [api]
title: Fyne API "fyne.ShortcutPaste"
---

# fyne.ShortcutPaste
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ShortcutPaste

```go
type ShortcutPaste struct {
	Clipboard Clipboard
}
```

ShortcutPaste describes a shortcut paste action.

#### func (*ShortcutPaste) ShortcutName

```go
func (se *ShortcutPaste) ShortcutName() string
```
ShortcutName returns the shortcut name
