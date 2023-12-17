package model

import (
	"reflect"
	"testing"
)

var ()

// Creating New CalcModel object
func TestNewCalcModel(t *testing.T) {
	r := *NewCalcModel(configCalc)
	s := reflect.TypeOf(r).String()
	if s != "model.calcModel" {
		t.Errorf("Creating object incorrect expected: %v: actual:%v", "model.calcModel", s)
	} else if reflect.TypeOf(r.history).String() != "model.calcHistory" {
		t.Errorf("Creating object incorrect expected: %v: actual:%v", "model.calcHistory", reflect.TypeOf(r.history).String())
	}
}

//-------------------Implementing Models Interface

func TestGetCalcResult(t *testing.T) {
	in := createMInputStructEqual("x^2", "4", 4)
	calc := NewCalcModel(configCalc)
	out := calc.GetCalcResult(in)
	if reflect.TypeOf(out).String() != "model.ModelsOutput" {
		t.Errorf("Creating object incorrect expected: %v: actual:%v", "model.ModelsOutput", reflect.TypeOf(out).String())
	}
	if out.Mode != 1 {
		t.Errorf("Incorrect Mode, expected: %v: actual:%v", 1, out.Mode)
	}
	if out.Err {
		t.Errorf("Incorrect Err, expected: %v: actual:%v", "false", out.Err)
	}
}

func TestCleanHistoryModel(t *testing.T) {
	in := createMInputStructEqual("x^2", "4", 4)
	calc := NewCalcModel(configCalc)
	calc.GetCalcResult(in)
	lenght := len(calc.history.historyData)
	calc.CleanHistory()
	if lenght > 0 && len(calc.history.historyData) != 0 {
		t.Errorf("Clean History incorrect - expected len=: %v ; actual: %v", 0, len(calc.history.historyData))
	}
}

func TestGetHistory(t *testing.T) {
	calc := NewCalcModel(configCalc)
	out := calc.GetHistory()
	if reflect.TypeOf(out).String() != "[]domains.HistoryItem" {
		t.Errorf("Creating object incorrect expected: %v: actual:%v", "[]domains.HistoryItem", reflect.TypeOf(out).String())
	}
}
