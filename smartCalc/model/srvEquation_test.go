package model

import (
	"reflect"
	"testing"
)

var (
	testCasesCheckBrackets = []struct {
		val    string
		expect bool
	}{
		{"({[]})", true},
		{"({}))", false},
		{"({)}", false},
	}

	testCasesStartCalculate = []struct {
		val    string
		expect float64
	}{
		{"5*6+(2-9)", 23},
		{"6 + 3 * (1 + 4 * 5) * 2 =", 132},
		{"1 / 2", 0.5},
		{" 5 + (-2 + 3)", 6},
		{" 5 + (+2 + 3)", 10},
		{" 5 + ( -2 + 3)", 6},
		{" 1.5 + 1.5", 3},
		{"-1.5 + (-1.5)", -3},
		{"0.66E+4 + 2", 6602},
		{"0.66E+4 + 2e+2 + 300", 7100},
		{"1.1e+10 + 1.1e+10", 2.2e+10},
		{"2.4e+10 - 2.4e+10", 0},
		{"-2 + 3", 1},
		{"(5.2e+4 + sin(0.1) - 10) + (-0.2)", 51989.80174532837},
	}

	testCasesCheckUnary = []struct {
		enter  string
		expect string
	}{
		{"-2", "(0-2)"},
		{"-2*3", "(0-2)*3"},
		{"-2*(-3)", "(0-2)*((0-3))"},
		{"34/(-2)", "34/((0-2))"},
		{"    tan   (   1   )", "tan(1)"},
		{"0.66 E +     4 + 2e +  2 + 300", "0.66e+4+2e+2+300"},
	}

	testCasesInsertSpaces = []struct {
		enter  string
		expect string
	}{
		{"(0-2)", " ( 0 - 2 ) "},
		{"( 0-2 )", " ( 0 - 2 ) "},
		{"( 0-.5 )", " ( 0 - 0.5 ) "},
		{".4*( 0-.345 )", "0.4 *  ( 0 - 0.345 ) "},
		{"(0-2)*3", " ( 0 - 2 )  * 3"},
		{"(0-2)*((0-3))", " ( 0 - 2 )  *  (  ( 0 - 3 )  ) "},
		{"0.66e+4+2e+2+300", "0.66e+4 + 2e+2 + 300"},
	}

	testCasesPoland = []struct {
		val    []string
		expect []string
	}{
		{[]string{"654", "+", "3", "*", "(", "1", "+", "4", "*", "5", ")", "*", "2"}, []string{"654", "3", "1", "4", "5", "*", "+", "*", "2", "*", "+"}},
		{[]string{"5", "*", "6", "+", "(", "2", "-", "9", ")"}, []string{"5", "6", "*", "2", "9", "-", "+"}},
	}

	testCasesCalculate = []struct {
		enter  []string
		expect float64
	}{
		{[]string{"2", "0", "cos", "*", "0", "cos", "+"}, 3},
		{[]string{"2", "90", "sin", "*", "0", "sin", "+"}, 2},
		{[]string{"90", "sin"}, 1},
		{[]string{"2", "1", "asin", "*", "1", "asin", "+"}, 270},
		{[]string{"1", "asin"}, 90},
		{[]string{"2", "0", "acos", "*", "-1", "acos", "+"}, 360},
		{[]string{"0", "acos"}, 90},
		{[]string{"25", "sqrt"}, 5},
		{[]string{"25", "sqrt", "25", "sqrt", "*", "36", "sqrt", "+"}, 31},
		{[]string{"1", "tan"}, 0.017455064928217585},
		{[]string{"180", "tan", "180", "tan", "*", "45", "tan", "+", "6", "+"}, 7},
		{[]string{"1", "atan"}, 45},
		{[]string{"1", "atan", "1", "atan", "+", "0", "atan", "-"}, 90},
		{[]string{"7.38905609893065", "ln"}, 2},
		{[]string{"4", "3.38905609893065", "+", "ln"}, 2},
		{[]string{"4", "sqrt", "7.38905609893065", "ln", "+"}, 4},
		{[]string{"100", "log"}, 2},
		{[]string{"1000", "log"}, 3},
		{[]string{"1000", "log", "10", "log", "+", "sqrt"}, 2},
		{[]string{"1000", "8999", "+", "10", "log", "+", "log"}, 4},
		{[]string{"1", "2", "+"}, 3},
		{[]string{"3", "1", "-"}, 2},
		{[]string{"4", "5", "*"}, 20},
		{[]string{"24", "6", "/"}, 4},
		{[]string{"5", "3", "mod"}, 2},
		{[]string{"2", "8", "^"}, 256},
	}
)

func TestNewEquation(t *testing.T) {
	e := createMInputStructEquation("34-23/2")
	er := *NewEquation(e)
	if reflect.TypeOf(er).String() != "model.equationModel" {
		t.Errorf("Creating object incorrect - expected: %v; actual: %v", "model.equationModel", reflect.TypeOf(er).String())
	} else if er.equation != "34-23/2" {
		t.Errorf("incorrect equation - expected: %v; actual: %v", "34-23/2", er.equation)
	}
}

