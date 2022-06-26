package ui

import (
	"darrenclevenger/kswitch/internal/k8s"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var main *tview.Grid
var title *tview.TextView
var current_cluster_text *tview.TextView
var gutter_right *tview.Box
var gutter_left *tview.Box
var cluster_list *tview.List
var footer *tview.Box

func RenderMainWindow() *tview.Grid {

	// Create the main title
	title := tview.NewTextView()
	title.SetText("K8s Selector")
	title.SetTextColor(tcell.ColorDeepSkyBlue)

	//Right, left, footer padding border
	gutter_right = tview.NewBox()
	gutter_left = tview.NewBox()
	footer = tview.NewBox()

	//Create the current cluster label.
	current_cluster_text = tview.NewTextView()
	current_cluster_text.SetTextAlign(tview.AlignRight)
	current_cluster_text.SetTextColor(tcell.ColorLightYellow)

	// Set the currently selected cluster.
	SetCurrentCluster()

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

	fillClustersList()

	main := tview.NewGrid()
	main.SetRows(2, 0, 2)
	main.SetColumns(6, 0, 0, 6)
	main.SetBorder(true)

	main.AddItem(gutter_left, 0, 0, 2, 1, 0, 0, false)
	main.AddItem(title, 0, 1, 1, 1, 10, 0, false)
	main.AddItem(current_cluster_text, 0, 2, 1, 1, 0, 60, false)
	main.AddItem(gutter_right, 0, 3, 2, 1, 0, 0, false)
	main.AddItem(cluster_list, 1, 1, 1, 2, 0, 0, true)
	main.AddItem(footer, 2, 0, 1, 4, 10, 0, false)

	return main
}

func onItemSelected() {

	idx := cluster_list.GetCurrentItem()
	if idx != -1 {
		cluster_name, _ := cluster_list.GetItemText(idx)
		k8s.SetCurrentClusterContext(cluster_name)
	}
}

func SetCurrentCluster() {

	current_cluster := k8s.GetCurrentCluster()
	current_cluster_text.SetText(current_cluster)
}

func fillClustersList() {

	list := k8s.GetClusterNames()

	for _, c := range list {
		cluster_list.AddItem(c.Name, "", 0, onItemSelected)
	}
}
