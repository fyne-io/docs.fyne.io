---
layout: page
tags: [api]
title: Fyne API widget
---

# widget
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type TabLocation

```go
type TabLocation int
```

TabLocation is the location where the tabs of a tab container should be rendered

```go
const (
	TabLocationTop TabLocation = iota
	TabLocationLeading
	TabLocationBottom
	TabLocationTrailing
)
```
TabLocation values
