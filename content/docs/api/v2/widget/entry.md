---
tags: [api]
title: widget.Entry
slug: entry

aliases:
- /api/v2.0/widget/entry
- /api/v2.1/widget/entry
- /api/v2.2/widget/entry
- /api/v2.3/widget/entry
- /api/v2.4/widget/entry
- /api/v2.5/widget/entry
- /api/v2.6/widget/entry
- /api/v2.7/widget/entry

package: fyne.io/fyne/v2/widget
---


---
```go
import "fyne.io/fyne/v2/widget"
```

## Usage

#### type Entry

```go
type Entry struct {
	DisableableWidget

	Text string
	// Since: 2.0
	TextStyle   fyne.TextStyle
	PlaceHolder string
	OnChanged   func(string) `json:"-"`
	// Since: 2.0
	OnSubmitted func(string) `json:"-"`
	Password    bool
	MultiLine   bool
	Wrapping    fyne.TextWrap

	// Scroll can be used to turn off the scrolling of our entry when Wrapping is WrapNone.
	//
	// Since: 2.4
	Scroll fyne.ScrollDirection

	// Set a validator that this entry will check against
	// Since: 1.4
	Validator fyne.StringValidator `json:"-"`

	// If true, the Validator runs automatically on render without user interaction.
	// It will reflect any validation errors found or those explicitly set via SetValidationError().
	// Since: 2.7
	AlwaysShowValidationError bool

	CursorRow, CursorColumn int
	OnCursorChanged         func() `json:"-"`

	// Icon is displayed at the outer left of the entry.
	// It is not clickable, but can be used to indicate the purpose of the entry.
	// Since: 2.7
	Icon fyne.Resource `json:"-"`

	// ActionItem is a small item which is displayed at the outer right of the entry (like a password revealer)
	ActionItem fyne.CanvasObject `json:"-"`
}
```

Entry widget allows simple text to be input when focused.

#### func  NewEntry

```go
func NewEntry() *Entry
```
NewEntry creates a new single line entry widget.

#### func  NewEntryWithData

```go
func NewEntryWithData(data binding.String) *Entry
```
NewEntryWithData returns an Entry widget connected to the specified data source.


<div class="since">Since: <code>
2.0</code></div>

#### func  NewMultiLineEntry

```go
func NewMultiLineEntry() *Entry
```
NewMultiLineEntry creates a new entry that allows multiple lines

#### func  NewPasswordEntry

```go
func NewPasswordEntry() *Entry
```
NewPasswordEntry creates a new entry password widget

#### func (*Entry) AcceptsTab

```go
func (e *Entry) AcceptsTab() bool
```
AcceptsTab returns if Entry accepts the Tab key or not.


<div class="since">Since: <code>
2.1</code></div>

#### func (*Entry) Append

```go
func (e *Entry) Append(text string)
```
Append appends the text to the end of the entry.


<div class="since">Since: <code>
2.4</code></div>

#### func (*Entry) Bind

```go
func (e *Entry) Bind(data binding.String)
```
Bind connects the specified data source to this Entry. The current value will be displayed and any changes in the data will cause the widget to update. User interactions with this Entry will set the value into the data source.


<div class="since">Since: <code>
2.0</code></div>

#### func (*Entry) CreateRenderer

```go
func (e *Entry) CreateRenderer() fyne.WidgetRenderer
```
CreateRenderer is a private method to Fyne which links this widget to its renderer

#### func (*Entry) Cursor

```go
func (e *Entry) Cursor() desktop.Cursor
```
Cursor returns the cursor type of this widget

#### func (*Entry) CursorPosition

```go
func (e *Entry) CursorPosition() fyne.Position
```
CursorPosition returns the relative position of this Entry widget's cursor.


<div class="since">Since: <code>
2.7</code></div>

#### func (*Entry) CursorTextOffset

```go
func (e *Entry) CursorTextOffset() (pos int)
```
CursorTextOffset returns how many runes into the source text the cursor is positioned at.


<div class="since">Since: <code>
2.7</code></div>

#### func (*Entry) DoubleTapped

```go
func (e *Entry) DoubleTapped(_ *fyne.PointEvent)
```
DoubleTapped is called when this entry has been double tapped so we should select text below the pointer

#### func (*Entry) DragEnd

```go
func (e *Entry) DragEnd()
```
DragEnd is called at end of a drag event.

#### func (*Entry) Dragged

```go
func (e *Entry) Dragged(d *fyne.DragEvent)
```
Dragged is called when the pointer moves while a button is held down. It updates the selection accordingly.

#### func (*Entry) ExtendBaseWidget

```go
func (e *Entry) ExtendBaseWidget(wid fyne.Widget)
```
ExtendBaseWidget is used by an extending widget to make use of BaseWidget functionality.

#### func (*Entry) FocusGained

```go
func (e *Entry) FocusGained()
```
FocusGained is called when the Entry has been given focus.

#### func (*Entry) FocusLost

```go
func (e *Entry) FocusLost()
```
FocusLost is called when the Entry has had focus removed.

#### func (*Entry) Hide

```go
func (e *Entry) Hide()
```
Hide hides the entry.

#### func (*Entry) KeyDown

```go
func (e *Entry) KeyDown(key *fyne.KeyEvent)
```
KeyDown handler for keypress events - used to store shift modifier state for text selection

