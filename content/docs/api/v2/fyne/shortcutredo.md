---
tags: [api]
title: fyne.ShortcutRedo
slug: shortcutredo

aliases:
- /api//shortcutredo
- /api//shortcutredo.html
- /api/v2.0//shortcutredo
- /api/v2.0//shortcutredo.html
- /api/v2.1//shortcutredo
- /api/v2.1//shortcutredo.html
- /api/v2.2//shortcutredo
- /api/v2.2//shortcutredo.html
- /api/v2.3//shortcutredo
- /api/v2.3//shortcutredo.html
- /api/v2.4//shortcutredo
- /api/v2.4//shortcutredo.html
- /api/v2.5//shortcutredo
- /api/v2.5//shortcutredo.html
- /api/v2.6//shortcutredo
- /api/v2.6//shortcutredo.html
- /api/v2.7//shortcutredo
- /api/v2.7//shortcutredo.html

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ShortcutRedo

```go
type ShortcutRedo struct{}
```

ShortcutRedo describes a shortcut redo action.


<div class="since">Since: <code>
2.5</code></div>

#### func (*ShortcutRedo) Key

```go
func (se *ShortcutRedo) Key() KeyName
```
Key returns the [KeyName] for this shortcut.

#### func (*ShortcutRedo) Mod

```go
func (se *ShortcutRedo) Mod() KeyModifier
```
Mod returns the [KeyModifier] for this shortcut.

#### func (*ShortcutRedo) ShortcutName

```go
func (se *ShortcutRedo) ShortcutName() string
```
ShortcutName returns the shortcut name
