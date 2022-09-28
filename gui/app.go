package gui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/samorton/git-repo-scraper/utils"
)

func Start() {
	utils.LogMsg("starting gui...")

	var app = tview.NewApplication()
	var text = tview.NewTextView().
		SetTextColor(tcell.ColorGreen).
		SetText("(q) to quit")

	if err := app.SetRoot(text, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		}
		return event
	})
}
