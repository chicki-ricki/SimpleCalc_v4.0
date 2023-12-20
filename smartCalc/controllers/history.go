package controllers

import (
	"encoding/json"
	// "fmt"
	m "smartCalc/model"
	// d "smartCalc/domains"
)

func loadHistoryFromModel() []byte {
	historyFromModel := m.ModelCalc.GetHistory()
	data, _ := json.MarshalIndent(historyFromModel, "", "    ")
	strData := "9_"
	strData += string(data)
	ret := []byte(strData)
	// fmt.Println("ret:", string(ret))
	return ret
}
