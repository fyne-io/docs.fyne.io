---
layout: page
title: Adding app translations
---

Most apps will want to add translations at some point, and since v2.5.0 Fyne helps to make this really simple. The translation features simply sit in your project as `.json` files and the
content can be crowd-sourced by popular platforms like [Weblate](https://weblate.org).

## Setup

The Fyne toolkit is already translated and will replace recognised strings in standard locations.
To make use of this functionality for your own app strings you make use of the `fyne.io/fyne/v2/lang` package and the helper functions it provides.

The simplest way to prepare for translations is to use the `L` or `Localize` function to mark a string as translatable - if the translation is not found then the string will be used as a fallback.

```go
    title := widget.NewLabel(lang.L("My App Title"))
```

In some cases it may be desirable to label a string with a unique ID instead of using the default value - for disambiguation or other reason. In this case you would use `LocalizeKey` or the `X` alias.

```go
    title := widget.NewLabel(lang.X("window.title", "My App Window Title"))
```

That may be all you need to know to get started - skip to [translation files[(translation-files).

## Translation files

Each translation file is a simple JSON file, if you use the `L` form all you need to do is insert the string with it's translation for each language - 1 per file. For example, this may be your app `en.json`:

```json
{
  "My App Title": "My App Title"
}
```

Then your French translation could look like:

```json
{
  "My App Title": "Titre de mon application"
}
```

Each file can be most easily loaded using the Go `embed` feature - place each of the files in a directory called "translation" and then define them simply as:

```go
//go:embed translation
var translations embed.FS
```

Finally you can tell Fyne to load these translations with a single function call between `app.New` and `Run()` in your `main` function:

```go
	lang.AddTranslationsFS(translations, "translation")
```

This uses the embedded filesystem and specifies the name of the directory that the files are stored in.

When your app starts it will display using the translations for the current user's langauge configuration.

## Plurals

In more complex cases the string will change based on the number of items it refers to. For this the `lang.LocalizePlural` function (aliased to `lang.N`) is avaialble.

```go
    age := widget.NewLabel(lang.N("{{.Years}} years old", years, map[string]any{"Years": years}))
```

You can pass data in this way to any of the calls, using template syntax to insert the value. A struct with exported fields, or a map as illustrated above, can be used to insert data.

The translation file will look a little more complex for this - the key has two sub-values called "one" and "other" (some languages may use more keys).

```json
{
  "Age": {
    "one": "1 year old",
    "other": "{{.Years}} years old"
  }
}
```

