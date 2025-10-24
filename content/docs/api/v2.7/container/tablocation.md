---
tags: [api]
title: container.TabLocation
slug: tablocation

aliases:
- /api/v2.7/container/tablocation

package: fyne.io/fyne/v2/container
---


---
```go
import "fyne.io/fyne/v2/container"
```

#

###

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
