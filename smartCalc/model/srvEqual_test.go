package model

import (
	"fmt"
	"testing"
)

var (
	in ModelsInput
	e  equalModel

	testCasesGetResult = []struct {
		s    string
		xStr string
		rStr string
	}{
		{"x^2 + 16x", "2", "36"},
		{"4x^2+3x-2", "-2", "8"},
		{"log(x^(2)) + 2x", "100", "204"},
	}
	testCasesGetResult2 = []struct {
		s    string
		xStr string
		rStr string
	}{
		{"x^2 + 16x", "2ecofds", "36"},
		{"4x^2+3x/(2-2)-2", "-2", "8"},
		{"log(x^(2)) + 2x", "ln100", "204"},
	}
)

func createEqualStruct(s string, xStr string) (r equalModel) {
	r.equation.equation = s
	r.equalString = s
	r.xEqualStr = xStr
	return
}

func createMInputStructEqual(s string, xStr string, x float64) (in ModelsInput) {
	in.ModelEqualData.EqualValue = s
	in.ModelEqualData.XEqualStr = xStr
	in.ModelEqualData.XEqual = x
	in.ModelEqualData.Mode = 1
	in.Mode = 1
	return
}

func createMInputStructEquation(s string) (in ModelsInput) {
	in.ModelEquationData.EqualValue = s
	return
}

func createMInputStructGraph(s string, xFrom string, xTo string, yFrom string, yTo string) (in ModelsInput) {
	in.ModelGraphData.EqualValue = s
	in.ModelGraphData.XFromStr = xFrom
	in.ModelGraphData.XToStr = xTo
	in.ModelGraphData.YFromStr = yFrom
	in.ModelGraphData.YToStr = yTo
	in.ModelGraphData.Mode = 2
	in.Mode = 2
	return
}

func TestCalculate(t *testing.T) {
	pixel := e.calculate(e.addStaplesForX("4x-2"), 2)
	if pixel.x != 2 || pixel.y != 6 || pixel.err {
		t.Errorf("Result is incorrect, %v", pixel)
	}
	pixel = e.calculate(e.addStaplesForX("x^2 - 3"), 2)
	if pixel.x != 2 || pixel.y != 1 || pixel.err {
		t.Errorf("Result is incorrect, %v", pixel)
	}
	pixel = e.calculate(e.addStaplesForX("1/x"), 0)
	if pixel.x != 0 || !pixel.err {
		fmt.Printf("pixel: %v", pixel)
		t.Errorf("Result is incorrect, %v", pixel)
	}
}

func TestEqualPrepareString(t *testing.T) {
	e.xEqualStr = "0.2"
	str := e.equalPrepareString("5x^2+x-4x")
	if str != "5*(0.2)^2+(0.2)-4*((0.2))" {
		t.Errorf("Result was incorrect, expected %s, actual: %s", "5*(0.2)^2+(0.2)-4*((0.2))", str)
	}
}

func TestGetResult(t *testing.T) {
	var r equalModel
	for _, tc := range testCasesGetResult {
		r = createEqualStruct(tc.s, tc.xStr)
		out := r.GetResult()
		if out.Mode != 1 || out.Err == true || out.ModelEqualResult.ResultStr != tc.rStr {
			t.Errorf("Result was incorrect, %v", out)
		}
	}
}

func TestGetResult2(t *testing.T) {
	var r equalModel
	for _, tc := range testCasesGetResult2 {
		r = createEqualStruct(tc.s, tc.xStr)
		out := r.GetResult()
		if out.Err != true {
			t.Errorf("Result was incorrect, %v", out)
		}
	}
}

func TestNewEqual(t *testing.T) {
	in.ModelEqualData.EqualValue = "4x-5"
	in.ModelEqualData.XEqualStr = "15"
	in.ModelEqualData.XEqual = 15
	e = *NewEqual(in)
	if e.equation.equation != in.ModelEqualData.EqualValue ||
		e.equalString != in.ModelEqualData.EqualValue ||
		e.xEqualStr != in.ModelEqualData.XEqualStr ||
		e.xEqual != in.ModelEqualData.XEqual {
		t.Errorf("Creating equal object incorrect expected: %v: actual:%v", in.ModelEqualData, e)
	}

}

func TestAddStaplesForX(t *testing.T) {
	var e *equalModel
	if result := e.addStaplesForX("45(3+6)(3-6)"); result != "45*(3+6)*(3-6)" {
		t.Errorf("Result was incorrect, expected %s, actual: %s", "45(3+6)(3-6)", result)
	}
	if result := e.addStaplesForX("45x(3+6x)(x-6)"); result != "45*(x)*(3+6*(x))*((x)-6)" {
		t.Errorf("Result was incorrect, expected %s, actual: %s", "45(3+6)(3-6)", result)
	}
	if result := e.addStaplesForX("-x+45x(3+6x)(x-6)"); result != "0-(x)+45*(x)*(3+6*(x))*((x)-6)" {
		t.Errorf("Result was incorrect, expected %s, actual: %s", "45(3+6)(3-6)", result)
	}
	if result := e.addStaplesForX("45x(3+6x)(x-6)56"); result != "45*(x)*(3+6*(x))*((x)-6)*56" {
		t.Errorf("Result was incorrect, expected %s, actual: %s", "45(3+6)(3-6)", result)
	}
}

func BenchmarkGetResult(b *testing.B) {
	var r equalModel
	for _, tc := range testCasesGetResult {
		b.Run(fmt.Sprintf("%v", tc), func(b *testing.B) {
			r = createEqualStruct(tc.s, tc.xStr)
			r.GetResult()
		})
	}
}
