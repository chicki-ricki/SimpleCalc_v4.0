package model

import (
	"fmt"
	"math"
	d "smartCalc/domains"
	"strconv"
	"strings"
)

type equalModel struct {
	equalString string
	xEqualStr   string
	xEqual      float64
	result      float64
	equation    equationModel
}

func NewEqual(in ModelsInput) *equalModel {
	var equal equalModel
	d.DbgPrint(fmt.Sprint("NewEquaL", in))
	equal.equation.equation = in.ModelEqualData.EqualValue
	equal.equalString = in.ModelEqualData.EqualValue
	equal.xEqualStr = in.ModelEqualData.XEqualStr
	equal.xEqual = in.ModelEqualData.XEqual
	return &equal
}

func (e *equalModel) setError(out *ModelsOutput) *ModelsOutput {
	out.Err = true
	out.ModelEqualResult.Err = true
	out.ModelEqualResult.ResultStr = "Error"
	return out
}

// Implementing request interface for equalModel
func (e *equalModel) GetResult() (out ModelsOutput) {
	out.ModelEqualResult.Mode = 1
	out.Mode = 1

	str, err := e.equation.onlyCheck()
	if err != nil {
		return *e.setError(&out)
	}

	if e.xEqual, err = strconv.ParseFloat(e.xEqualStr, 64); err != nil {
		return *e.setError(&out)
	}

	e.result, err = e.equation.onlyCalculate(e.equation.prepareString(e.equalPrepareString(str)))
	if err != nil {
		return *e.setError(&out)
	}

	out.ModelEqualResult.ResultStr = strconv.FormatFloat(e.result, 'f', -1, 64)
	return
}

// Calculate checked string with inputed x
func (e *equalModel) calculate(str string, x float64) (pixel SCoordinates) {
	var err error
	pixel.x = x
	pixel.y, err = e.equation.onlyCalculate(e.equation.prepareString(strings.ReplaceAll(str, "x", strconv.FormatFloat(x, 'f', -1, 64))))
	if err != nil || math.IsNaN(pixel.y) || math.IsInf(pixel.y, 0) {
		pixel.err = true
		return
	}
	return
}

// Adding staples and Multi for x and brackets in equal
func (e *equalModel) addStaplesForX(str string) string {

	// handle sign before first x at the begining of string
	templength := len(str)
	if templength > 1 && (string(str[0:2]) == "-x" || string(str[0:2]) == "+x" || string(str[0:2]) == "*x" ||
		string(str[0:2]) == "/x") {
		str = fmt.Sprint("0" + string(str[0:]))
	}

	// handle x at the begining of string
	templength = len(str)
	if string(str[0:1]) == "x" {
		switch templength == 1 {
		case false:
			str = fmt.Sprint("(x)" + string(str[1:]))
		case true:
			str = "x"
			return str
		}
	}

	// handle x at the tail of string
	templength = len(str)
	if string(str[templength-1:templength]) == "x" {
		switch strings.Contains("0123456789x)", string(str[templength-2:templength-1])) {
		case true:
			str = fmt.Sprint(string(str[0:templength-1]) + "*(x)")
		case false:
			str = fmt.Sprint(string(str[0:templength-1]) + "(x)")
		}
	}

	// handle x at the middle of string
	for i := len(str) - 2; i >= 1; i-- {
		if string(str[i:i+1]) == "x" && strings.Contains("0123456789x)", string(str[i-1:i])) {
			str = fmt.Sprint(str[0:i] + "*(x)" + string(str[i+1:]))
		} else if string(str[i:i+1]) == "x" && !strings.Contains("0123456789", string(str[i-1:i])) {
			str = fmt.Sprint(string(str[0:i]) + "(x)" + string(str[i+1:]))
		} else if string(str[i:i+1]) == "(" && strings.Contains("0123456789x)", string(str[i-1:i])) {
			str = fmt.Sprint(str[0:i] + "*(" + string(str[i+1:]))
		}
	}

	// handle ")num" situation
	for i := len(str) - 2; i >= 1; i-- {
		if strings.Contains("0123456789", string(str[i:i+1])) && string(str[i-1:i]) == ")" {
			str = fmt.Sprint(str[0:i] + "*" + string(str[i:]))
		}
	}

	return str
}

// insert value of x and insert staples fo x
func (e *equalModel) equalPrepareString(str string) string {
	return strings.ReplaceAll(e.addStaplesForX(str), "x", e.xEqualStr)
}
