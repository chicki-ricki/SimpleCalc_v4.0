package model

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	d "smartCalc/domains"
)

// const FILENAME string = "history.json"

type calcHistory struct {
	config          d.Cfg
	fileName        string
	historyData     []d.HistoryItem
	tempHistoryItem d.HistoryItem
}

//--------------------------------------Public Methods

// create calcHistory object
func NewCalcHistory(c d.Cfg) *calcHistory {
	return &calcHistory{
		config:      c,
		fileName:    c.HistoryFile,
		historyData: readHistoryJson(c.HistoryFile),
	}
}

// history hendler
func (h *calcHistory) HistoryHandler(in ModelsInput, out ModelsOutput) {
	h.createHistoryItem(in, out)
	h.addHistoryString()
	h.writeHistoryJson(h.createHistoryJson())
}

func (h *calcHistory) CleanHistory() {
	h.historyData = []d.HistoryItem{}
	h.writeHistoryJson(h.createHistoryJson())
}

//----------------------------------------Creating historyItem START

// creating history string
func (h *calcHistory) createHistoryItem(in ModelsInput, out ModelsOutput) {
	// Creating new history string
	switch in.Mode {
	case 0:
		h.createHistoryCalc(in, out)
	case 1:
		h.createHistoryEqual(in, out)
	case 2:
		h.createHistoryGraph(in, out)
	}
}

// creating history string for Calc mode
func (h *calcHistory) createHistoryCalc(in ModelsInput, out ModelsOutput) {
	h.tempHistoryItem.Mode = "calc"
	h.tempHistoryItem.Equation = in.ModelEquationData.EqualValue
	h.tempHistoryItem.Entrys = ""
	if out.Err {
		h.tempHistoryItem.Result = "error"
	} else {
		h.tempHistoryItem.Result = out.ModelEquationResult.ResultStr
	}
}

// creating history string for Equal mode
func (h *calcHistory) createHistoryEqual(in ModelsInput, out ModelsOutput) {

	h.tempHistoryItem.Mode = "equal"
	h.tempHistoryItem.Equation = in.ModelEqualData.EqualValue
	h.tempHistoryItem.XEqual = in.ModelEqualData.XEqualStr
	h.tempHistoryItem.Entrys = "{X=" + in.ModelEqualData.XEqualStr + "} "
	if out.Err {
		h.tempHistoryItem.Result = "error"
	} else {
		h.tempHistoryItem.Result = out.ModelEqualResult.ResultStr
	}
}

// creating history string for Graph mode
func (h *calcHistory) createHistoryGraph(in ModelsInput, out ModelsOutput) {

	h.tempHistoryItem.Mode = "graph"
	h.tempHistoryItem.Equation = in.ModelGraphData.EqualValue
	h.tempHistoryItem.XFrom = in.ModelGraphData.XFromStr
	h.tempHistoryItem.XTo = in.ModelGraphData.XToStr
	h.tempHistoryItem.YFrom = in.ModelGraphData.YFromStr
	h.tempHistoryItem.YTo = in.ModelGraphData.YToStr
	h.tempHistoryItem.Entrys = fmt.Sprintf("X{%s .. %s} Y{%s .. %s} ",
		in.ModelGraphData.XFromStr,
		in.ModelGraphData.XToStr,
		in.ModelGraphData.YFromStr,
		in.ModelGraphData.YToStr)
	if out.Err {
		h.tempHistoryItem.Result = "error"
	} else {
		h.tempHistoryItem.Result = out.ModelGraphResult.ResultStr
	}
}

//----------------------------------------Creating historyItem END
//----------------------------------------Handle history START

// adding history string to history base
func (h *calcHistory) addHistoryString() {
	h.historyData = append(h.historyData, h.tempHistoryItem) // adding historyItem to history Base
}

// creating Json data from
func (h *calcHistory) createHistoryJson() (data []byte) {
	data, _ = json.MarshalIndent(h.historyData, "", "    ") // creating Json
	return data
}

// writing entire history base to file
func (h *calcHistory) writeHistoryJson(data []byte) {
	historyFd, err := os.OpenFile(h.fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(h.fileName, data, 0777) // write json([]byte) to file
	historyFd.Close()
}

// reading entire history from file
func readHistoryJson(fileName string) (hdata []d.HistoryItem) {

	historyFd, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}
	dataFromFile, err := os.ReadFile(fileName)
	err = json.Unmarshal(dataFromFile, &hdata)

	historyFd.Close()
	return hdata
}

//----------------------------------------Handle history END
