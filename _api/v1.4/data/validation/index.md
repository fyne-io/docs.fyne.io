---
layout: page
tags: [api]
title: Fyne API "validation"
---

# validation
---
```go
import "fyne.io/fyne/data/validation"
```

Package validation provides validation for data inside widgets

## Usage

#### func  NewRegexp

```go
func NewRegexp(regexpstr, reason string) fyne.StringValidator
```
NewRegexp creates a new validator that uses regular expression parsing. The validator will return nil if valid, otherwise returns an error with a reason text.


<div class="since">Since: <code>
1.4</code></div>

#### types
