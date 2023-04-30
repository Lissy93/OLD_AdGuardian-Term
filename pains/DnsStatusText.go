package pains

import (
	"fmt"
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/lissy93/adguardian-term/fetch"
	"github.com/lissy93/adguardian-term/values"
)

func DnsStatus() *widgets.Paragraph {
	statusEndpoint := values.ENDPOINT + "/control/status"
	stats, err := fetch.GetAdGuardStatus(statusEndpoint, values.USERNAME, values.PASSWORD)
	if err != nil {
		log.Fatalf("failed to fetch AdGuard status: %v", err)
	}
	chartBlockPercent := widgets.NewParagraph()
	chartBlockPercent.BorderStyle.Fg = ui.ColorBlue
	text := ""
	if stats.Running {
		text += "Status: [Running](fg:green,mod:bold)"
	} else {
		text += "Status: [Not Running](fg:red,mod:bold)"
	}
	text += " | "
	if stats.DHCPAvailable {
		text += "DHCP: [Available](fg:green,mod:bold)"
	} else {
		text += "DHCP: [Unavailable](fg:red,mod:bold)"
	}
	text += " | "
	if stats.ProtectionEnabled {
		text += "Protection: [Enabled](fg:green,mod:bold)"
	} else {
		text += "Protection: [Disabled](fg:red,mod:bold)"
	}
	text += " | "
	text += fmt.Sprintf("Version: [%s](fg:blue,mod:bold)", stats.Version)

	chartBlockPercent.Text = text
	return chartBlockPercent
}
