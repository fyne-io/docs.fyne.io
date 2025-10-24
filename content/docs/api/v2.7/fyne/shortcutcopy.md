---
tags: [api]
title: fyne.ShortcutCopy
slug: shortcutcopy

aliases:
- /api/v2.7//shortcutcopy

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type ShortcutCopy struct {
	Clipboard Clipboard
}
```

ShortcutCopy describes a shortcut copy action.

###

```go
func (se *ShortcutCopy) Key() KeyName
```
Key returns the [KeyName] for this shortcut.

###

```go
func (se *ShortcutCopy) Mod() KeyModifier
```
Mod returns the [KeyModifier] for this shortcut.

###

```go
func (se *ShortcutCopy) ShortcutName() string
```
ShortcutName returns the shortcut name
