---
tags: [api]
title: fyne.ShortcutSelectAll
slug: shortcutselectall

aliases:
- /api/v2.7//shortcutselectall

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type ShortcutSelectAll struct{}
```

ShortcutSelectAll describes a shortcut selectAll action.

###

```go
func (se *ShortcutSelectAll) Key() KeyName
```
Key returns the [KeyName] for this shortcut.

###

```go
func (se *ShortcutSelectAll) Mod() KeyModifier
```
Mod returns the [KeyModifier] for this shortcut.

###

```go
func (se *ShortcutSelectAll) ShortcutName() string
```
ShortcutName returns the shortcut name
