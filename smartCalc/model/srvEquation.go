package model

import (
	"errors"
	"fmt"
	"math"
	t "smartCalc/tools"
	"strconv"
	"strings"
	"unicode"
)

type equationModel struct {
	err       error
	equation  string
	Checked   string
	prepared  string
	ResultStr string
	Result    float64
}

func NewEquation(in ModelsInput) *equationModel {
	t.DbgPrint(fmt.Sprint("NewEquation", in))
	return &equationModel{
		equation: in.ModelEquationData.EqualValue,
	}
}

func (e *equationModel) setError(out *ModelsOutput) *ModelsOutput {
	out.Err = true
	out.ModelEquationResult.Err = true
	out.ModelEquationResult.ResultStr = "Error"
	return out
}

// Implementing GetResult interface for equationModel
func (e *equationModel) GetResult() (out ModelsOutput) {
	// var rez calcViewResult
	out.Mode = 0
	out.ModelEquationResult.Mode = 0

	if e.Checked, e.err = e.onlyCheck(); e.err != nil {
		return *e.setError(&out)
	}
	if e.Result, e.err = e.onlyCalculate(e.prepareString(e.Checked)); e.err != nil {
		return *e.setError(&out)
	}
	out.ModelEquationResult.ResultStr = strconv.FormatFloat(e.Result, 'f', -1, 64)
	return
}

// insert spaces between expressions and zero before unary (-2 > 0-2)
func (e *equationModel) prepareString(str string) string {
	return e.insertSpases(e.replaceUnary(str))
}

// checking for empty string and brackets
func (e *equationModel) onlyCheck() (string, error) {
	if e.equation != "" && e.checkBrackets(e.equation) {
		e.Checked = e.equation
	} else {
		e.err = errors.New("Invalid brackets or empty string")
	}
	return e.Checked, e.err
}

// calculate prepared string
func (e *equationModel) onlyCalculate(str string) (rez float64, err error) {
	if str != "" {
		if rez, err = e.startCalculate(str); err != nil {
			e.err = err
		}
	} else {
		e.err = errors.New("Empty request")
	}
	return rez, e.err
}

// check for brackets balance
func (e *equationModel) checkBrackets(str string) bool {
	var stack []string
	for _, char := range str {
		lenght := len(stack) - 1
		switch {
		case char == '(':
			stack = append(stack, ")")
		case char == '[':
			stack = append(stack, "]")
		case char == '{':
			stack = append(stack, "}")
		case char == ')' || char == '}' || char == ']':
			if len(stack) == 0 || stack[lenght] != string(char) {
				return false
			}
			stack[lenght] = ""
			stack = stack[:lenght]
		default:
			break
		}
	}
	if len(stack) == 0 {
		return true
	} else {
		return false
	}
}

// replace unary expressions with 0 (-2 > 0-2)
func (e *equationModel) replaceUnary(str string) string {
	var retStr string
	var flagIn bool
	s := strings.ToLower(str)
	for i, char := range s {
		if char == ' ' {
			continue
		}
		if len(retStr) == 0 && (char == '-' || char == '+') {
			retStr += "(0" + string(char)
			flagIn = true
		} else if (char == '+' || char == '-') && (len(retStr) > 1 && retStr[len(retStr)-1:] == "(") {
			retStr += "(0" + string(char)
			flagIn = true
		} else if flagIn && strings.Contains(")(^+-*/", string(char)) {
			retStr += ")" + string(char)
			flagIn = false
		} else if flagIn && len(str) == i+1 {
			retStr += string(char) + ")"
			flagIn = false
		} else {
			retStr += string(char)
		}
	}
	return retStr
}

// insert spaces between expressions
func (e *equationModel) insertSpases(str string) string {
	var retStr string
	s := strings.ToLower(str)
	for _, char := range s {
		if char == ' ' {
			continue
		}
		if strings.Contains(")(^+-*/", string(char)) {
			retStr += " " + string(char) + " "
		} else if char == 'm' {
			retStr += " " + string(char)
		} else if char == 'd' {
			retStr += string(char) + " "
		} else if unicode.IsDigit(char) || unicode.IsLetter(char) || char == '.' || (len(retStr) > 1 && retStr[len(retStr)-1:] == "e") {
			retStr += string(char)
		} else {
			retStr += string(char) + " "
		}
	}
	// t.DbgPrint(fmt.Sprint("insertSpaces - after spaces added:", retStr))

	if string(retStr[0:1]) == "." && strings.Contains("0123456789", string(retStr[1:2])) {
		retStr = fmt.Sprint("0" + retStr)
	}

	for i := 0; i < len(retStr)-2; i++ {
		if string(retStr[i:i+1]) == " " && string(retStr[i+1:i+2]) == "." && strings.Contains("0123456789", string(retStr[i+2:i+3])) {
			retStr = fmt.Sprint(retStr[0:i+1] + "0" + string(retStr[i+1:]))
		}
	}
	// t.DbgPrint(fmt.Sprint("insertSpaces - after zero added:", retStr))
	return retStr
}

