package ui_components

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Dropdown struct {
	Selected   string
	Options    []string
	Expanded   bool
	Clickable  widget.Clickable
	OptionBtns []widget.Clickable
}

func NewDropdown(options []string) *Dropdown {
	optionBtns := make([]widget.Clickable, len(options))
	return &Dropdown{
		Options:    options,
		OptionBtns: optionBtns,
	}
}

func (dd *Dropdown) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	if dd.Clickable.Clicked(gtx) {
		dd.Expanded = !dd.Expanded
	}

	items := []layout.FlexChild{
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			btnText := "Filter by pack"
			if dd.Selected != "" {
				btnText = dd.Selected
			}
			return material.Button(th, &dd.Clickable, btnText).Layout(gtx)
		}),
	}

	if dd.Expanded {
		for i, option := range dd.Options {
			i := i
			items = append(items, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				if dd.OptionBtns[i].Clicked(gtx) {
					dd.Selected = option
					dd.Expanded = false
				}
				return material.Button(th, &dd.OptionBtns[i], option).Layout(gtx)
			}))
		}
	}

	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx, items...)
}
