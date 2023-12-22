package controllers

import (
	"encoding/json"
	"fmt"
	"strings"
	m "smartCalc/model"
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

func lastHistory(input m.ModelsInput, output m.ModelsOutput) []byte {
	// historyFromModel := m.ModelCalc.GetHistory()
	// data, _ := json.MarshalIndent(historyFromModel, "", "    ")
	// fmt.Println("lenght json:", len(data))
	// fmt.Println("last elem in json:", string(data[len(data) - 1]))

	strData := "8_"

	fmt.Println("input.Mode: ", input.Mode);
	switch input.Mode {
	case 0:
		fmt.Println("data calculate:", input.ModelEquationData.EqualValue)
		strData += "{\"mode\":\"calc\"," + "\"equation\":\"" + input.ModelEquationData.EqualValue + "\"}"

	case 1:
		fmt.Println("data equal:", input.ModelEqualData)
	case 2:
		fmt.Println("data graph:", input.ModelGraphData)
	}
	fmt.Println("ret: ", strData)



	ret := []byte(strData)
	return ret
}

func historyParseAddition(input string) {
	args := strings.Fields(input)
	for _, arg := range args {
		fmt.Println("arg: ", arg)
	}
}
