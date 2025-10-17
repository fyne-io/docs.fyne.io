---
layout: page
tags: [api]
title: Fyne API "fyne.ShortcutUndo"
package: fyne.io/fyne/v2
---

# fyne.ShortcutUndo
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
