---
tags: [api]
title: fyne.ShortcutHandler
slug: shortcuthandler

aliases:
- /api/v2.0//shortcuthandler
- /api/v2.1//shortcuthandler
- /api/v2.2//shortcuthandler
- /api/v2.3//shortcuthandler
- /api/v2.4//shortcuthandler
- /api/v2.5//shortcuthandler
- /api/v2.6//shortcuthandler
- /api/v2.7//shortcuthandler

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type ShortcutHandler

```go
type ShortcutHandler struct {
}
```

ShortcutHandler is a default implementation of the shortcut handler for [CanvasObject].

#### func (*ShortcutHandler) AddShortcut

```go
func (sh *ShortcutHandler) AddShortcut(shortcut Shortcut, handler func(shortcut Shortcut))
```
AddShortcut register a handler to be executed when the shortcut action is triggered

#### func (*ShortcutHandler) RemoveShortcut

```go
func (sh *ShortcutHandler) RemoveShortcut(shortcut Shortcut)
```
RemoveShortcut removes a registered shortcut

#### func (*ShortcutHandler) TypedShortcut

```go
func (sh *ShortcutHandler) TypedShortcut(shortcut Shortcut)
```
TypedShortcut handle the registered shortcut
