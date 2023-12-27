package model

import (
	"image/draw"
	d "smartCalc/domains"
)

//---------------------------------------Types START

// structure for output Model data
type ModelsOutput struct {
	Err                 bool
	Mode                int
	ModelEquationResult ModelResultEquation
	ModelEqualResult    ModelResultEqual
	ModelGraphResult    ModelResultGraph
}

// structure for Input Model data
type ModelsInput struct {
	Mode              int
	ModelEquationData ModelDataEquation
	ModelEqualData    ModelDataEqual
	ModelGraphData    ModelDataGraph
}

// structure for Input Equation data
type ModelDataEquation struct {
	Mode       int
	EqualValue string // string of equation for calc
}

// structure for Input Equal data
type ModelDataEqual struct {
	Mode       int
	EqualValue string  // string of equal for equal or graph
	XEqualStr  string  // raw data
	XEqual     float64 // Value X for equal
}

// structure for Input Graph data
type ModelDataGraph struct {
	Mode                               int
	EqualValue                         string  // string of equal for equal or graph
	XFrom, XTo, YFrom, YTo             float64 // border Value for graph
	XFromStr, XToStr, YFromStr, YToStr string  //raw data
}

// structure for Ouput Equation data
type ModelResultEquation struct {
	Err       bool   // true = error
	Mode      int    // calc - 0, equal - 1 or graph - 2
	ResultStr string // raw data
}

// structure for Ouput Equal data
type ModelResultEqual struct {
	Err       bool   // true = error
	Mode      int    // calc - 0, equal - 1 or graph - 2
	ResultStr string // raw data
}

// structure for Ouput Graph data
type ModelResultGraph struct {
	Err        bool   // true = error
	Mode       int    // calc - 0, equal - 1 or graph - 2
	ResultStr  string // raw data
	GraphImage draw.Image
}

// common interface for services
type request interface {
	GetResult() ModelsOutput
}

// struct for CalcModel
type CalcModel struct {
	Config  *d.Cfg
	request request
	history calcHistory
}

var ModelCalc = NewCalcModel(d.Config)

//---------------------------------------Types END

// Creating New CalcModel object
func NewCalcModel(cfgm *d.Cfg) *CalcModel {

	return &CalcModel{
		Config:  cfgm,
		history: *NewCalcHistory(*cfgm),
	}
}

//-------------------Implementing Models Interface

func (m *CalcModel) GetCalcResult(in ModelsInput) (out ModelsOutput) {

	if in.Mode > -1 && in.Mode < 3 {
		switch in.Mode {
		case 0:
			m.request = NewEquation(in)
		case 1:
			m.request = NewEqual(in)
		case 2:
			m.request = m.NewGraph(in)
		}

		out = m.request.GetResult()
		m.history.HistoryHandler(in, out)
		return
	}

	out.Err = true
	return
}

func (m *CalcModel) CleanHistory() {
	m.history.CleanHistory()
}

func (m *CalcModel) GetHistory() []d.HistoryItem {
	return m.history.historyData
}
