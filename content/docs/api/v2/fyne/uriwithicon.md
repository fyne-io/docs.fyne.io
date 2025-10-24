---
tags: [api]
title: fyne.URIWithIcon
slug: uriwithicon

aliases:
- /api/v2.0//uriwithicon
- /api/v2.1//uriwithicon
- /api/v2.2//uriwithicon
- /api/v2.3//uriwithicon
- /api/v2.4//uriwithicon
- /api/v2.5//uriwithicon
- /api/v2.6//uriwithicon
- /api/v2.7//uriwithicon

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type URIWithIcon

```go
type URIWithIcon interface {
	URI

	Icon() Resource
}
```

URIWithIcon describes a [URI] that should be rendered with a certain icon in file browsers.


<div class="since">Since: <code>
2.5</code></div>
