---
tags: [api]
title: fyne.ShortcutUndo
slug: shortcutundo

aliases:
- /api/v2.7//shortcutundo

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type ShortcutUndo struct{}
```

ShortcutUndo describes a shortcut undo action.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func (se *ShortcutUndo) Key() KeyName
```
Key returns the [KeyName] for this shortcut.

###

```go
func (se *ShortcutUndo) Mod() KeyModifier
```
Mod returns the [KeyModifier] for this shortcut.

###

```go
func (se *ShortcutUndo) ShortcutName() string
```
ShortcutName returns the shortcut name