#### func (*Entry) KeyUp

```go
func (e *Entry) KeyUp(key *fyne.KeyEvent)
```
KeyUp handler for key release events - used to reset shift modifier state for text selection

#### func (*Entry) Keyboard

```go
func (e *Entry) Keyboard() mobile.KeyboardType
```
Keyboard implements the Keyboardable interface

#### func (*Entry) MinSize

```go
func (e *Entry) MinSize() fyne.Size
```
MinSize returns the size that this widget should not shrink below.

#### func (*Entry) MouseDown

```go
func (e *Entry) MouseDown(m *desktop.MouseEvent)
```
MouseDown called on mouse click, this triggers a mouse click which can move the cursor, update the existing selection (if shift is held), or start a selection dragging operation.

#### func (*Entry) MouseUp

```go
func (e *Entry) MouseUp(m *desktop.MouseEvent)
```
MouseUp called on mouse release If a mouse drag event has completed then check to see if it has resulted in an empty selection, if so, and if a text select key isn't held, then disable selecting

#### func (*Entry) Redo

```go
func (e *Entry) Redo()
```
Redo un-does the last undo action.


<div class="since">Since: <code>
2.5</code></div>

#### func (*Entry) Refresh

```go
func (e *Entry) Refresh()
```

#### func (*Entry) SelectedText

```go
func (e *Entry) SelectedText() string
```
SelectedText returns the text currently selected in this Entry. If there is no selection it will return the empty string.

#### func (*Entry) SetIcon

```go
func (e *Entry) SetIcon(res fyne.Resource)
```
SetIcon sets the leading icon resource for the entry. The icon will be displayed at the outer left of the entry, but is not clickable. This can be used to indicate the purpose of the entry, such as an email or password field.


<div class="since">Since: <code>
2.7</code></div>

#### func (*Entry) SetMinRowsVisible

```go
func (e *Entry) SetMinRowsVisible(count int)
```
SetMinRowsVisible forces a multi-line entry to show `count` number of rows without scrolling. This is not a validation or requirement, it just impacts the minimum visible size. Use this carefully as Fyne apps can run on small screens so you may wish to add a scroll container if this number is high. Default is 3.


<div class="since">Since: <code>
2.2</code></div>

#### func (*Entry) SetOnValidationChanged

```go
func (e *Entry) SetOnValidationChanged(callback func(error))
```
SetOnValidationChanged is intended for parent widgets or containers to hook into the validation. The function might be overwritten by a parent that cares about child validation (e.g. widget.Form).

#### func (*Entry) SetPlaceHolder

```go
func (e *Entry) SetPlaceHolder(text string)
```
SetPlaceHolder sets the text that will be displayed if the entry is otherwise empty

#### func (*Entry) SetText

```go
func (e *Entry) SetText(text string)
```
SetText manually sets the text of the Entry to the given text value. Calling SetText resets all undo history.

#### func (*Entry) SetValidationError

```go
func (e *Entry) SetValidationError(err error)
```
SetValidationError manually updates the validation status until the next input change.

#### func (*Entry) Tapped

```go
func (e *Entry) Tapped(ev *fyne.PointEvent)
```
Tapped is called when this entry has been tapped. We update the cursor position in device-specific callbacks (MouseDown() and TouchDown()).

#### func (*Entry) TappedSecondary

```go
func (e *Entry) TappedSecondary(pe *fyne.PointEvent)
```
TappedSecondary is called when right or alternative tap is invoked.

Opens the PopUpMenu with `Paste` item to paste text from the clipboard.

#### func (*Entry) TouchCancel

```go
func (e *Entry) TouchCancel(*mobile.TouchEvent)
```
TouchCancel is called when this entry gets a touch cancel event on mobile device (app was removed from focus).


<div class="since">Since: <code>
2.1</code></div>

#### func (*Entry) TouchDown

```go
func (e *Entry) TouchDown(ev *mobile.TouchEvent)
```
TouchDown is called when this entry gets a touch down event on mobile device, we ensure we have focus.


<div class="since">Since: <code>
2.1</code></div>

#### func (*Entry) TouchUp

```go
func (e *Entry) TouchUp(*mobile.TouchEvent)
```
TouchUp is called when this entry gets a touch up event on mobile device.


<div class="since">Since: <code>
2.1</code></div>

#### func (*Entry) TypedKey

```go
func (e *Entry) TypedKey(key *fyne.KeyEvent)
```
TypedKey receives key input events when the Entry widget is focused.

#### func (*Entry) TypedRune

```go
func (e *Entry) TypedRune(r rune)
```
TypedRune receives text input events when the Entry widget is focused.

#### func (*Entry) TypedShortcut

```go
func (e *Entry) TypedShortcut(shortcut fyne.Shortcut)
```
TypedShortcut implements the Shortcutable interface

#### func (*Entry) Unbind

```go
func (e *Entry) Unbind()
```
Unbind disconnects any configured data source from this Entry. The current value will remain at the last value of the data source.


<div class="since">Since: <code>
2.0</code></div>

#### func (*Entry) Undo

```go
func (e *Entry) Undo()
```
Undo un-does the last modifying user-action.


<div class="since">Since: <code>
2.5</code></div>

#### func (*Entry) Validate

```go
func (e *Entry) Validate() error
```
Validate validates the current text in the widget.
