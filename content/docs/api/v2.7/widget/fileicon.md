---
tags: [api]
title: widget.FileIcon
slug: fileicon

aliases:
- /api/v2.7/widget/fileicon

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type FileIcon struct {
	BaseWidget

	// Deprecated: Selection is now handled externally.
	Selected bool
	URI      fyne.URI
}
```

FileIcon is an adaption of widget.Icon for showing files and folders


<div class="since">Since: <code>
1.4</code></div>

###

```go
func NewFileIcon(uri fyne.URI) *FileIcon
```
NewFileIcon takes a filepath and creates an icon with an overlaid label using the detected mimetype and extension


<div class="since">Since: <code>
1.4</code></div>

###

```go
func (i *FileIcon) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

###

```go
func (i *FileIcon) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

###

```go
func (i *FileIcon) SetSelected(selected bool)
```
SetSelected makes the file look like it is selected.


<div class="deprecated">
Deprecated: Selection is now handled externally.</div>

###

```go
func (i *FileIcon) SetURI(uri fyne.URI)
```
SetURI changes the URI and makes the icon reflect a different file
