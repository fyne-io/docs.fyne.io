---
layout: page
tags: [api]
title: Fyne API "embedded.Driver"
package: fyne.io/fyne/v2/driver/embedded
---

# embedded.Driver
---
```go
import "fyne.io/fyne/v2/driver/embedded"
```

## Usage

#### type Driver

```go
type Driver interface {
	Render(image.Image)
	Run(func())

	ScreenSize() fyne.Size
	Queue() chan Event
}
```

Driver is an embedded driver designed for handling custom hardware. Various standard driver implementations are available in the fyne-x project.


<div class="since">Since: <code>
2.7</code></div>
