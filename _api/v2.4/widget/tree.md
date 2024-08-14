---
layout: page
tags: [api]
title: Fyne API "widget.Tree"
package: fyne.io/fyne/v2/widget
---

# widget.Tree
---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Tree

```go
type Tree struct {
	BaseWidget
	Root TreeNodeID

	ChildUIDs      func(uid TreeNodeID) (c []TreeNodeID)                     `json:"-"` // Return a sorted slice of Children TreeNodeIDs for the given Node TreeNodeID
	CreateNode     func(branch bool) (o fyne.CanvasObject)                   `json:"-"` // Return a CanvasObject that can represent a Branch (if branch is true), or a Leaf (if branch is false)
	IsBranch       func(uid TreeNodeID) (ok bool)                            `json:"-"` // Return true if the given TreeNodeID represents a Branch
	OnBranchClosed func(uid TreeNodeID)                                      `json:"-"` // Called when a Branch is closed
	OnBranchOpened func(uid TreeNodeID)                                      `json:"-"` // Called when a Branch is opened
	OnSelected     func(uid TreeNodeID)                                      `json:"-"` // Called when the Node with the given TreeNodeID is selected.
	OnUnselected   func(uid TreeNodeID)                                      `json:"-"` // Called when the Node with the given TreeNodeID is unselected.
	UpdateNode     func(uid TreeNodeID, branch bool, node fyne.CanvasObject) `json:"-"` // Called to update the given CanvasObject to represent the data at the given TreeNodeID
}
```

Tree widget displays hierarchical data. Each node of the tree must be identified by a Unique TreeNodeID.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewTree

```go
func NewTree(childUIDs func(TreeNodeID) []TreeNodeID, isBranch func(TreeNodeID) bool, create func(bool) fyne.CanvasObject, update func(TreeNodeID, bool, fyne.CanvasObject)) *Tree
```
NewTree returns a new performant tree widget defined by the passed functions. childUIDs returns the child TreeNodeIDs of the given node. isBranch returns true if the given node is a branch, false if it is a leaf. create returns a new template object that can be cached. update is used to apply data at specified data location to the passed template CanvasObject.


<div class="since">Since: <code>
1.4</code></div>

#### func  NewTreeWithData

```go
func NewTreeWithData(data binding.DataTree, createItem func(bool) fyne.CanvasObject, updateItem func(binding.DataItem, bool, fyne.CanvasObject)) *Tree
```
NewTreeWithData creates a new tree widget that will display the contents of the provided data.


<div class="since">Since: <code>
2.4</code></div>

#### func  NewTreeWithStrings

```go
func NewTreeWithStrings(data map[string][]string) (t *Tree)
```
NewTreeWithStrings creates a new tree with the given string map. Data must contain a mapping for the root, which defaults to empty string ("").


<div class="since">Since: <code>
1.4</code></div>

#### func (*Tree) CloseAllBranches

```go
func (t *Tree) CloseAllBranches()
```
CloseAllBranches closes all branches in the tree.

#### func (*Tree) CloseBranch

```go
func (t *Tree) CloseBranch(uid TreeNodeID)
```
CloseBranch closes the branch with the given TreeNodeID.

#### func (*Tree) CreateRenderer

```go
func (t *Tree) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer.

#### func (*Tree) FocusGained

```go
func (t *Tree) FocusGained()
```
FocusGained is called after this Tree has gained focus.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Tree) FocusLost

```go
func (t *Tree) FocusLost()
```
FocusLost is called after this Tree has lost focus.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Tree) IsBranchOpen

```go
func (t *Tree) IsBranchOpen(uid TreeNodeID) bool
```
IsBranchOpen returns true if the branch with the given TreeNodeID is expanded.

#### func (*Tree) MinSize

```go
func (t *Tree) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below.

#### func (*Tree) OpenAllBranches

```go
func (t *Tree) OpenAllBranches()
```
OpenAllBranches opens all branches in the tree.

#### func (*Tree) OpenBranch

```go
func (t *Tree) OpenBranch(uid TreeNodeID)
```
OpenBranch opens the branch with the given TreeNodeID.

#### func (*Tree) RefreshItem

```go
func (t *Tree) RefreshItem(id TreeNodeID)
```
RefreshItem refreshes a single item, specified by the item ID passed in.


<div class="since">Since: <code>
2.4</code></div>

#### func (*Tree) Resize

```go
func (t *Tree) Resize(size fyne.Size)
```
Resize sets a new size for a widget.

#### func (*Tree) ScrollTo

```go
func (t *Tree) ScrollTo(uid TreeNodeID)
```
ScrollTo scrolls to the node with the given id.

Since 2.1

#### func (*Tree) ScrollToBottom

```go
func (t *Tree) ScrollToBottom()
```
ScrollToBottom scrolls to the bottom of the tree.

Since 2.1

#### func (*Tree) ScrollToTop

```go
func (t *Tree) ScrollToTop()
```
ScrollToTop scrolls to the top of the tree.

Since 2.1

#### func (*Tree) Select

```go
func (t *Tree) Select(uid TreeNodeID)
```
Select marks the specified node to be selected.

#### func (*Tree) ToggleBranch

```go
func (t *Tree) ToggleBranch(uid string)
```
ToggleBranch flips the state of the branch with the given TreeNodeID.

#### func (*Tree) TypedKey

```go
func (t *Tree) TypedKey(event *fyne.KeyEvent)
```
TypedKey is called if a key event happens while this Tree is focused.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Tree) TypedRune

```go
func (t *Tree) TypedRune(_ rune)
```
TypedRune is called if a text event happens while this Tree is focused.


<div class="implements">Implements: <code>
fyne.Focusable</code></div>

#### func (*Tree) Unselect

```go
func (t *Tree) Unselect(uid TreeNodeID)
```
Unselect marks the specified node to be not selected.

#### func (*Tree) UnselectAll

```go
func (t *Tree) UnselectAll()
```
UnselectAll sets all nodes to be not selected.


<div class="since">Since: <code>
2.1</code></div>
