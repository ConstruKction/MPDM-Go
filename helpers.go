package main

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"github.com/kbinani/screenshot"
)

func getScreenSize() (int, int, error) {
	b := screenshot.GetDisplayBounds(0)
	x := b.Dx()
	y := b.Dy()

	return x, y, nil
}

func setInset(t int, b int, l int, r int) layout.Inset {
	return layout.Inset{
		Top: 	unit.Dp(t),
		Bottom:	unit.Dp(b),
		Left: 	unit.Dp(l),
		Right:	unit.Dp(r),
	}
}