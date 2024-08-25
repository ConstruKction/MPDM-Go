package main

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Widget struct {
	Layout func(gtx layout.Context, th *material.Theme) layout.Dimensions
}

type Drawer struct {
	inset layout.Inset
}

func NewDrawer(i layout.Inset) *Drawer {
	return &Drawer{inset: i}
}

func (d *Drawer) DrawWidget(gtx layout.Context, th *material.Theme, widget Widget) layout.FlexChild {
	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return d.inset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return widget.Layout(gtx, th)
		})
	})
}