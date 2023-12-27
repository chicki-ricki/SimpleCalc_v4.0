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

func makeCalcResponse(input m.ModelsInput, outputs []string) string {
	out := ""
	out += `{"mode":"calc","equation":"`
	out += input.ModelEquationData.EqualValue + "\","
	out += `"result":"`
	out += outputs[1] + "\","
	out += `"xEqual":"","xFrom":"","xTo":"","yFrom":"","yTo":""}`
	return out
}

func makeEqualResponse(input m.ModelsInput, outputs []string) string {
	out := ""
	out += `{"mode":"equal","equation":"`
	out += input.ModelEqualData.EqualValue + "\","
	out += `"result":"`
	out += outputs[1] + "\","
	out += `"xEqual":"`
	out += input.ModelEqualData.XEqualStr + "\","
	out += `"xFrom":"","xTo":"","yFrom":"","yTo":""}`
	return out
}

func makeGraphResponse(input m.ModelsInput, outputs []string) string {
	out := ""
	out += `{"mode":"graph","equation":"`
	out += input.ModelGraphData.EqualValue + "\","
	out += `"result":"`
	out += outputs[1] + "\","
	out += `"xEqual":"","xFrom":"`
	out += input.ModelGraphData.XFromStr + "\","
	out += `"xTo":"`
	out += input.ModelGraphData.XToStr + "\","
	out += `"yFrom":"`
	out += input.ModelGraphData.YFromStr + "\","
	out += `"yTo":"`
	out += input.ModelGraphData.YToStr + "\"}"
	return out
}

func lastHistory(input m.ModelsInput, output string) []byte {
	strData := "8_"
	outputs := strings.Fields(output)
	switch input.Mode {
	case 0:
		strData += makeCalcResponse(input, outputs)
	case 1:
		strData += makeEqualResponse(input, outputs)
	case 2:
		strData += makeGraphResponse(input, outputs)
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
