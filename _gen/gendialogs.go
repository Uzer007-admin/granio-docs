package main

import (
	"image/color"
	"image/png"
	"os"
	"path/filepath"

	"github.com/Uzer007-admin/granio/v2"
	"github.com/Uzer007-admin/granio/v2/canvas"
	"github.com/Uzer007-admin/granio/v2/dialog"
	"github.com/Uzer007-admin/granio/v2/test"
	"github.com/Uzer007-admin/granio/v2/theme"
	"github.com/Uzer007-admin/granio/v2/widget"
)

type drawItem struct {
	name string
	dia  func(granio.Window)
}

var (
	imgDir string
)

func makeDrawList() []drawItem {
	return []drawItem{
		{"information", func(w granio.Window) { dialog.ShowInformation("Some Information", "This is important", w) }},
		{"color", func(w granio.Window) {
			dialog.ShowColorPicker("Choose a color", "Please pick a color...", func(color.Color) {}, w)
		}},
		{"confirm", func(w granio.Window) { dialog.ShowConfirm("Please Confirm", "Are you sure..?", func(bool) {}, w) }},
		{"fileopen", func(w granio.Window) {
			dialog.ShowFileOpen(func(granio.URIReadCloser, error) {}, w)
		}},
		{"form", func(w granio.Window) {
			dialog.ShowForm("Form Input", "Enter", "Cancel",
				[]*widget.FormItem{widget.NewFormItem("Enter a string...", widget.NewEntry())}, func(bool) {}, w)
		}},
		{"custom", func(w granio.Window) {
			text := widget.NewTextGrid()
			text.SetText("Custom content")
			text.SetStyleRange(0, 7, 0, 14, widget.TextGridStyleWhitespace)
			dialog.ShowCustom("Custom Dialog", "Cancel", text, w)
		}},
	}
}

func draw(scene func(granio.Window), name string, w granio.Window, c test.WindowlessCanvas, themeName string) {
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

	c.Resize(granio.NewSize(460, 350)) // too big, but give space, will be sized down
	c.SetScale(2.0)                    // get HiDPI output so we can render nicely on fancy screens :)
	for _, over := range c.Overlays().List() {
		c.Overlays().Remove(over)
	}

	c.SetContent(canvas.NewRectangle(theme.BackgroundColor()))
	scene(w)
	if name == "fileopen" {
		w.Resize(granio.NewSize(458, 318))
	} else {
		min := c.Overlays().Top().MinSize()
		w.Resize(min.Add(granio.NewSize(theme.Padding()*2, theme.Padding()*2)))
		c.Overlays().Top().Refresh()
	}

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
	imgDir = filepath.Join(pwd, "..", "static", "images", "dialogs")

	granio.CurrentApp().Settings().SetTheme(theme.LightTheme())
	for _, item := range makeDrawList() {
		draw(item.dia, item.name, w, c, "light")
	}

	granio.CurrentApp().Settings().SetTheme(theme.DarkTheme())
	for _, item := range makeDrawList() {
		draw(item.dia, item.name, w, c, "dark")
	}
}
