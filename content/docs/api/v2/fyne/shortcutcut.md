---
tags: [api]
title: fyne.ShortcutCut
slug: shortcutcut

aliases:
- /api/v2/fyne/shortcutcut.html
- /api/v2.0//shortcutcut
- /api/v2.0//shortcutcut.html
- /api/v2.1//shortcutcut
- /api/v2.1//shortcutcut.html
- /api/v2.2//shortcutcut
- /api/v2.2//shortcutcut.html
- /api/v2.3//shortcutcut
- /api/v2.3//shortcutcut.html
- /api/v2.4//shortcutcut
- /api/v2.4//shortcutcut.html
- /api/v2.5//shortcutcut
- /api/v2.5//shortcutcut.html
- /api/v2.6//shortcutcut
- /api/v2.6//shortcutcut.html
- /api/v2.7//shortcutcut
- /api/v2.7//shortcutcut.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ShortcutCut

```go
type ShortcutCut struct {
	Clipboard Clipboard
}
```

ShortcutCut describes a shortcut cut action.

#### func (*ShortcutCut) Key

```go
func (se *ShortcutCut) Key() KeyName
```
Key returns the [KeyName] for this shortcut.

#### func (*ShortcutCut) Mod

```go
func (se *ShortcutCut) Mod() KeyModifier
```
Mod returns the [KeyModifier] for this shortcut.

#### func (*ShortcutCut) ShortcutName

```go
func (se *ShortcutCut) ShortcutName() string
```
ShortcutName returns the shortcut name
