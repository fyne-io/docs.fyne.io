package main

import (
	"image/color"
	"image/png"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type drawItem struct {
	name string
	dia  func(fyne.Window)
}

var (
	imgDir string
)

func makeDrawList() []drawItem {
	return []drawItem{
		{"information", func(w fyne.Window) { dialog.ShowInformation("Some Information", "This is a thing to know", w) }},
		{"color", func(w fyne.Window) {
			dialog.ShowColorPicker("Choose a color", "Please pick a color...", func(color.Color) {}, w)
		}},
		{"confirm", func(w fyne.Window) { dialog.ShowConfirm("Please Confirm", "Are you sure..?", func(bool) {}, w) }},
		{"fileopen", func(w fyne.Window) {
			dialog.ShowFileOpen(func(fyne.URIReadCloser, error) {}, w)}},
		{"form", func(w fyne.Window) {
			dialog.ShowForm("Form Input", "Enter", "Cancel",
				[]*widget.FormItem{widget.NewFormItem("Enter a string...", widget.NewEntry())}, func(bool) {}, w)
		}},
		{"custom", func(w fyne.Window) {
			text := widget.NewTextGrid()
			text.SetText("Custom content")
			text.SetStyleRange(0, 7, 0, 14, widget.TextGridStyleWhitespace)
			dialog.ShowCustom("Custom Dialog", "Cancel", text, w)
		}},
	}
}

func draw(scene func(fyne.Window), name string, w fyne.Window, c test.WindowlessCanvas, themeName string) {
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
	for _, over := range c.Overlays().List() {
		c.Overlays().Remove(over)
	}

	c.SetContent(canvas.NewRectangle(theme.BackgroundColor()))
	scene(w)
	if name == "fileopen" {
		w.Resize(fyne.NewSize(440, 320))
		c.Overlays().Top().Resize(fyne.NewSize(432, 312))
	} else {
		min := c.Overlays().Top().MinSize()
		w.Resize(min.Add(fyne.NewSize(theme.Padding()*2, theme.Padding()*2)))
		c.Overlays().Top().Resize(min)
	}

	img := c.Capture()
	err = png.Encode(file, img)
	if err != nil {
		fyne.LogError("Unable to write image", err)
	}
}

func main() {
	w := test.NewWindow(nil)
	c := w.Canvas().(test.WindowlessCanvas)

	pwd, _ := os.Getwd()
	imgDir = filepath.Join(pwd, "..", "static", "images", "dialogs")

	fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
	for _, item := range makeDrawList() {
		draw(item.dia, item.name, w, c, "light")
	}

	fyne.CurrentApp().Settings().SetTheme(theme.DarkTheme())
	for _, item := range makeDrawList() {
		draw(item.dia, item.name, w, c, "dark")
	}
}
