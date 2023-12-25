package controllers

import (
	d "smartCalc/domains"
	m "smartCalc/model"
	// "fyne.io/fyne/v2"
)

type View interface {
	// SetLink(*chan string)
	GetUIData() interface{}
	DisplayResult(interface{})
	UpdateHistory([]d.HistoryItem)
}

type Model interface {
	GetHistory() []d.HistoryItem
	GetCalcResult(m.ModelsInput) m.ModelsOutput
	CleanHistory()
}

// type errStruct struct {
// 	Err       bool
// 	mode      int
// 	resultStr string
// }

// presenters structure
type presenter struct {
	// cnv             convert
	ViewDataChannel chan string
}

// create presenters object
func NewPresenter() *presenter {
	return &presenter{
		ViewDataChannel: make(chan string),
	}
}
