---
tags: [api]
title: fyne.Locale
slug: locale

aliases:
- /api/v2.0//locale
- /api/v2.1//locale
- /api/v2.2//locale
- /api/v2.3//locale
- /api/v2.4//locale
- /api/v2.5//locale
- /api/v2.6//locale
- /api/v2.7//locale

package: fyne.io/fyne/v2
---


---
```go
import "fyne.io/fyne/v2"
```

## Usage

#### type Locale

```go
type Locale string
```

Locale represents a user's locale (language, region and script)


<div class="since">Since: <code>
2.5</code></div>

#### func (Locale) LanguageString

```go
func (l Locale) LanguageString() string
```
LanguageString returns a version of the local without the script portion. For example "en" or "fr-FR".

#### func (Locale) String

```go
func (l Locale) String() string
```
String returns the complete locale as a standard string.
