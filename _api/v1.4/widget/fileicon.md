---
layout: page
tags: [api]
title: Fyne API "widget.FileIcon"
---

# widget.FileIcon
---
```go
import "fyne.io/fyne/widget"
```

## Usage

#### type FileIcon

```go
type FileIcon struct {
	BaseWidget

	Selected bool
	URI      fyne.URI
}
```

FileIcon is an adaption of widget.Icon for showing files and folders


<div class="since">Since: <code>
1.4</code></div>

#### func  NewFileIcon

```go
func NewFileIcon(uri fyne.URI) *FileIcon
```
NewFileIcon takes a filepath and creates an icon with an overlayed label using the detected mimetype and extension


<div class="since">Since: <code>
1.4</code></div>

#### func (*FileIcon) CreateRenderer

```go
func (i *FileIcon) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*FileIcon) MinSize

```go
func (i *FileIcon) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below

#### func (*FileIcon) SetSelected

```go
func (i *FileIcon) SetSelected(selected bool)
```
SetSelected makes the file look like it is selected

#### func (*FileIcon) SetURI

```go
func (i *FileIcon) SetURI(uri fyne.URI)
```
SetURI changes the URI and makes the icon reflect a different file
