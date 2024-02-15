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

#### type ShortcutCut

```go
type ShortcutCut struct {
	Clipboard Clipboard
}
```

ShortcutCut describes a shortcut cut action.

#### func (*ShortcutCut) ShortcutName

```go
func (se *ShortcutCut) ShortcutName() string
```
ShortcutName returns the shortcut name
