package ui_components

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type PathInput struct {
	Editor *widget.Editor
	Button *widget.Clickable
}

func NewPathInput() *PathInput {
	return &PathInput{
		Editor: new(widget.Editor),
		Button: new(widget.Clickable),
	}
}

func (pi *PathInput) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Editor(th, pi.Editor, "Mod/Song Directory").Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.Inset{Left: unit.Dp(10)}.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Button(th, pi.Button, "Browse").Layout(gtx)
			})
		}),
	)
}