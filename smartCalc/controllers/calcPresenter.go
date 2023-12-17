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

type errStruct struct {
	Err       bool
	mode      int
	resultStr string
}

// presenters structure
type presenter struct {
	// app             fyne.App
	cnv             convert
	ViewDataChannel chan string
}

// create presenters object
func NewPresenter() *presenter {
	return &presenter{
		ViewDataChannel: make(chan string),
	}
}

// main presenters function
// func (pr *presenter) CrossRoad(v View, m Model) {
// 	e := errStruct{
// 		Err:       true,
// 		mode:      0,
// 		resultStr: "error",
// 	}

// 	for {

// 		// put the signal from View
// 		viewSignal := <-pr.ViewDataChannel

// 		// handle signal
// 		switch viewSignal {
// 		case "cleanhistory": // Mode Clean History

// 			m.CleanHistory()                // clean history in model
// 			v.UpdateHistory(m.GetHistory()) // update history in view

// 		case "ready": // start main handle process(calculating)

// 			if temp, err := pr.cnv.UIToModel(v.GetUIData()); !err {
// 				v.DisplayResult(pr.cnv.ModelToUI(m.GetCalcResult(temp)))
// 				v.UpdateHistory(m.GetHistory()) // update history in view
// 			} else {
// 				v.DisplayResult(e)
// 			}
// 		}
// 	}
// }