// Convert infix to Poland notation
func (e *equationModel) toPolandNotation(strArr []string) (expression []string) {
	var stack []string
	precedence := map[string]int{
		"+":    1,
		"-":    1,
		"mod":  3, // % (остаток от деления)
		"*":    4,
		"/":    4,
		"^":    5,
		"cos":  6,
		"sin":  6,
		"tan":  6,
		"acos": 6,
		"asin": 6,
		"atan": 6,
		"sqrt": 6,
		"ln":   6,
		"log":  6,
		"(":    7,
		"{":    7,
		"[":    7,
	}
	open_brackets := "({["
	close_brackets := ")}]"
	operators := "+-^*/modacosasinatansqrtlnlog"
	for _, char := range strArr {
		switch {
		case char != "" && char[0] >= '0' && char[0] <= '9':
			expression = append(expression, string(char))
		case strings.Contains(operators, char):
			lenght := len(stack)
			if lenght == 0 || stack[lenght-1] == "(" {
				stack = append(stack, string(char))
			} else {
				for {
					if lenght > 0 && char != "" && precedence[char] <= precedence[stack[lenght-1]] && stack[lenght-1] != "(" {
						expression = append(expression, stack[lenght-1])
						stack[lenght-1] = ""
						stack = stack[:lenght-1]
						lenght = len(stack)
					} else {
						if char != "" && char != ")" {
							stack = append(stack, char)
						}
						break
					}
				}
			}
		case strings.Contains(open_brackets, char):
			stack = append(stack, string(char))
		case strings.Contains(close_brackets, char):
			lenght := len(stack)
			for {
				if strings.Contains(open_brackets, stack[lenght-1]) {
					break
				}
				expression = append(expression, stack[lenght-1])
				stack[lenght-1] = ""
				stack = stack[:lenght-1]
				lenght = len(stack)
			}
			if char == ")" {
				stack[lenght-1] = ""
				stack = stack[:lenght-1]
			}
		default:
			break
		}
	}
	if len(stack) > 0 {
		lenght := len(stack)
		for {
			if lenght == 0 {
				break
			}
			expression = append(expression, stack[lenght-1])
			stack[lenght-1] = ""
			stack = stack[:lenght-1]
			lenght = len(stack)
		}
	}
	return
}

// Calculate expression in poland notation
func (e *equationModel) calculate(expression []string) (float64, error) {
	operators := "+-^*/modacosasinatansqrtlnlog"
	unaryop := "acosasinatansqrtlnlog"
	var stack []float64
	var err error = nil

	for _, val := range expression {
		var temp float64
		if strings.Contains(operators, val) {
			if len(stack) < 2 {
				if strings.Contains(unaryop, val) {
					lenght := len(stack)
					n1 := stack[lenght-1]
					stack = stack[:lenght-1]
					switch {
					case val == "cos":
						temp = math.Cos(n1 * math.Pi / 180)
					case val == "sin":
						temp = math.Sin(n1 * math.Pi / 180)
					case val == "tan":
						temp = math.Tan(n1 * math.Pi / 180)
					case val == "acos":
						temp = math.Acos(n1) * 180 / math.Pi
					case val == "asin":
						temp = math.Asin(n1) * 180 / math.Pi
					case val == "atan":
						temp = math.Atan(n1) * 180 / math.Pi
					case val == "sqrt":
						temp = math.Sqrt(n1)
					case val == "ln":
						temp = math.Log(n1)
					case val == "log":
						temp = math.Log10(n1)
					}
					stack = append(stack, temp)
					continue
				} else {
					err = errors.New("Too few arguments")
					break
				}
			}
			lenght := len(stack)
			n1 := stack[lenght-1]
			stack = stack[:lenght-1]
			lenght = len(stack)
			n2 := stack[lenght-1]
			stack = stack[:lenght-1]
			switch {
			case val == "+":
				temp = n2 + n1
			case val == "-":
				temp = n2 - n1
			case val == "^":
				temp = math.Pow(n2, n1)
			case val == "*":
				temp = n2 * n1
			case val == "mod":
				temp = float64(int(n2) % int(n1))
			case val == "/":
				if n1 == 0 {
					err = errors.New("Error: division by zero")
					break
				}
				temp = n2 / n1
			case val == "cos":
				stack = append(stack, n2)
				temp = math.Cos(n1)
			case val == "sin":
				stack = append(stack, n2)
				temp = math.Sin(n1)
			case val == "tan":
				stack = append(stack, n2)
				temp = math.Tan(n1)
			case val == "acos":
				stack = append(stack, n2)
				temp = math.Acos(n1)
			case val == "asin":
				stack = append(stack, n2)
				temp = math.Asin(n1)
			case val == "atan":
				stack = append(stack, n2)
				temp = math.Atan(n1)
			case val == "sqrt":
				stack = append(stack, n2)
				temp = math.Sqrt(n1)
			case val == "ln":
				stack = append(stack, n2)
				temp = math.Log(n1)
			case val == "log":
				stack = append(stack, n2)
				temp = math.Log10(n1)
			}
			stack = append(stack, temp)
		} else {
			if num, err := strconv.ParseFloat(val, 64); err == nil {
				stack = append(stack, num)
			} else {
				err = fmt.Errorf("Error in strconv: %v", err)
			}
		}
	}
	return stack[0], err
}

// Common function of calculate with input prepared string (with unary and spaces)
func (e *equationModel) startCalculate(str string) (rez float64, err error) {
	return e.calculate(e.toPolandNotation(strings.Fields(str)))
}
