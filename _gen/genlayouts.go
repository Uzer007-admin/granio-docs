package main

import (
	"image/color"
	"image/png"
	"os"
	"path/filepath"

	"github.com/Uzer007-admin/granio/v2"
	"github.com/Uzer007-admin/granio/v2/canvas"
	"github.com/Uzer007-admin/granio/v2/layout"
	"github.com/Uzer007-admin/granio/v2/test"
	"github.com/Uzer007-admin/granio/v2/theme"
)

type drawItem struct {
	name string
	lay  granio.CanvasObject
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
	border := granio.NewContainerWithLayout(layout.NewBorderLayout(bObjs[0], nil, bObjs[1], nil), bObjs...)

	smaller := makeObjs()
	for _, s := range smaller {
		s.(*canvas.Rectangle).SetMinSize(granio.NewSize(s.MinSize().Width, 25))
	}
	top := granio.NewContainerWithLayout(layout.NewHBoxLayout(), smaller...)
	left := granio.NewContainerWithLayout(layout.NewVBoxLayout(), makeObjs()...)
	content := granio.NewContainerWithLayout(layout.NewGridWrapLayout(granio.NewSize(40, 40)), makeObjs()...)
	combined := granio.NewContainerWithLayout(layout.NewBorderLayout(top, nil, left, nil), top, left, content)

	return []drawItem{
		{"hbox", granio.NewContainerWithLayout(layout.NewHBoxLayout(), makeObjs()...)},
		{"vbox", granio.NewContainerWithLayout(layout.NewVBoxLayout(), makeObjs()...)},
		{"border", border},
		{"center", granio.NewContainerWithLayout(layout.NewCenterLayout(), makeObjs()...)},
		{"form", granio.NewContainerWithLayout(layout.NewFormLayout(), append(makeObjs(), canvas.NewRectangle(boxColor()))...)},
		{"grid", granio.NewContainerWithLayout(layout.NewGridLayout(2), makeObjs()...)},
		{"gridwrap", granio.NewContainerWithLayout(layout.NewGridWrapLayout(granio.NewSize(75, 75)), makeObjs()...)},
		{"max", granio.NewContainerWithLayout(layout.NewMaxLayout(), makeObjs()...)},
		{"padded", granio.NewContainerWithLayout(layout.NewPaddedLayout(), makeObjs()...)},
		{"rowwrap", granio.NewContainerWithLayout(layout.NewRowWrapLayout(), makeMoreObjs()...)},
		{"custompadded", granio.NewContainerWithLayout(layout.NewCustomPaddedLayout(10, 20, 30, 15), makeObjs()...)},
		{"combined", combined},
	}
}

func makeObjs() []granio.CanvasObject {
	prop1 := canvas.NewRectangle(boxColor())
	prop1.SetMinSize(granio.NewSize(50, 50))
	prop2 := canvas.NewRectangle(boxColor())
	prop2.SetMinSize(granio.NewSize(25, 25))
	prop3 := canvas.NewRectangle(boxColor())
	prop3.SetMinSize(granio.NewSize(75, 75))
	return []granio.CanvasObject{prop1, prop2, prop3}
}

func makeMoreObjs() []granio.CanvasObject {
	prop1 := canvas.NewRectangle(boxColor())
	prop1.SetMinSize(granio.NewSize(50, 50))
	prop2 := canvas.NewRectangle(boxColor())
	prop2.SetMinSize(granio.NewSize(100, 25))
	prop3 := canvas.NewRectangle(boxColor())
	prop3.SetMinSize(granio.NewSize(50, 75))
	prop4 := canvas.NewRectangle(boxColor())
	prop4.SetMinSize(granio.NewSize(75, 75))
	return []granio.CanvasObject{prop1, prop2, prop3, prop4}
}

func draw(scene granio.CanvasObject, name string, c test.WindowlessCanvas, themeName string) {
	fileName := filepath.Join(imgDir, name+"-"+themeName+".png")
	file, err := os.Create(fileName)
	if err != nil {
		granio.LogError("err", err)
		file, err = os.Open(fileName)
		if err != nil {
			granio.LogError("Unable to open file for writing", err)
			return
		}
	}

	c.SetScale(2.0) // get HiDPI output so we can render nicely on fancy screens :)
	c.SetPadded(false)
	c.SetContent(scene)
	c.Resize(granio.NewSize(220, 180))
	img := c.Capture()
	err = png.Encode(file, img)
	if err != nil {
		granio.LogError("Unable to write image", err)
	}
}

func main() {
	w := test.NewWindow(nil)
	c := w.Canvas().(test.WindowlessCanvas)

	pwd, _ := os.Getwd()
	imgDir = filepath.Join(pwd, "..", "static", "images", "layouts")

	granio.CurrentApp().Settings().SetTheme(transparentBGTheme(theme.LightTheme()))
	for _, item := range makeDrawList() {
		draw(item.lay, item.name, c, "light")
	}

	granio.CurrentApp().Settings().SetTheme(transparentBGTheme(theme.DarkTheme()))
	for _, item := range makeDrawList() {
		draw(item.lay, item.name, c, "dark")
	}
}

type transparent struct {
	th granio.Theme
}

func transparentBGTheme(th granio.Theme) granio.Theme {
	return &transparent{th: th}
}

func (t *transparent) Color(n granio.ThemeColorName, v granio.ThemeVariant) color.Color {
	if n == theme.ColorNameBackground {
		return color.Transparent
	}

	return t.th.Color(n, v)
}

func (t *transparent) Font(s granio.TextStyle) granio.Resource {
	return t.th.Font(s)
}

func (t *transparent) Icon(n granio.ThemeIconName) granio.Resource {
	return t.th.Icon(n)
}

func (t *transparent) Size(n granio.ThemeSizeName) float32 {
	return t.th.Size(n)
}
