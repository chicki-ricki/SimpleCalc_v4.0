package controllers

import (
	"encoding/json"
	"fmt"
	"strings"

	m "smartCalc/model"
	t "smartCalc/tools"
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
		strData += `{"mode":"equal","equation":"`
		strData += input.ModelEqualData.EqualValue + "\","
		strData += `"result":"`
		strData += outputs[1] + "\","
		strData += `"xEqual":"`
		strData += input.ModelEqualData.XEqualStr + "\","
		strData += `"xFrom":"","xTo":"","yFrom":"","yTo":""}`
	case 2:
		strData += `{"mode":"graph","equation":"`
		strData += input.ModelGraphData.EqualValue + "\","
		strData += `"result":"`
		strData += outputs[1] + "\","
		strData += `"xEqual":"","xFrom":"`
		strData += input.ModelGraphData.XFromStr + "\","
		strData += `"xTo":"`
		strData += input.ModelGraphData.XToStr + "\","
		strData += `"yFrom":"`
		strData += input.ModelGraphData.YFromStr + "\","
		strData += `"yTo":"`
		strData += input.ModelGraphData.YToStr + "\"}"
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
