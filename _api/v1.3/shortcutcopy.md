---
layout: page
tags: [api]
title: Fyne API fyne
---

# fyne
---
```go
import "fyne.io/fyne"
```

## Usage

#### type ShortcutCopy

```go
type ShortcutCopy struct {
	Clipboard Clipboard
}
```

ShortcutCopy describes a shortcut copy action.

#### func (*ShortcutCopy) ShortcutName

```go
func (se *ShortcutCopy) ShortcutName() string
```
ShortcutName returns the shortcut name
