---
layout: page
tags: [api]
title: Fyne API "fyne.ShortcutRedo"
package: fyne.io/fyne/v2
---

# fyne.ShortcutRedo
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


<div class="implements">Implements: <code>
[KeyboardShortcut]</code></div>

#### func (*ShortcutRedo) Mod

```go
func (se *ShortcutRedo) Mod() KeyModifier
```
Mod returns the [KeyModifier] for this shortcut.


<div class="implements">Implements: <code>
[KeyboardShortcut]</code></div>

#### func (*ShortcutRedo) ShortcutName

```go
func (se *ShortcutRedo) ShortcutName() string
```
ShortcutName returns the shortcut name
