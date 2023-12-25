package controllers

import (
	d "smartCalc/domains"
	m "smartCalc/model"
)

type View interface {
	GetUIData() interface{}
	DisplayResult(interface{})
	UpdateHistory([]d.HistoryItem)
}

type Model interface {
	GetHistory() []d.HistoryItem
	GetCalcResult(m.ModelsInput) m.ModelsOutput
	CleanHistory()
}

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
