---
tags: [api]
title: widget.Menu
slug: menu

aliases:
- /api/v2.7/widget/menu

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

#

###

```go
type Menu struct {
	BaseWidget

	Items     []fyne.CanvasObject
	OnDismiss func() `json:"-"`
}
```

Menu is a widget for displaying a fyne.Menu.

###

```go
func NewMenu(menu *fyne.Menu) *Menu
```
NewMenu creates a new Menu.

###

```go
func (m *Menu) ActivateLastSubmenu() bool
```
ActivateLastSubmenu finds the last active menu item traversing through the open submenus and activates its submenu if any. It returns `true` if there was a submenu and it was activated and `false` elsewhere. Activating a submenu does show it and activate its first item.

###

```go
func (m *Menu) ActivateNext()
```
ActivateNext activates the menu item following the currently active menu item. If there is no menu item active, it activates the first menu item. If there is no menu item after the current active one, it does nothing. If a submenu is open, it delegates the activation to this submenu.

###

```go
func (m *Menu) ActivatePrevious()
```
ActivatePrevious activates the menu item preceding the currently active menu item. If there is no menu item active, it activates the last menu item. If there is no menu item before the current active one, it does nothing. If a submenu is open, it delegates the activation to this submenu.

###

```go
func (m *Menu) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer returns a new renderer for the menu.

###

```go
func (m *Menu) DeactivateChild()
```
DeactivateChild deactivates the active menu item and hides its submenu if any.

###

```go
func (m *Menu) DeactivateLastSubmenu() bool
```
DeactivateLastSubmenu finds the last open submenu traversing through the open submenus, deactivates its active item and hides it. This also deactivates any submenus of the deactivated submenu. It returns `true` if there was a submenu open and closed and `false` elsewhere.

###

```go
func (m *Menu) Dismiss()
```
Dismiss dismisses the menu by dismissing and hiding the active child and performing OnDismiss.

###

```go
func (m *Menu) MinSize() fyne.Size
```
MinSize returns the minimal size of the menu.

###

```go
func (m *Menu) Refresh()
```
Refresh updates the menu to reflect changes in the data.

###

```go
func (m *Menu) Tapped(*fyne.PointEvent)
```
Tapped catches taps on separators and the menu background. It doesn't perform any action.

###

```go
func (m *Menu) TriggerLast()
```
TriggerLast finds the last active menu item traversing through the open submenus and triggers it.
