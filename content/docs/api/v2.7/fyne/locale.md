---
tags: [api]
title: fyne.Locale
slug: locale

aliases:
- /api/v2.7//locale

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

#

###

```go
type Locale string
```

Locale represents a user's locale (language, region and script)


<div class="since">Since: <code>
2.5</code></div>

###

```go
func (l Locale) LanguageString() string
```
LanguageString returns a version of the local without the script portion. For example "en" or "fr-FR".

###

```go
func (l Locale) String() string
```
String returns the complete locale as a standard string.
