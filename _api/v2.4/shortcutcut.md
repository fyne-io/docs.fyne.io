---
layout: page
tags: [api]
title: Fyne API "fyne.ShortcutCut"
package: fyne.io/fyne/v2
---

# fyne.ShortcutCut
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
Key returns the KeyName for this shortcut.


<div class="implements">Implements: <code>
KeyboardShortcut</code></div>

#### func (*ShortcutCut) Mod

```go
func (se *ShortcutCut) Mod() KeyModifier
```
Mod returns the KeyModifier for this shortcut.


<div class="implements">Implements: <code>
KeyboardShortcut</code></div>

#### func (*ShortcutCut) ShortcutName

```go
func (se *ShortcutCut) ShortcutName() string
```
ShortcutName returns the shortcut name
