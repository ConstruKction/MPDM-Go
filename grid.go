package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"mpdm.test/ui_components"
)

var (
	pi = ui_components.NewPathInput()
	dd  = ui_components.NewDropdown([]string{})
)

func drawGrid(gtx layout.Context, th *material.Theme) layout.Dimensions {
	wgs := map[string]Widget{
		"inputPath": {Layout: func(gtx layout.Context, th *material.Theme) layout.Dimensions { return pi.Layout(gtx, th) }},
		"dropdown":  {Layout: func(gtx layout.Context, th *material.Theme) layout.Dimensions { return dd.Layout(gtx, th) }},
	}

	dw := NewDrawer(setInset(10, 10, 10, 10))

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		dw.DrawWidget(gtx, th, wgs["inputPath"]),
		dw.DrawWidget(gtx, th, wgs["dropdown"]),
	)
}