---
layout: page
tags: [api]
title: Fyne API "widget.TabLocation"
---

# widget.TabLocation
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type TabLocation

```go
type TabLocation int
```

TabLocation is the location where the tabs of a tab container should be rendered.


<div class="deprecated">
Deprecated: use container.TabLocation instead.</div>

```go
const (
	// Deprecated: use container.TabLocationTop
	TabLocationTop TabLocation = iota
	// Deprecated: use container.TabLocationLeading
	TabLocationLeading
	// Deprecated: use container.TabLocationBottom
	TabLocationBottom
	// Deprecated: use container.TabLocationTrailing
	TabLocationTrailing
)
```
TabLocation values
