package main

import (
	"darrenclevenger/kswitch/internal"
	"darrenclevenger/kswitch/internal/ui"

	"github.com/rivo/tview"
)

var app *tview.Application
var grid *tview.Grid

func main() {

	app = tview.NewApplication()

	//header := internal.GetHeader()
	header := internal.GetHeaderGrid()
	list := internal.GetClusterSelectionList()
	ui.RenderMainWindow()

	grid = tview.NewGrid().
		SetRows(2, 0).
		SetColumns(0).
		SetBorders(true).
		AddItem(header, 0, 0, 1, 3, 0, 0, false).
		AddItem(list, 1, 0, 1, 3, 40, 0, true)

	//grid = tview.NewGrid()
	//grid.SetRows(20)
	//grid.SetColumns(100)
	//grid.SetBorders(true)
	//grid.AddItem(header, 0, 0, 0, 0, 20, 20, false)
	app.SetRoot(grid, true).SetFocus(grid).Run()
}
