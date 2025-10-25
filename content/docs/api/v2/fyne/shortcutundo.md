---
tags: [api]
title: fyne.ShortcutUndo
slug: shortcutundo

aliases:
- /api/v2/fyne/shortcutundo.html
- /api/v2.0//shortcutundo
- /api/v2.0//shortcutundo.html
- /api/v2.1//shortcutundo
- /api/v2.1//shortcutundo.html
- /api/v2.2//shortcutundo
- /api/v2.2//shortcutundo.html
- /api/v2.3//shortcutundo
- /api/v2.3//shortcutundo.html
- /api/v2.4//shortcutundo
- /api/v2.4//shortcutundo.html
- /api/v2.5//shortcutundo
- /api/v2.5//shortcutundo.html
- /api/v2.6//shortcutundo
- /api/v2.6//shortcutundo.html
- /api/v2.7//shortcutundo
- /api/v2.7//shortcutundo.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ShortcutUndo

```go
type ShortcutUndo struct{}
```

ShortcutUndo describes a shortcut undo action.


<div class="since">Since: <code>
2.5</code></div>

#### func (*ShortcutUndo) Key

```go
func (se *ShortcutUndo) Key() KeyName
```
Key returns the [KeyName] for this shortcut.

#### func (*ShortcutUndo) Mod

```go
func (se *ShortcutUndo) Mod() KeyModifier
```
Mod returns the [KeyModifier] for this shortcut.

#### func (*ShortcutUndo) ShortcutName

```go
func (se *ShortcutUndo) ShortcutName() string
```
ShortcutName returns the shortcut name
