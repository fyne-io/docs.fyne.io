---
tags: [api]
title: fyne.ShortcutPaste
slug: shortcutpaste

aliases:
- /api/v2.7//shortcutpaste

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type ShortcutPaste struct {
	Clipboard Clipboard
}
```

ShortcutPaste describes a shortcut paste action.

###

```go
func (se *ShortcutPaste) Key() KeyName
```
Key returns the [KeyName] for this shortcut.

###

```go
func (se *ShortcutPaste) Mod() KeyModifier
```
Mod returns the [KeyModifier] for this shortcut.

###

```go
func (se *ShortcutPaste) ShortcutName() string
```
ShortcutName returns the shortcut name
