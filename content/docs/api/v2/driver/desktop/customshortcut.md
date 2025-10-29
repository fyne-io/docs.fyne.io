---
tags: [api]
title: desktop.CustomShortcut
slug: customshortcut

aliases:
- /api/driver/desktop/customshortcut
- /api/driver/desktop/customshortcut.html
- /api/v2.0/driver/desktop/customshortcut
- /api/v2.0/driver/desktop/customshortcut.html
- /api/v2.1/driver/desktop/customshortcut
- /api/v2.1/driver/desktop/customshortcut.html
- /api/v2.2/driver/desktop/customshortcut
- /api/v2.2/driver/desktop/customshortcut.html
- /api/v2.3/driver/desktop/customshortcut
- /api/v2.3/driver/desktop/customshortcut.html
- /api/v2.4/driver/desktop/customshortcut
- /api/v2.4/driver/desktop/customshortcut.html
- /api/v2.5/driver/desktop/customshortcut
- /api/v2.5/driver/desktop/customshortcut.html
- /api/v2.6/driver/desktop/customshortcut
- /api/v2.6/driver/desktop/customshortcut.html
- /api/v2.7/driver/desktop/customshortcut
- /api/v2.7/driver/desktop/customshortcut.html

package: fyne.io/fyne/v2/driver/desktop
---


---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type CustomShortcut

```go
type CustomShortcut struct {
	fyne.KeyName
	Modifier fyne.KeyModifier
}
```

CustomShortcut describes a shortcut desktop event.

#### func (*CustomShortcut) Key

```go
func (cs *CustomShortcut) Key() fyne.KeyName
```
Key returns the key name of this shortcut.

#### func (*CustomShortcut) Mod

```go
func (cs *CustomShortcut) Mod() fyne.KeyModifier
```
Mod returns the modifier of this shortcut.

#### func (*CustomShortcut) ShortcutName

```go
func (cs *CustomShortcut) ShortcutName() string
```
ShortcutName returns the shortcut name associated to the event.
