---
layout: page
tags: [api]
title: Fyne API "fyne.ShortcutCopy"
package: fyne.io/fyne/v2
---

# fyne.ShortcutCopy
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ShortcutCopy

```go
type ShortcutCopy struct {
	Clipboard Clipboard
}
```

ShortcutCopy describes a shortcut copy action.

#### func (*ShortcutCopy) Key

```go
func (se *ShortcutCopy) Key() KeyName
```
Key returns the [KeyName] for this shortcut.

#### func (*ShortcutCopy) Mod

```go
func (se *ShortcutCopy) Mod() KeyModifier
```
Mod returns the [KeyModifier] for this shortcut.

#### func (*ShortcutCopy) ShortcutName

```go
func (se *ShortcutCopy) ShortcutName() string
```
ShortcutName returns the shortcut name
