---
title: AppTabs

weight: 4070

aliases:
- /container/apptabs.html
- /tour/container/apptabs
---

The AppTabs container is used to allow the user to switch
between different content panels. Tabs are either just text
or text and an icon. It is recommended not to mix some tabs
having icons and some without. A tab container is created
using `container.NewAppTabs(...)` and passing
`container.TabItem` items (that can be created using
`container.NewTabItem(...)`).

The tab container can be configured by setting the location
of tabs, one of `container.TabLocationTop`, `container.TabLocationBottom`,
`container.TabLocationLeading` and `container.TabLocationTrailing`.
The default location is top.

```go
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
```

When loaded on a mobile device the tab location may be ignored.
In a portrait orientation a leading or trailing position will
be changed to bottom. When in landscape orientation, the top or bottom
positions will be moved to leading.
