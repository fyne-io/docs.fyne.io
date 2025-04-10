---
layout: page
tags: [api]
title: Fyne API "fyne.KeyboardShortcut"
package: fyne.io/fyne/v2
---

# fyne.KeyboardShortcut
---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type KeyboardShortcut

```go
type KeyboardShortcut interface {
	Shortcut
	Key() KeyName
	Mod() KeyModifier
}
```

KeyboardShortcut describes a shortcut meant to be triggered by a keyboard action.
