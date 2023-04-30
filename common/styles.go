package common

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

const (
	defaultForeground = ui.ColorClear
	borderColor       = ui.ColorBlue
	selectedItemBg    = ui.ColorBlack
	bold              = ui.ModifierBold
)

func SetCommonStyles(w ui.Drawable) {
	switch widget := w.(type) {
	case *widgets.List:
		widget.TitleStyle.Fg = defaultForeground
		widget.TitleStyle.Modifier = bold
		widget.BorderStyle.Fg = borderColor
		widget.TextStyle = ui.NewStyle(defaultForeground)
		widget.WrapText = false
		widget.SelectedRowStyle = ui.NewStyle(defaultForeground, selectedItemBg, bold)
	case *widgets.Table:
		widget.TitleStyle.Fg = defaultForeground
		widget.TitleStyle.Modifier = bold
		widget.BorderStyle.Fg = borderColor
		widget.TextStyle = ui.NewStyle(defaultForeground)
		widget.TextAlignment = ui.AlignCenter
		widget.FillRow = true
	case *widgets.Tree:
		widget.TitleStyle.Fg = defaultForeground
		widget.TitleStyle.Modifier = bold
		widget.BorderStyle.Fg = borderColor
		widget.TextStyle = ui.NewStyle(defaultForeground)
		widget.SelectedRowStyle = ui.NewStyle(defaultForeground, selectedItemBg, bold)
		widget.WrapText = false
	case *widgets.SparklineGroup:
		widget.TitleStyle.Fg = defaultForeground
		widget.TitleStyle.Modifier = bold
		widget.BorderStyle.Fg = borderColor
	case *widgets.PieChart:
		widget.TitleStyle.Fg = defaultForeground
		widget.TitleStyle.Modifier = bold
		widget.BorderStyle.Fg = borderColor
	case *widgets.Plot:
		widget.TitleStyle.Fg = defaultForeground
		widget.TitleStyle.Modifier = bold
		widget.BorderStyle.Fg = borderColor
		widget.AxesColor = defaultForeground
	}
}
