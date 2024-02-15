---
layout: page
tags: [api]
title: Fyne API "fyne.ShortcutSelectAll"
---

# fyne.ShortcutSelectAll
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ShortcutSelectAll

```go
type ShortcutSelectAll struct{}
```

ShortcutSelectAll describes a shortcut selectAll action.

#### func (*ShortcutSelectAll) Key

```go
func (se *ShortcutSelectAll) Key() KeyName
```
Key returns the KeyName for this shortcut.


<div class="implements">Implements: <code>
KeyboardShortcut</code></div>

#### func (*ShortcutSelectAll) Mod

```go
func (se *ShortcutSelectAll) Mod() KeyModifier
```
Mod returns the KeyModifier for this shortcut.


<div class="implements">Implements: <code>
KeyboardShortcut</code></div>

#### func (*ShortcutSelectAll) ShortcutName

```go
func (se *ShortcutSelectAll) ShortcutName() string
```
ShortcutName returns the shortcut name
