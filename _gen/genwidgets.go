package main

import (
	"image/color"
	"image/png"
	"os"
	"path/filepath"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/test"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type drawItem struct {
	name string
	obj fyne.CanvasObject
}

var (
	imgDir string
	drawList = []drawItem{
		{"button", widget.NewButtonWithIcon("Cancel", theme.CancelIcon(), func() {})},
		{"check", &widget.Check{Text: "Check", Checked: true}},
		{"entry", &widget.Entry{PlaceHolder: "Entry"}},
		{"form", &widget.Form{Items:
			[]*widget.FormItem{
				{Text: "Username", Widget: widget.NewEntry()},
				{Text: "Password", Widget: widget.NewPasswordEntry()}},
			OnSubmit: func() {}, OnCancel: func() {}}},
		{"hyperlink", widget.NewHyperlink("fyne.io", nil)},
		{"icon", widget.NewIcon(theme.ContentPasteIcon())},
		{"label", widget.NewLabel("Text label")},
		{"password", &widget.Entry{PlaceHolder:"Password", Password: true}},
		{"progress", &widget.ProgressBar{Value:0.74}},
		{"progressinf", widget.NewProgressBarInfinite()},
		{"radio", widget.NewRadio([]string{"Item 1"}, func(string) {})},
		{"scroll", widget.NewScrollContainer(widget.NewLabel("Scroll"))},
		{"select", widget.NewSelect([]string{"1", "2"}, func(string) {})},
		{"slider", widget.NewSlider(-5, 25)},
		{"tabcontainer", widget.NewTabContainer(
			widget.NewTabItem("Tab1", canvas.NewRectangle(color.Transparent)),
			widget.NewTabItem("Tab2", canvas.NewRectangle(color.Transparent)))},
		{"toolbar", widget.NewToolbar(widget.NewToolbarAction(theme.MailComposeIcon(), func() {}),
			widget.NewToolbarSeparator(),
			widget.NewToolbarSpacer(),
			widget.NewToolbarAction(theme.ContentCutIcon(), func() {}),
			widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
			widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		)},
	}
)

func draw(obj fyne.CanvasObject, name string, c fyne.Canvas, themeName string) {
	fileName := filepath.Join(imgDir, name+"-"+themeName+".png")
	file, err := os.Create(fileName)
	if err != nil {
		fyne.LogError("err", err)
		file, err = os.Open(fileName)
		if err != nil {
			fyne.LogError("Unable to open file for writing", err)
			return
		}
	}

	c.SetScale(2.0) // get HiDPI output so we can render nicely on fancy screens :)
	c.SetContent(obj)
	c.Content().Refresh() // for some reason a few themed items did not pick up the current theme :(
	if scroll, ok := c.Content().(*widget.ScrollContainer); ok {
		scroll.Content.Refresh() // as above because scrollcontainer refresh does not propagate
	}
	img := c.Capture()
	err = png.Encode(file, img)
	if err != nil {
		fyne.LogError("Unable to write image", err)
	}
}

func main() {
	w := test.NewWindow(nil)
	c := w.Canvas()

	pwd, _ := os.Getwd()
	imgDir = filepath.Join(pwd, "images", "widgets")

	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
	for _, item := range drawList {
		draw(item.obj, item.name, c, "light")
	}

	fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
	for _, item := range drawList {
		draw(item.obj, item.name, c, "dark")
	}
}
