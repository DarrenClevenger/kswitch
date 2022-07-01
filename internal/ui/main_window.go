package ui

import (
	"darrenclevenger/kswitch/internal/k8s"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app *tview.Application
var main *tview.Grid
var title *tview.TextView
var current_cluster_text *tview.TextView
var gutter_right *tview.Box
var gutter_left *tview.Box
var cluster_list *tview.List
var footer *tview.TextView

func RenderMainWindow(application *tview.Application) *tview.Grid {

	app = application

	// Create the main title
	title = tview.NewTextView()
	title.SetText("K8s Selector")
	title.SetTextColor(tcell.ColorDeepSkyBlue)

	//Right, left, footer padding border
	gutter_right = tview.NewBox()
	gutter_left = tview.NewBox()
	footer = tview.NewTextView()

	//Create the current cluster label.
	current_cluster_text = tview.NewTextView()
	current_cluster_text.SetTextAlign(tview.AlignRight)
	current_cluster_text.SetTextColor(tcell.ColorLightYellow)

	// Set the currently selected cluster.
	cluster_name := SetCurrentCluster()

	//Create the cluster List
	cluster_list = tview.NewList()
	cluster_list.SetTitle("Clusters")
	cluster_list.SetBorder(true)
	cluster_list.SetBorderPadding(1, 1, 1, 1)
	cluster_list.SetBorderColor(tcell.ColorSkyblue)
	cluster_list.SetTitleAlign(tview.AlignLeft)
	cluster_list.SetHighlightFullLine(true)
	cluster_list.ShowSecondaryText(false)
	cluster_list.SetBorderAttributes(tcell.AttrDim)

	fillClustersList(cluster_name)

	main = tview.NewGrid()
	main.SetRows(2, 0, 4)
	main.SetColumns(6, 0, 0, 6)
	main.SetBorder(true)

	main.AddItem(gutter_left, 0, 0, 2, 1, 0, 0, false)
	main.AddItem(title, 0, 1, 1, 1, 10, 0, false)
	main.AddItem(current_cluster_text, 0, 2, 1, 1, 0, 60, false)
	main.AddItem(gutter_right, 0, 3, 2, 1, 0, 0, false)
	main.AddItem(cluster_list, 1, 1, 1, 2, 0, 0, true)
	main.AddItem(footer, 2, 0, 1, 4, 10, 0, false)

	main.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		switch event.Rune() {
		case 'r':
			fillClustersList("")
			return event
		case 'd':
			getClusterDetails()
			return event
		}

		return event
	})

	footer.SetText("\n\t(q) Quit\t(r) Refresh Clusters")

	return main
}

func getClusterDetails() {

}

func onItemSelected() {

	idx := cluster_list.GetCurrentItem()
	if idx != -1 {
		cluster_name, _ := cluster_list.GetItemText(idx)
		err := k8s.SetCurrentClusterContext(cluster_name)
		if err != nil {
			panic(err)
		}

		SetCurrentCluster()
	}
}

func SetCurrentCluster() string {

	current_cluster := k8s.GetCurrentCluster()
	current_cluster_text.SetText(current_cluster)

	return current_cluster
}

func fillClustersList(cluster_name string) {

	list := k8s.GetClusterNames()
	cluster_list.Clear()

	for i, c := range list {
		cluster_list.AddItem(c.Name, "", 0, onItemSelected)
		if cluster_name != "" {
			if cluster_name == c.Name {
				cluster_list.SetCurrentItem(i)
			}
		}
	}

}
