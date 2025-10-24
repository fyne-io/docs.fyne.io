---
tags: [api]
title: fyne.ShortcutPaste
slug: shortcutpaste

aliases:
- /api/v2.0//shortcutpaste
- /api/v2.1//shortcutpaste
- /api/v2.2//shortcutpaste
- /api/v2.3//shortcutpaste
- /api/v2.4//shortcutpaste
- /api/v2.5//shortcutpaste
- /api/v2.6//shortcutpaste
- /api/v2.7//shortcutpaste

package: fyne.io/fyne/v2
---


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

#### func (*ShortcutPaste) Key

```go
func (se *ShortcutPaste) Key() KeyName
```
Key returns the [KeyName] for this shortcut.

#### func (*ShortcutPaste) Mod

```go
func (se *ShortcutPaste) Mod() KeyModifier
```
Mod returns the [KeyModifier] for this shortcut.

#### func (*ShortcutPaste) ShortcutName

```go
func (se *ShortcutPaste) ShortcutName() string
```
ShortcutName returns the shortcut name
