---
tags: [api]
title: fyne.ShortcutHandler
slug: shortcuthandler

aliases:
- /api/v2.7//shortcuthandler

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type ShortcutHandler struct {
}
```

ShortcutHandler is a default implementation of the shortcut handler for [CanvasObject].

###

```go
func (sh *ShortcutHandler) AddShortcut(shortcut Shortcut, handler func(shortcut Shortcut))
```
AddShortcut register a handler to be executed when the shortcut action is triggered

###

```go
func (sh *ShortcutHandler) RemoveShortcut(shortcut Shortcut)
```
RemoveShortcut removes a registered shortcut

###

```go
func (sh *ShortcutHandler) TypedShortcut(shortcut Shortcut)
```
TypedShortcut handle the registered shortcut
