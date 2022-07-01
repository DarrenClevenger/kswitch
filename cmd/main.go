package main

import (
	"darrenclevenger/kswitch/internal/ui"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application

func main() {

	app = tview.NewApplication()

	app.SetBeforeDrawFunc(func(screen tcell.Screen) bool {
		screen.Clear()
		return false
	})

	main := ui.RenderMainWindow(app)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Rune() == 'q' {
			app.Stop()
		}

		return event
	})

	app.SetRoot(main, true).SetFocus(main).Run()
}
