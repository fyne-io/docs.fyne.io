---
title: Tree

weight: 6030
--- 

The `Tree` collection widget is like the [List](/collection/list) widget (another of the toolkit's collection widgets) with a multi-level data structure.
Like `List` this is designed to help build really performant
interfaces when lots of data is being presented.
Because of this the widget is not created with all the data embedded, but instead calls out to the data source when needed.

The `Tree` uses callback functions to ask for data when it is required.
There are 4 main callbacks, `ChildUIDs`, `IsBranch`, `CreateNode` and `UpdateNode`.
The ChildUIDs callback is where we communicate the unique ID of each child node to the one requested.
This will be called with the `TreeNodeID` `""` to first get a list of all the IDs that appear inside the root of the tree.
The `IsBranch` callback should return `true` if the node is a branch. This is usually true if the node ID has child nodes - but you can have an empty branch.

**It is crucial that the id of each tree node is unique to the tree.**
For example if you are building a file manager the ID should be the file path rather than its name.

The other two relate to the content templates.

The `CreateNode` callback returns a new template object, just like list, though there is an additional `bool` parameter that is `true` if the node can have child elements (is a branch).
As previously the `UpdateNode` is called to apply data to a cell template.
You should update the content according to the `TreeNodeID` and `isBranch` parameters.


```go
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Table Widget")

	tree := widget.NewTree(
		func(id widget.TreeNodeID) []widget.TreeNodeID {
			switch id {
			case "":
				return []widget.TreeNodeID{"a", "b", "c"}
			case "a":
				return []widget.TreeNodeID{"a1", "a2"}
			}
			return []string{}
		},
		func(id widget.TreeNodeID) bool {
			return id == "" || id == "a"
		},
		func(branch bool) fyne.CanvasObject {
			if branch {
				return widget.NewLabel("Branch template")
			}
			return widget.NewLabel("Leaf template")
		},
		func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
			text := id
			if branch {
				text += " (branch)"
			}
			o.(*widget.Label).SetText(text)
		})

	myWindow.SetContent(tree)
	myWindow.ShowAndRun()
}
```
