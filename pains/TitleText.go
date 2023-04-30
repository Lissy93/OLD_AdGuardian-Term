package pains

import (
	"fmt"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lissy93/adguardian-term/values"
)

func Title() *widgets.Paragraph {
	title := widgets.NewParagraph()
	title.Text = fmt.Sprintf("[%s](mod:bold)", values.TITLE)
	title.BorderStyle.Fg = ui.ColorBlue
	return title
}
