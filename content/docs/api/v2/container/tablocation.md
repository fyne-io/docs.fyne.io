---
tags: [api]
title: container.TabLocation
slug: tablocation

aliases:
- /api/v2/container/tablocation.html
- /api/v2.0/container/tablocation
- /api/v2.0/container/tablocation.html
- /api/v2.1/container/tablocation
- /api/v2.1/container/tablocation.html
- /api/v2.2/container/tablocation
- /api/v2.2/container/tablocation.html
- /api/v2.3/container/tablocation
- /api/v2.3/container/tablocation.html
- /api/v2.4/container/tablocation
- /api/v2.4/container/tablocation.html
- /api/v2.5/container/tablocation
- /api/v2.5/container/tablocation.html
- /api/v2.6/container/tablocation
- /api/v2.6/container/tablocation.html
- /api/v2.7/container/tablocation
- /api/v2.7/container/tablocation.html

package: fyne.io/fyne/v2/container
---


---
```go
import "fyne.io/fyne/v2/container"
```

## Usage

#### type TabLocation

```go
type TabLocation int
```

TabLocation is the location where the tabs of a tab container should be rendered


<div class="since">Since: <code>
1.4</code></div>

```go
const (
	TabLocationTop TabLocation = iota
	TabLocationLeading
	TabLocationBottom
	TabLocationTrailing
)
```
TabLocation values
