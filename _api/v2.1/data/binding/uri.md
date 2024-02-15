---
layout: page
tags: [api]
title: Fyne API "binding.URI"
---

# binding.URI
---
```go
import "fyne.io/fyne/v2/data/binding"
```

## Usage

#### type URI

```go
type URI interface {
	DataItem
	Get() (fyne.URI, error)
	Set(fyne.URI) error
}
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
