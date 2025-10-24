---
tags: [api]
title: fyne.ShortcutSelectAll
slug: shortcutselectall

aliases:
- /api/v2.0//shortcutselectall
- /api/v2.1//shortcutselectall
- /api/v2.2//shortcutselectall
- /api/v2.3//shortcutselectall
- /api/v2.4//shortcutselectall
- /api/v2.5//shortcutselectall
- /api/v2.6//shortcutselectall
- /api/v2.7//shortcutselectall

package: fyne.io/fyne/v2
---


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
Key returns the [KeyName] for this shortcut.

#### func (*ShortcutSelectAll) Mod

```go
func (se *ShortcutSelectAll) Mod() KeyModifier
```
Mod returns the [KeyModifier] for this shortcut.

#### func (*ShortcutSelectAll) ShortcutName

```go
func (se *ShortcutSelectAll) ShortcutName() string
```
ShortcutName returns the shortcut name
