package controllers

import (
	"fmt"
	m "smartCalc/model"
	t "smartCalc/tools"
	"strings"
)

// const FILENAME string = "history.json"

type convert struct {
}

// create ModelsInput with graph
func (c *convert) copyGraphForModel(inArray []string) (input m.ModelsInput, err bool) {
	input.ModelGraphData.Mode = 2
	input.Mode = 2
	if input.ModelGraphData.EqualValue = inArray[5]; strings.EqualFold(input.ModelGraphData.EqualValue, "") {
		return input, true
	} else if input.ModelGraphData.XFromStr = inArray[1]; strings.EqualFold(input.ModelGraphData.XFromStr, "") {
		input.ModelGraphData.XFromStr = "0"
	} else if input.ModelGraphData.XToStr = inArray[2]; strings.EqualFold(input.ModelGraphData.XToStr, "") {
		input.ModelGraphData.XToStr = "0"
	} else if input.ModelGraphData.YFromStr = inArray[3]; strings.EqualFold(input.ModelGraphData.YFromStr, "") {
		input.ModelGraphData.YFromStr = "0"
	} else if input.ModelGraphData.YToStr = inArray[4]; strings.EqualFold(input.ModelGraphData.YToStr, "") {
		input.ModelGraphData.YToStr = "0"
	}
	if strings.EqualFold(input.ModelGraphData.XFromStr, input.ModelGraphData.XToStr) ||
		strings.EqualFold(input.ModelGraphData.YFromStr, input.ModelGraphData.YToStr) {
		return input, true
	}
	return
}

// create ModelsInput with equal
func (c *convert) copyEqualForModel(inArray []string) (input m.ModelsInput, err bool) {
	input.ModelEqualData.Mode = 1
	input.Mode = 1
	t.DbgPrint(fmt.Sprint("copyEqualForModel inArray:", inArray))
	t.DbgPrint(fmt.Sprint("copyEqualForModel EqualValue", inArray[2]))
	t.DbgPrint(fmt.Sprint("copyEqualForModel XEqualStr", inArray[1]))
	if input.ModelEqualData.EqualValue = inArray[2]; strings.EqualFold(input.ModelEqualData.EqualValue, "") {
		return input, true
	} else if input.ModelEqualData.XEqualStr = inArray[1]; strings.EqualFold(input.ModelEqualData.XEqualStr, "") {
		input.ModelEqualData.XEqualStr = "0"
		return input, false
	}
	return
}

// create ModelsInput with equation
func (c *convert) copyEquationForModel(inArray []string) (input m.ModelsInput, err bool) {
	t.DbgPrint(fmt.Sprint("Convert to equation"))
	input.ModelEquationData.Mode = 0
	input.Mode = 0
	if input.ModelEquationData.EqualValue = inArray[1]; inArray[1] == "" {
		return input, true
	}
	return input, false
}

// converted interface to ModelsInput for Model
func (c *convert) UIToModel(in string) (m.ModelsInput, bool) {

	t.DbgPrint(fmt.Sprint("in:", in))
	if in == "" {
		return m.ModelsInput{}, true
	}

	inArray := strings.Fields(in)
	for i, val := range inArray {
		fmt.Printf("inArray i = %d, str = %s\n", i, val)
	}
	t.DbgPrint(fmt.Sprint("inArray:", inArray))
	t.DbgPrint(fmt.Sprint("len(inArray):", len(inArray)))
	if len(inArray) < 2 {
		return m.ModelsInput{}, true
	}

	switch inArray[0] {
	case "calculate":
		t.DbgPrint(fmt.Sprint("Choice Calculate"))
		if len(inArray) == 2 {
			return c.copyEquationForModel(inArray)
		}
	case "equal":
		t.DbgPrint(fmt.Sprint("Choice Equal"))
		if len(inArray) == 3 {
			return c.copyEqualForModel(inArray)
		}
	case "graph":
		t.DbgPrint(fmt.Sprint("Choice Graph"))
		if len(inArray) == 6 {
			return c.copyGraphForModel(inArray)
		}
	}

	return m.ModelsInput{}, true
}

// converted modelsOutput to interface for View
func (c *convert) ModelToUI(output m.ModelsOutput) string {

	switch output.Mode {
	case 0:
		t.DbgPrint(fmt.Sprint("EOUT:", output.ModelEquationResult))
		return fmt.Sprint(output.ModelEquationResult.Mode, " ", output.ModelEquationResult.ResultStr)
	case 1:
		t.DbgPrint(fmt.Sprint("EQLOUT:", output.ModelEqualResult))
		return fmt.Sprint(output.ModelEqualResult.Mode, " ", output.ModelEqualResult.ResultStr)
	case 2:
		t.DbgPrint(fmt.Sprint("GROUT:", output.ModelGraphResult))
		return fmt.Sprint(output.ModelGraphResult.Mode, " ", output.ModelGraphResult.ResultStr)
	}

	return "0 errorPr"
}
