---
layout: page
tags: [api]
title: Fyne API "fyne.ShortcutHandler"
---

# fyne.ShortcutHandler
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

ShortcutHandler is a default implementation of the shortcut handler for the canvasObject

#### func (*ShortcutHandler) AddShortcut

```go
func (sh *ShortcutHandler) AddShortcut(shortcut Shortcut, handler func(shortcut Shortcut))
```
AddShortcut register an handler to be executed when the shortcut action is triggered

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
