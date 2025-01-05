---
layout: page
title: Testing Graphical Apps
---

Part of a good test suite is being able to quickly write tests and run them on a regular basis.
Fyne's API is designed to make testing applications easy. By separating component logic from its rendering definition we can load applications without actually displaying them and test the functionality completely.

### Example

We can demonstrate unit testing by extending our [Hello World](/started/hello)
app to include space for users to input their name to be greeted.
We start by updating the user interface to have two elements,
a `Label` for the greeting and an `Entry` for the name input.
We display them, one above another, using `container.NewVBox` (a
vertical box container). The updated user interface code will look as follows:

```go
func makeUI() (*widget.Label, *widget.Entry) {
	return widget.NewLabel("Hello world!"),
		widget.NewEntry()
}

func main() {
	a := app.New()
	w := a.NewWindow("Hello Person")

	w.SetContent(container.NewVBox(makeUI()))
	w.ShowAndRun()
}
```

To test this input behaviour we create a new file (with a name ending
 `_test.go` to mark it as tests) that defines a `TestGreeter` function.

 ```go
package main

import (
	"testing"
)

func TestGreeting(t *testing.T) {
}
```

We can add an initial test that verifies the initial state, to do this
we test the `Text` field of the `Label` that is returned from `makeUI`
and error the test if it is not correct. Add the following code to your test method:

```go
	out, in := makeUI()

	if out.Text != "Hello world!" {
		t.Error("Incorrect initial greeting")
	}
```

This test will pass - next we add to the test to validate the greeter.
We use the Fyne `fyne.io/fyne/v2/test` package which assists in
test scenarios, calling `test.Type` to simulate user input.
The following test code will check that the output updates when the
user's name is input (be sure to add the import as well):

```go
	test.Type(in, "Andy")
	if out.Text != "Hello Andy!" {
		t.Error("Incorrect user greeting")
	}
```

You can run all of these tests using `go test .` - just like any other tests.
Doing so you will now see a failure - because we did not add the greeter logic. Update the `makeUI` function to the following code:

```go
func makeUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Hello world!")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		out.SetText("Hello " + content + "!")
	}
	return out, in
}
```

Doing so you will see that the tests now pass. You can also run the
full application (using `go run .`) and see the greeting update as
you enter a name in the `Entry` field.
Notice also that these tests all run without displaying a window
or stealing your mouse - this is another benefit of the Fyne unit
testing setup.

{% include youtube.html id="AVIQn1areC4" %}
