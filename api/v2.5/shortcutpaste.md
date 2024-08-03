---
layout: page
tags: [api]
title: Fyne API "fyne.ShortcutPaste"
package: fyne.io/fyne/v2
---

# fyne.ShortcutPaste
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ShortcutPaste

```go
type ShortcutPaste struct {
	Clipboard Clipboard
}
```

ShortcutPaste describes a shortcut paste action.

#### func (*ShortcutPaste) Key

```go
func (se *ShortcutPaste) Key() KeyName
```
Key returns the KeyName for this shortcut.


<div class="implements">Implements: <code>
KeyboardShortcut</code></div>

#### func (*ShortcutPaste) Mod

```go
func (se *ShortcutPaste) Mod() KeyModifier
```
Mod returns the KeyModifier for this shortcut.


<div class="implements">Implements: <code>
KeyboardShortcut</code></div>

#### func (*ShortcutPaste) ShortcutName

```go
func (se *ShortcutPaste) ShortcutName() string
```
ShortcutName returns the shortcut name
