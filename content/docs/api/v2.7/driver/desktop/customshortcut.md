---
tags: [api]
title: desktop.CustomShortcut
slug: customshortcut

aliases:
- /api/v2.7/driver/desktop/customshortcut

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

#

###

```go
type CustomShortcut struct {
	fyne.KeyName
	Modifier fyne.KeyModifier
}
```

CustomShortcut describes a shortcut desktop event.

###

```go
func (cs *CustomShortcut) Key() fyne.KeyName
```
Key returns the key name of this shortcut.

###

```go
func (cs *CustomShortcut) Mod() fyne.KeyModifier
```
Mod returns the modifier of this shortcut.

###

```go
func (cs *CustomShortcut) ShortcutName() string
```
ShortcutName returns the shortcut name associated to the event.
