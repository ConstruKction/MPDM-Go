package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func setWindowSize() (int, int, error) {
	x, y, err := getScreenSize()
	if err != nil {
		return 0, 0, err
	}

	width := (x * 40) / 100
	height := (y * 70) / 100

	return width, height, nil
}

func main() {
	go func() {
		win := new(app.Window)

		width, height, err := setWindowSize()
		if err != nil {
			log.Fatal(err)
		}
		win.Option(app.Size(unit.Dp(float32(width)), unit.Dp(float32(height))))

		err = run(win)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(win *app.Window) error {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	win.Option(app.Title("MPDM"))

	var ops op.Ops
	for {
		switch e := win.Event().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			drawGrid(gtx, th)
			e.Frame(gtx.Ops)
		}
	}
}
