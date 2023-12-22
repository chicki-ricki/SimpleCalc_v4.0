package controllers

import (
	"encoding/json"
	"fmt"
	"strings"
	m "smartCalc/model"
	t "smartCalc/tools"
	// d "smartCalc/domains"
)

func loadHistoryFromModel() []byte {
	historyFromModel := m.ModelCalc.GetHistory()
	data, _ := json.MarshalIndent(historyFromModel, "", "    ")
	strData := "9_"
	strData += string(data)
	ret := []byte(strData)
	return ret
}

func lastHistory(input m.ModelsInput, output string) []byte {

    /* 
    {
        "mode": "calc",
        "equation": "6*9",
        "result": "54",
        "entrys": "",
        "xEqual": "",
        "xFrom": "",
        "xTo": "",
        "yFrom": "",
        "yTo": ""
    }
	*/

	strData := "8_"
	outputs := strings.Fields(output)
	switch input.Mode {
	case 0:
		strData += `{"mode":"calc","equation":"`
		strData += input.ModelEquationData.EqualValue + "\","
		strData += `"result":"`
		strData += outputs[1] + "\","
		strData += `"xEqual":"","xFrom":"","xTo":"","yFrom":"","yTo":""}`

	case 1:
		t.Clg.Debug(fmt.Sprint("_lastHistory_ data equal input:", input.ModelEqualData))
		t.Clg.Debug(fmt.Sprint("_lastHistory_ data equal output:", output))
	case 2:
		t.Clg.Debug(fmt.Sprint("_lastHistory_ data graph input:", input.ModelEqualData))
		t.Clg.Debug(fmt.Sprint("_lastHistory_ data graph output:", output))
	}
		t.Clg.Debug(fmt.Sprint("_lastHistory_ ret: ", strData))

	ret := []byte(strData)
	return ret
}

func clearHistory() []byte {
	m.ModelCalc.CleanHistory()
	strData := "7_"
	ret := []byte(strData)
	return ret
}
