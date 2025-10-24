---
title: Data Binding
slug: data

weight: 7010

aliases:
- /tour/binding
---

Data binding is a powerful new addition to the Fyne toolkit
that was introduced in version `v2.0.0`.
By using data binding we can avoid manually managing many standard
objects like `Label`s, `Button`s and `List`s.

The builtin bindings support many primitive types (like `Int`, `String`, `Float` etc), lists (such as `StringList`, `BoolList`) as well as `Map` and `Struct` bindings. Each of these types can be created using a simple constructor function. For example to create a new string binding with a zero value you can use `binding.NewString()`. You can get or set the value of most data bindings using `Get` and `Set` methods.

It is also possible to bind to an existing value using similar functions who's names start `Bind` and they all accept a pointer to the type bound.
To bind to an existing `int` we could use `binding.BindInt(&myInt)`.
By keeping a reference to a bound value instead of the original 
variable we can configure widgets and functions to respond to any changes automatically.
If you change the external data directly, be sure to call `Reload`()
to ensure that the binding system reads the new value.

```go
package main

import (
	"log"

	"fyne.io/fyne/v2/data/binding"
)

func main() {
	boundString := binding.NewString()
	s, _ := boundString.Get()
	log.Printf("Bound = '%s'", s)

	myInt := 5
	boundInt := binding.BindInt(&myInt)
	i, _ := boundInt.Get()
	log.Printf("Source = %d, bound = %d", myInt, i)
}
```

We start learning about [simple values](/binding/simple) widget
binding next.
