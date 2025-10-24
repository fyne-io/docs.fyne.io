---
tags: [api]
title: fyne.AnimationCurve
slug: animationcurve

aliases:
- /api/v2.0//animationcurve
- /api/v2.1//animationcurve
- /api/v2.2//animationcurve
- /api/v2.3//animationcurve
- /api/v2.4//animationcurve
- /api/v2.5//animationcurve
- /api/v2.6//animationcurve
- /api/v2.7//animationcurve

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type AnimationCurve

```go
type AnimationCurve func(float32) float32
```

AnimationCurve represents an animation algorithm for calculating the progress through a timeline. Custom animations can be provided by implementing the "func(float32) float32" definition. The input parameter will start at 0.0 when an animation starts and travel up to 1.0 at which point it will end. A linear animation would return the same output value as is passed in.
