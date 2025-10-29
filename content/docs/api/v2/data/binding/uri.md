---
tags: [api]
title: binding.URI
slug: uri

aliases:
- /api/data/binding/uri
- /api/data/binding/uri.html
- /api/v2.0/data/binding/uri
- /api/v2.0/data/binding/uri.html
- /api/v2.1/data/binding/uri
- /api/v2.1/data/binding/uri.html
- /api/v2.2/data/binding/uri
- /api/v2.2/data/binding/uri.html
- /api/v2.3/data/binding/uri
- /api/v2.3/data/binding/uri.html
- /api/v2.4/data/binding/uri
- /api/v2.4/data/binding/uri.html
- /api/v2.5/data/binding/uri
- /api/v2.5/data/binding/uri.html
- /api/v2.6/data/binding/uri
- /api/v2.6/data/binding/uri.html
- /api/v2.7/data/binding/uri
- /api/v2.7/data/binding/uri.html

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type URI

```go
type URI = Item[fyne.URI]
```

URI supports binding a fyne.URI value.


<div class="since">Since: <code>
2.1</code></div>

#### func  NewURI

```go
func NewURI() URI
```
NewURI returns a bindable fyne.URI value that is managed internally.


<div class="since">Since: <code>
2.1</code></div>

#### func  StringToURI

```go
func StringToURI(str String) URI
```
StringToURI creates a binding that connects a String data item to a URI. Changes to the String will be parsed and pushed to the URI if the parse was successful, and setting the URI update the String binding.


<div class="since">Since: <code>
2.1</code></div>
