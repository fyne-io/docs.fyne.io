---
title: Theme and Customisation
slug: theme

weight: 11020
---

In this page we answer some common questions about the design of Fyne themes and widgets.

## Customisation

**Q: How can I change the colour of text for a `Label` widget?**

**A:** All of the standard widgets use the current `Theme` definition to set the colour, font and sizes. To make changes to your application consider using a 
[custom theme](/extend/custom-theme).

If your application requires text that is a different colour you can use the `canvas.Text` type instead.
That allows directly setting the colour and size of the text. Be careful when doing this because a user can choose between light or dark theme variations, so you should test with both.

**Q: How can I remove the background colour from my `Entry` widget?**

**A:** The input background is set by the theme `InputBackground` color. You can change that to `color.Transparent` to remove all input background boxes. It is not possible to edit the style of a single input element - the theme API is designed to give a customisable, but consistent, design.

## Theme API

**Q: How can I use my custom theme written before v2.0.0?**

**A:** Over time you should consider updating to use the new theme API. However it is possible to use a simple adapter that was included to allow the usage of old themes during the transition time. You will find `theme.FromLegacy` function that can adapt an old theme instance to the new API.

```go
myTheme := &myOldThemeType{}
updated := theme.FromLegacy(myTheme)
app.Settings().SetTheme(updated)
```

There are no performance penalties when using a theme in this mode, but in a future release this API will be removed.
