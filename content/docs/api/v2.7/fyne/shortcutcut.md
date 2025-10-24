---
tags: [api]
title: fyne.ShortcutCut
slug: shortcutcut

aliases:
- /api/v2.7//shortcutcut

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type ShortcutCut struct {
	Clipboard Clipboard
}
```

ShortcutCut describes a shortcut cut action.

###

```go
func (se *ShortcutCut) Key() KeyName
```
Key returns the [KeyName] for this shortcut.

###

```go
func (se *ShortcutCut) Mod() KeyModifier
```
Mod returns the [KeyModifier] for this shortcut.

###

```go
func (se *ShortcutCut) ShortcutName() string
```
ShortcutName returns the shortcut name
