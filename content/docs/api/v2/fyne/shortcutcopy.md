---
tags: [api]
title: fyne.ShortcutCopy
slug: shortcutcopy

aliases:
- /api//shortcutcopy
- /api//shortcutcopy.html
- /api/v2.0//shortcutcopy
- /api/v2.0//shortcutcopy.html
- /api/v2.1//shortcutcopy
- /api/v2.1//shortcutcopy.html
- /api/v2.2//shortcutcopy
- /api/v2.2//shortcutcopy.html
- /api/v2.3//shortcutcopy
- /api/v2.3//shortcutcopy.html
- /api/v2.4//shortcutcopy
- /api/v2.4//shortcutcopy.html
- /api/v2.5//shortcutcopy
- /api/v2.5//shortcutcopy.html
- /api/v2.6//shortcutcopy
- /api/v2.6//shortcutcopy.html
- /api/v2.7//shortcutcopy
- /api/v2.7//shortcutcopy.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ShortcutCopy

```go
type ShortcutCopy struct {
	Clipboard Clipboard
}
```

ShortcutCopy describes a shortcut copy action.

#### func (*ShortcutCopy) Key

```go
func (se *ShortcutCopy) Key() KeyName
```
Key returns the [KeyName] for this shortcut.

#### func (*ShortcutCopy) Mod

```go
func (se *ShortcutCopy) Mod() KeyModifier
```
Mod returns the [KeyModifier] for this shortcut.

#### func (*ShortcutCopy) ShortcutName

```go
func (se *ShortcutCopy) ShortcutName() string
```
ShortcutName returns the shortcut name
