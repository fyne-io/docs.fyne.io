---
tags: [api]
title: binding.URI
slug: uri

aliases:
- /api/v2.7/data/binding/uri

package: fyne.io/fyne/v2/data/binding
---


---
```go
import "fyne.io/fyne/v2/data/binding"
```

#

###

```go
type URI = Item[fyne.URI]
```

URI supports binding a fyne.URI value.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func NewURI() URI
```
NewURI returns a bindable fyne.URI value that is managed internally.


<div class="since">Since: <code>
2.1</code></div>

###

```go
func StringToURI(str String) URI
```
StringToURI creates a binding that connects a String data item to a URI. Changes to the String will be parsed and pushed to the URI if the parse was successful, and setting the URI update the String binding.


<div class="since">Since: <code>
2.1</code></div>
