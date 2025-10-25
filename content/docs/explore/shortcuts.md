---
layout: page
title: Adding Shortcuts to an App
slug: shortcuts

weight: 2070

aliases:
- /explore/shortcuts.html
- /tutorial/shortcuts
---

Shortcuts are common tasks that can be triggered by keyboard combinations or context menus. Shortcuts, much like keyboard events, can be attached to a focused element or registered on the `Canvas` to always be available in a `Window`.

## Registering with a Canvas

There are many standard shortcuts defined (such as `fyne.ShortcutCopy`) which are connected to standard keyboard shortcuts and right-click menus. The first step to adding a new `Shortcut` is to define the shortcut. For most uses this will be a keyboard triggered shortcut, which is a desktop extension. To do this we use `desktop.CustomShortcut`, for example to use the Tab key and Control modifier you might do the following:

```go
	ctrlTab    := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}
	ctrlAltTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl | fyne.KeyModifierAlt}
```

Notice that this shortcut can be re-used so you could attach it to menus or other items as well. For this example we want it to be always available, so we register it with our window's `Canvas` as follows:

```go
	ctrlTab := &desktop.CustomShortcut{KeyName: fyne.KeyTab, Modifier: fyne.KeyModifierControl}
	w.Canvas().AddShortcut(ctrlTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Tab")
	})
	w.Canvas().AddShortcut(ctrlAltTab, func(shortcut fyne.Shortcut) {
		log.Println("We tapped Ctrl+Alt+Tab")
	})
```

As you can see there are two parts to registering a shortcut in this way - passing the shortcut definition and also a callback function. If the user types the keyboard shortcut then the function will be called and the output printed.

Shortcuts only work in combination with modifier keys. In order to react to keyboard input without modifiers use `canvas.OnTypedRune` or `canvas.OnTypedKey`.

## Adding shortcuts to an Entry

It can also be helpful to have a shortcut apply only when the current item is focused. This approach can be used for any focusable widget, and is managed by extending that widget and adding a `TypedShortcut` handler. This is much like adding key handlers, except the value passed in will be a `fyne.Shortcut`.

```go
type myEntry struct {
	widget.Entry
}

func (m *myEntry) TypedShortcut(s fyne.Shortcut) {
	if _, ok := s.(*desktop.CustomShortcut); !ok {
		m.Entry.TypedShortcut(s)
		return
	}

	log.Println("Shortcut typed:", s)
}
```

From the excerpt above you can see how a `TypedShortcut` handler might be implemented. Inside this function you should check whether the shortcut is of the custom type used earlier. If the shortcut is a standard one it's a good idea to call the original shortcut handler (if the widget had one).
With those checks done you can compare the shortcut with the various types you are handling (if there are multiple).
