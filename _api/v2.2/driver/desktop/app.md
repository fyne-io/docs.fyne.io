---
layout: page
tags: [api]
title: Fyne API "desktop.App"
---

# desktop.App
---
```go
import "fyne.io/fyne/v2/driver/desktop"
```

## Usage

#### type App

```go
type App interface {
	SetSystemTrayMenu(menu *fyne.Menu)
	SetSystemTrayIcon(icon fyne.Resource)
}
```

App defines the desktop specific extensions to a fyne.App.


<div class="since">Since: <code>
2.2</code></div>
