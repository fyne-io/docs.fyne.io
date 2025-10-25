---
tags: [api]
title: dialog.Dialog
slug: dialog

aliases:
- /api/v2/dialog/dialog.html
- /api/v2.0/dialog/dialog
- /api/v2.0/dialog/dialog.html
- /api/v2.1/dialog/dialog
- /api/v2.1/dialog/dialog.html
- /api/v2.2/dialog/dialog
- /api/v2.2/dialog/dialog.html
- /api/v2.3/dialog/dialog
- /api/v2.3/dialog/dialog.html
- /api/v2.4/dialog/dialog
- /api/v2.4/dialog/dialog.html
- /api/v2.5/dialog/dialog
- /api/v2.5/dialog/dialog.html
- /api/v2.6/dialog/dialog
- /api/v2.6/dialog/dialog.html
- /api/v2.7/dialog/dialog
- /api/v2.7/dialog/dialog.html

package: fyne.io/fyne/v2/dialog
---


---
```go
import "fyne.io/fyne/v2/dialog"
```

## Usage

#### type Dialog

```go
type Dialog interface {
	Show()
	Hide()
	SetDismissText(label string)
	SetOnClosed(closed func())
	Refresh()
	Resize(size fyne.Size)

	// MinSize returns the size that this dialog should not shrink below.
	//
	// Since: 2.1
	MinSize() fyne.Size

	// Dismiss instructs the dialog to close without any affirmative action.
	//
	// Since: 2.6
	Dismiss()
}
```

Dialog is the common API for any dialog window with a single dismiss button

#### func  NewError

```go
func NewError(err error, parent fyne.Window) Dialog
```
NewError creates a dialog over the specified window for an application error. The message is extracted from the provided error (should not be nil). After creation you should call Show().

#### func  NewInformation

```go
func NewInformation(title, message string, parent fyne.Window) Dialog
```
NewInformation creates a dialog over the specified window for user information. The title is used for the dialog window and message is the content. After creation you should call Show().
