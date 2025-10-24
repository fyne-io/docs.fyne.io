---
tags: [api]
title: fyne.ShortcutRedo
slug: shortcutredo

aliases:
- /api/v2.7//shortcutredo

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type ShortcutRedo struct{}
```

ShortcutRedo describes a shortcut redo action.


<div class="since">Since: <code>
2.5</code></div>

###

```go
func (se *ShortcutRedo) Key() KeyName
```
Key returns the [KeyName] for this shortcut.

###

```go
func (se *ShortcutRedo) Mod() KeyModifier
```
Mod returns the [KeyModifier] for this shortcut.

###

```go
func (se *ShortcutRedo) ShortcutName() string
```
ShortcutName returns the shortcut name
