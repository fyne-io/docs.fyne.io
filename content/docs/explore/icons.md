---
layout: page
title: Theme Icons
slug: icons

weight: 2060
---
Each of the following icons is available via the `theme` package as a function. 
For example `theme.InfoIcon()`.

The icons are also available via their source icon name by using the `ThemeIconName` 
with the `Icon` method on a struct implementing `fyne.Theme`. For example `theme.Icon(theme.IconNameInfo)`.

## List

<style>
  html:not([data-dark-mode]) .theme-icon-list svg :not([fill]),
  html:not([data-dark-mode]) .theme-icon-list svg [fill*="#"],
  html:not([data-dark-mode]) .theme-icon-list svg [style*="fill:#"] {
    fill: black !important;
  }
  html[data-dark-mode] .theme-icon-list svg :not([fill]),
  html[data-dark-mode] .theme-icon-list svg [fill*="#"],
  html[data-dark-mode] .theme-icon-list svg [style*="fill:#"] {
    fill: white !important;
  }
  
  .theme-icon-list figure {
	display: flex;
  }

  .theme-icon-list figure svg {
	height: 32pt !important;
	width: 32pt !important;
	margin-top: -7pt;
	margin-bottom: 2pt;
  }

  .theme-icon-list figure figcaption {
	padding-left: 8pt;
  }
</style>

{{< iconlist >}}

## Using other color sets

Each icon can be used as a source for a particular themed color using the various public helper methods:

* `NewDisabledThemedResource`
* `NewErrorThemedResource`
* `NewInvertedThemedResource`
* `NewPrimaryThemedResource`

By default, all icons adapt to the current theme foreground using `NewThemedResource`
which uses the theme foreground color. All Icons are SVG `width="24"`, `height="24"`.