// Implementing GetResult interface for equationModel
func TestGetResultEquation(t *testing.T) {

	er := NewEquation(createMInputStructEquation("34-24/2")).GetResult()
	if er.Mode != 0 || er.ModelEquationResult.Mode != 0 {
		t.Errorf("Equations ModelsOutput-Mode incorrect - expected: %v; actual: %v, %v", 0, er.Mode, er.ModelEquationResult.Mode)
	} else if er.ModelEquationResult.ResultStr != "22" {
		t.Errorf("Equations ModelsOutput-ResultStr incorrect - expected: %v; actual: %v", "12", er.ModelEquationResult.ResultStr)
	}
	// check fo errors
}

// insert spaces between expressions and zero before unary (-2 > 0-2)
func TestPrepareString(t *testing.T) {
	var e equationModel
	if e.prepareString("34-24/(-2)") != "34 - 24 /  (  ( 0 - 2 )  ) " {
		t.Errorf("Equations prepareSring incorrect - expected: |%v|; actual: |%v|", "34 - 24 /  (  ( 0 - 2 )  ) ", e.prepareString("34-24/(-2)"))
	}
}

// checking for empty string and brackets
func TestOnlyCheck(t *testing.T) {
	er := NewEquation(createMInputStructEquation("34-24/2"))
	if er.Checked, er.err = er.onlyCheck(); er.err != nil {
		t.Errorf("Equations onlyCheck incorrect - expected: |%v|; actual: |%v|", false, er.err == nil)
	}
	er = NewEquation(createMInputStructEquation("34-24/)(2"))
	if er.Checked, er.err = er.onlyCheck(); er.err == nil {
		t.Errorf("Equations onlyCheck incorrect - expected: |%v|; actual: |%v|", true, false)
	}
}

// calculate prepared string
func TestOnlyCalculate(t *testing.T) {
	er := NewEquation(createMInputStructEquation("34-24/2"))
	rez, err := er.onlyCalculate(er.prepareString(er.equation))
	if rez != 22 || err != nil {
		t.Errorf("Equations onlyCalculate incorrect - expected: rez=|%v|, err=|%v|; actual: rez=|%v|, err=|%v|", rez, err, rez, err)
	}
	er.equation = ("34-24/0")
	_, err = er.onlyCalculate(er.prepareString(er.equation))
	if err == nil {
		t.Errorf("Equations onlyCalculate incorrect - expected: err=|%v|; actual: err=|%v|", true, false)
	}
	er.equation = ("42e-2*350") // rez = 147
	rez, err = er.onlyCalculate(er.prepareString(er.equation))
	if err != nil || rez != 147 {
		t.Errorf("Equations onlyCalculate incorrect - expected: err=|%v|; actual: err=|%v|", nil, err)
	}
	_, err = er.onlyCalculate("")
	if err == nil {
		t.Errorf("Equations onlyCalculate incorrect - expected: err=|%v|; actual: err=|%v|", true, false)
	}
}

// check for brackets balance
func TestCheckBrackets(t *testing.T) {
	var er equationModel
	for _, testCase := range testCasesCheckBrackets {
		actual := er.checkBrackets(testCase.val)
		if actual != testCase.expect {
			t.Errorf("Result was incorrect, expected: %v, actual: %v\n", testCase.expect, actual)
		}
	}
}

// replace unary expressions with 0 (-2 > 0-2)
func TestReplaceUnary(t *testing.T) {
	var er equationModel
	for _, testCase := range testCasesCheckUnary {
		actual := er.replaceUnary(testCase.enter)
		if actual != testCase.expect {
			t.Errorf("Result ReplaceUnary was incorrect, expected: |%v|, actual: |%v|\n", testCase.expect, actual)
		}
	}
}

// insert spaces between expressions
func TestInsertSpases(t *testing.T) {
	var er equationModel
	for _, testCase := range testCasesInsertSpaces {
		actual := er.insertSpases(testCase.enter)
		if actual != testCase.expect {
			t.Errorf("Result InsertSpaces was incorrect, expected: |%v|, actual: |%v|\n", testCase.expect, actual)
		}
	}
}

// Convert infix to Poland notation
func TestToPolandNotation(t *testing.T) {
	var er equationModel
	for _, testCase := range testCasesPoland {
		actual := er.toPolandNotation(testCase.val)
		if !reflect.DeepEqual(actual, testCase.expect) {
			t.Errorf("Result was incorrect, expected: %v, actual: %v\n", testCase.expect, actual)
		}
	}
}

// Calculate expression in poland notation
func TestCalculateEquation(t *testing.T) {
	var er equationModel
	for _, testCase := range testCasesCalculate {
		actual, err := er.calculate(testCase.enter)
		if err == nil && actual != testCase.expect {
			t.Errorf("Result was incorrect, exected: %v, actual: %v\n", testCase.expect, actual)
		}
	}
}

// Common function of calculate with input prepared string (with unary and spaces)
func TestStartCalculate(t *testing.T) {
	var er equationModel
	for _, testCase := range testCasesStartCalculate {
		actual, err := er.startCalculate(er.prepareString(testCase.val))
		if err == nil && actual != testCase.expect {
			t.Errorf("Result was incorrect, expected: |%v|, actual: |%v|\n", testCase.expect, actual)
		}
	}
}
