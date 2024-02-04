---
layout: page
title: Creating a Custom Theme

redirect_from:
 - /tutorial/custom-theme
---

Applications are able to load custom themes that can apply small changes to the standard theme or provide a completely unique look. A theme will need to implement the functions of `fyne.Theme` interface, which is defined as follows:

```go
type Theme interface {
	Color(ThemeColorName, ThemeVariant) color.Color
	Font(TextStyle) Resource
	Icon(ThemeIconName) Resource
	Size(ThemeSizeName) float32
}
```

To apply our theme changes we will first define a new type that that implements this interface.

### Define your theme

We start by defining a new type that will be our theme, a simple empty struct will do:

```go
type myTheme struct {}
```

It is a good idea to assert that we implement an interface so that
compile errors are closer to the defining type.

```go
var _ fyne.Theme = (*myTheme)(nil)
```

At this point you could see compile errors because we still need to 
implement the methods, we start with color.

#### Customising colors

The `Color` function defined in the `Theme` interface asks us to define a
named color and also provides a hint for the variant that the user desires (for example `theme.VariantLight` or `theme.VariantDark`). In our theme we will return a custom background color - using a different value for light and dark.

```go
// The color package has to be imported from "image/color".

func (m myTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		if variant == theme.VariantLight {
			return color.White
		}
		return color.Black
	}

	return theme.DefaultTheme().Color(name, variant)
}
```

You will see the last line here references the `theme.DefaultTheme()` to
look up standard values. This allows us to provide custom values and yet
fall back to the standard theme when we don't want to provide our own values.

Of course colors are simpler than resources, we look at that for custom icons.

#### Overriding default icons

Icons (and Fonts) use `fyne.Resource` as values instead of simple types like `int` (for size) or `color.Color` for colors. We can build our own
resource using `fyne.NewStaticResource`, or you could pass in a value
that was created using [resource embedding](https://developer.fyne.io/tutorial/bundle).

```go
func (m myTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	if name == theme.IconNameHome {
		return fyne.NewStaticResource("myHome", homeBytes)
	}
	
	return theme.DefaultTheme().Icon(name)
}
```

As above we return the default theme icon if we don't want to provide
a specific override.


### Load the theme

Before we can load the theme you will need to implement the `Size` and `Font` methods as well. You can just use these empty implementations if
you are happy to use the defaults.

```go
func (m myTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
```

To set the theme for your app you will need to add the following line of code:

```go
app.Settings().SetTheme(&myTheme{})
```

With these changes you can apply your own style, to make small tweaks or
provide a completely custom looking application!
