---
title: Numerical Entry

weight: 8060

aliases:
- /extend/numerical-entry.html
 - /tutorial/keypress-on-entry.html
 - /tutorial/numerical-entry
---

In the traditional sense, GUI programs have used callbacks to customize actions for widgets. Fyne does not expose inserting custom callbacks to capture events on widgets, but it does not need to. The Go language is plenty extensible to make this work.

Instead we can simply use Type Embedding and extend the widget to only make it possible to enter numerical values.

First create a new type struct, we will call it `numericalEntry`.

```go
type numericalEntry struct {
    widget.Entry
}
```

As mentioned in [Extending existing widgets](https://developer.fyne.io/tutorial/extending-widgets), we follow good practice and create a constructor function that extends the `BaseWidget`.

```go
func newNumericalEntry() *numericalEntry {
    entry := &numericalEntry{}
    entry.ExtendBaseWidget(entry)
    return entry
}
```

Now we need to make the entry accept only numbers. This can be done by overriding the `TypedRune(rune)` method that's part of the `fyne.Focusable` interface.
This will allow us to intercept the standard handling of runes received from key presses and only pass through those that we want.
Inside this method, we will use a conditional to check if the rune matches any of the numbers between zero and nine. If they do, we delegate it to the standard `TypedRune(rune)` method of the embedded entry. If they do not, we just ignore the inputs.
This implementation will only allow integers to be entered, but can easily be extended to check for other keys in the future if necessary.

```go
func (e *numericalEntry) TypedRune(r rune) {
	if r >= '0' && r <= '9' {
		e.Entry.TypedRune(r)
	}
}
```

If we want to update the implementation to allow for decimal numbers as well, we can simply add `.` and `,` to the list of allowed runes (some languages use commas over dots for decimal notations).

```go
func (e *numericalEntry) TypedRune(r rune) {
	if (r >= '0' && r <= '9') || r == '.' || r == ',' {
			e.Entry.TypedRune(r)
	}
}
```

With this, the entry now only allows the user to enter numerical values when keys are pressed. However, the paste shortcut will still allow text to be entered.
To fix this, we can overwrite the `TypedShortcut(fyne.Shortcut)` method that is part of the `fyne.Shortcutable` interface.
First we need to do a type assertion to check if the given shortcut is of the type `*fyne.ShortcutPaste`. If it is not, we can just delegate the shortcut back to the embedded entry.
If it is, we check if the clipboard content is numerical, by using `strconv.ParseFloat()` (if you want to only allow integers, `strconv.Atoi()` will be just fine), and then delegating the shortcut back to the embedded entry if the clipboard content could be parsed without errors.

```go
func (e *numericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}
```

As a bonus, we can also make sure that mobile operating systems open the numerical keyboard instead of the default keyboard. This can be done by by first importing the `fyne.io/fyne/v2/driver/mobile` package and overwriting the `Keyboard() mobile.KeyboardType` method that is part of the `m̀obile.Keyboardable` interface. Inside the function, we then simply return the `mobile.NumberKeyboard` type.

```go
func (e *numericalEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}
```

In the end, the resulting code could look something like this:

```go
package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type numericalEntry struct {
	widget.Entry
}

func newNumericalEntry() *numericalEntry {
	entry := &numericalEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *numericalEntry) TypedRune(r rune) {
	if (r >= '0' && r <= '9') || r == '.' || r == ',' {
		e.Entry.TypedRune(r)
	}
}

func (e *numericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func (e *numericalEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}

func main() {
	a := app.New()
	w := a.NewWindow("Numerical")

	entry := newNumericalEntry()

	w.SetContent(entry)
	w.ShowAndRun()
}
```
