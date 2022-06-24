package internal

import (
	"github.com/rivo/tview"
)

func GetHeader() *tview.TextView {

	header := tview.NewTextView()
	header.SetText("Kubernetes Cluster Switcher")
	header.SetBorder(false)
	header.SetDynamicColors(true)
	header.SetTextAlign(tview.AlignCenter)
	return header
}

func GetHeaderGrid() *tview.Grid {

	title := tview.NewTextView()
	title.SetText("JRS Kubernetes Cluster Switcher")
	title.SetBorder(false)

	clusterName := tview.NewTextView()
	clusterName.SetText("Selected: prbuild.k8s.clusterName.com")
	clusterName.SetBorder(false)

	grid := tview.NewGrid()
	grid.SetRows(0)
	grid.SetColumns(0, 0, 0)
	grid.AddItem(title, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(clusterName, 0, 2, 1, 1, 0, 0, false)

	return grid
}

func GetClusterSelectionList() *tview.List {

	list := tview.NewList()
	list.AddItem("Line Item 1", "", 'a', nil)
	list.AddItem("Line Item 2", "", 'a', nil)
	list.AddItem("Line Item 3", "", 'a', nil)
	list.AddItem("Line Item 4", "", 'a', nil)
	return list

}
