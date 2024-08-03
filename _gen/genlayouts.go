package main

import (
	"image/color"
	"image/png"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/theme"
)

type drawItem struct {
	name string
	lay  fyne.CanvasObject
}

var (
	imgDir string
)

func boxColor() color.Color {
	base := theme.DisabledButtonColor()
	r, _, _, _ := base.RGBA()
	if r>>8 < 0x80 {
		return color.NRGBA{R: 0x18, G: 0x18, B: 0x18, A: 0xff}
	}

	return color.NRGBA{R: 0xe2, G: 0xe2, B: 0xe2, A: 0xff}
}

func makeDrawList() []drawItem {
	bObjs := makeObjs()
	border := fyne.NewContainerWithLayout(layout.NewBorderLayout(bObjs[0], nil, bObjs[1], nil), bObjs...)

	smaller := makeObjs()
	for _, s := range smaller {
		s.(*canvas.Rectangle).SetMinSize(fyne.NewSize(s.MinSize().Width, 25))
	}
	top := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), smaller...)
	left := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), makeObjs()...)
	content := fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(40, 40)), makeObjs()...)
	combined := fyne.NewContainerWithLayout(layout.NewBorderLayout(top, nil, left, nil), top, left, content)

	return []drawItem{
		{"hbox", fyne.NewContainerWithLayout(layout.NewHBoxLayout(), makeObjs()...)},
		{"vbox", fyne.NewContainerWithLayout(layout.NewVBoxLayout(), makeObjs()...)},
		{"border", border},
		{"center", fyne.NewContainerWithLayout(layout.NewCenterLayout(), makeObjs()...)},
		{"form", fyne.NewContainerWithLayout(layout.NewFormLayout(), append(makeObjs(), canvas.NewRectangle(boxColor()))...)},
		{"grid", fyne.NewContainerWithLayout(layout.NewGridLayout(2), makeObjs()...)},
		{"gridwrap", fyne.NewContainerWithLayout(layout.NewGridWrapLayout(fyne.NewSize(75, 75)), makeObjs()...)},
		{"max", fyne.NewContainerWithLayout(layout.NewMaxLayout(), makeObjs()...)},
		{"padded", fyne.NewContainerWithLayout(layout.NewPaddedLayout(), makeObjs()...)},
		{"custompadded", fyne.NewContainerWithLayout(layout.NewCustomPaddedLayout(5, 15, 10, 20), makeObjs()...)},
		{"combined", combined},
	}
}

func makeObjs() []fyne.CanvasObject {
	prop1 := canvas.NewRectangle(boxColor())
	prop1.SetMinSize(fyne.NewSize(50, 50))
	prop2 := canvas.NewRectangle(boxColor())
	prop2.SetMinSize(fyne.NewSize(25, 25))
	prop3 := canvas.NewRectangle(boxColor())
	prop3.SetMinSize(fyne.NewSize(75, 75))
	return []fyne.CanvasObject{prop1, prop2, prop3}
}

func draw(scene fyne.CanvasObject, name string, c test.WindowlessCanvas, themeName string) {
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
	c.SetPadded(false)
	c.SetContent(scene)
	c.Resize(fyne.NewSize(220, 180))
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
	imgDir = filepath.Join(pwd, "..", "images", "layouts")

	fyne.CurrentApp().Settings().SetTheme(transparentBGTheme(theme.LightTheme()))
	for _, item := range makeDrawList() {
		draw(item.lay, item.name, c, "light")
	}

	fyne.CurrentApp().Settings().SetTheme(transparentBGTheme(theme.DarkTheme()))
	for _, item := range makeDrawList() {
		draw(item.lay, item.name, c, "dark")
	}
}

type transparent struct {
	th fyne.Theme
}

func transparentBGTheme(th fyne.Theme) fyne.Theme {
	return &transparent{th: th}
}

func (t *transparent) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	if n == theme.ColorNameBackground {
		return color.Transparent
	}

	return t.th.Color(n, v)
}

func (t *transparent) Font(s fyne.TextStyle) fyne.Resource {
	return t.th.Font(s)
}

func (t *transparent) Icon(n fyne.ThemeIconName) fyne.Resource {
	return t.th.Icon(n)
}

func (t *transparent) Size(n fyne.ThemeSizeName) float32 {
	return t.th.Size(n)
}
