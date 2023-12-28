package model

import (
	"encoding/json"
	"os"
	"reflect"
	d "smartCalc/domains"
	"testing"
)

// create calcHistory object
func TestNewCalcHistory(t *testing.T) {
	ch := *NewCalcHistory(*configCalc)
	if reflect.TypeOf(ch).String() != "model.calcHistory" {
		t.Errorf("Creating object incorrect - expected: %v; actual: %v", "model.calcHistory", reflect.TypeOf(ch).String())
	} else if ch.fileName != configCalc.HistoryFile {
		t.Errorf("incorrect filename - expected: %v; actual: %v", d.Config.HistoryFile, ch.fileName)
	} else if reflect.TypeOf(ch.historyData).String() != "[]domains.HistoryItem" {
		t.Errorf("Creating object incorrect - expected: %v ; actual: %v", "[]domains.HistoryItem", reflect.TypeOf(ch.historyData).String())
	}

}

// history hendler
func TestHistoryHandler(t *testing.T) {
	in := createMInputStructEqual("x^2", "4", 4)
	calc := NewCalcModel(configCalc)
	lenght := len(calc.history.historyData)
	calc.GetCalcResult(in)
	if lenght+1 != len(calc.history.historyData) {
		t.Errorf("Adding History string incorrect - expected len=: %v ; actual: %v", lenght+1, len(calc.history.historyData))
	}
	in = createMInputStructEquation("x^2")

	lenght = len(calc.history.historyData)
	calc.GetCalcResult(in)
	if lenght+1 != len(calc.history.historyData) {
		t.Errorf("Adding History string incorrect - expected len=: %v ; actual: %v", lenght+1, len(calc.history.historyData))
	}
	in = createMInputStructGraph("x^2", "-4", "4", "-4", "4")

	lenght = len(calc.history.historyData)
	calc.GetCalcResult(in)
	if lenght+1 != len(calc.history.historyData) {
		t.Errorf("Adding History string incorrect - expected len=: %v ; actual: %v", lenght+1, len(calc.history.historyData))
	}
}

func TestCleanHistory(t *testing.T) {
	in := createMInputStructEqual("x^2", "4", 4)
	calc := NewCalcModel(configCalc)
	calc.GetCalcResult(in)
	lenght := len(calc.history.historyData)
	calc.history.CleanHistory()
	if lenght > 0 && len(calc.history.historyData) != 0 {
		t.Errorf("Clean History incorrect - expected len=: %v ; actual: %v", 0, len(calc.history.historyData))
	}
}

//----------------------------------------Creating historyItem START

// creating history string
func TestCreateHistoryItem(t *testing.T) {
	calc := NewCalcModel(configCalc)

	in := createMInputStructEquation("x^2")
	calc.GetCalcResult(in)
	th := calc.history.historyData[len(calc.history.historyData)-1]
	if th.Mode != "calc" {
		t.Errorf("CreateHistoryItem incorrect - expected len=: %v ; actual: %v", "calc", th.Mode)
	}

	in = createMInputStructEqual("x^2", "4", 4)
	calc.GetCalcResult(in)
	th = calc.history.historyData[len(calc.history.historyData)-1]
	if th.Mode != "equal" {
		t.Errorf("CreateHistoryItem incorrect - expected len=: %v ; actual: %v", "equal", th.Mode)
	}

	in = createMInputStructGraph("x^2", "-4", "4", "-4", "4")
	calc.GetCalcResult(in)
	th = calc.history.historyData[len(calc.history.historyData)-1]
	if th.Mode != "graph" {
		t.Errorf("CreateHistoryItem incorrect - expected len=: %v ; actual: %v", "graph", th.Mode)
	}
}

// creating history string for Calc mode
func TestCreateHistoryCalc(t *testing.T) {
	in := createMInputStructEquation("43-5*6")
	calc := NewCalcModel(configCalc)
	out := calc.GetCalcResult(in)
	calc.history.createHistoryCalc(in, out)
	th := calc.history.tempHistoryItem
	if th.Mode != "calc" || th.Equation != "43-5*6" || th.Result != "13" {
		t.Errorf("Adding History string incorrect - in=: %v ; history: %v", in, th)
	}
	in = createMInputStructEquation("43-5(*6")
	calc.history.createHistoryCalc(in, calc.GetCalcResult(in))
	th = calc.history.tempHistoryItem
	if th.Mode != "calc" || th.Equation != "43-5(*6" || th.Result != "error" {
		t.Errorf("Adding History string incorrect - in=: %v ; history: %v", in, th)
	}
}

// creating history string for Equal mode
func TestCreateHistoryEqual(t *testing.T) {

	calc := NewCalcModel(configCalc)

	in := createMInputStructEqual("43-x+5*6", "5", 5)
	calc.history.createHistoryEqual(in, calc.GetCalcResult(in))
	th := calc.history.tempHistoryItem
	if th.Mode != "equal" || th.Equation != "43-x+5*6" || th.Result != "68" || th.Entrys != "{X=5} " {
		t.Errorf("Adding History string incorrect - in=: %v ; history: %v", in, th)
	}
	in = createMInputStructEqual("43-x+5(*6", "5", 5)
	calc.history.createHistoryEqual(in, calc.GetCalcResult(in))
	th = calc.history.tempHistoryItem
	if th.Mode != "equal" || th.Equation != "43-x+5(*6" || th.Result != "error" || th.Entrys != "{X=5} " {
		t.Errorf("Adding History string incorrect - in=: %v ; history: %v", in, th)
	}
}

// creating history string for Graph mode
func TestCreateHistoryGraph(t *testing.T) {

	in := createMInputStructGraph("43-x+5*6", "-5", "5", "-5", "5")
	calc := NewCalcModel(configCalc)
	out := calc.GetCalcResult(in)
	calc.history.createHistoryGraph(in, out)
	th := calc.history.tempHistoryItem
	if th.Mode != "graph" || th.Equation != "43-x+5*6" || th.Result != "Y {0.00 .. 78.00}" || th.Entrys != "X{-5 .. 5} Y{-5 .. 5} " {
		t.Errorf("Adding History string incorrect - in=: %v ; history: %v", in, th)
	}
	in = createMInputStructGraph("43-x+5(*6", "-5", "5", "-5", "5")
	calc.history.createHistoryGraph(in, calc.GetCalcResult(in))
	th = calc.history.tempHistoryItem
	if th.Mode != "graph" || th.Equation != "43-x+5(*6" || th.Result != "error" {
		t.Errorf("Adding History string incorrect 2 - in=: %v ; history: %v", in, th)
	}
}

//----------------------------------------Creating historyItem END
//----------------------------------------Handle history START

// adding history string to history base
func TestAddHistoryString(t *testing.T) {
	// fmt.Println(os.Getwd())
	in := createMInputStructEquation("43-5*6")
	calc := NewCalcModel(configCalc)
	out := calc.GetCalcResult(in)
	calc.history.createHistoryCalc(in, out)
	lenght := len(calc.history.historyData)
	calc.history.addHistoryString()
	if lenght+1 != len(calc.history.historyData) {
		t.Errorf("Adding History string incorrect - expected: %v ; actual: %v", lenght+1, len(calc.history.historyData))
	}
}

// creating Json data from
func TestCreateHistoryJson(t *testing.T) {

	in := createMInputStructEquation("43-5*6")
	calc := NewCalcModel(configCalc)
	calc.history.fileName = "history_test.json"
	calc.history.historyData = []d.HistoryItem{}
	calc.GetCalcResult(in)

	data := calc.history.createHistoryJson()
	calc.history.historyData = []d.HistoryItem{}
	len0 := len(calc.history.historyData)
	err := json.Unmarshal(data, &calc.history.historyData)
	os.Remove("history_test.json")
	if len(calc.history.historyData) != 1 || len0 != 0 || err != nil {
		t.Errorf("Creating History JSON incorrect - expected: %v ; actual: %v", 1, len(calc.history.historyData))
	}
}

// writing entire history base to file
func TestWriteHistoryJson(t *testing.T) {
	in := createMInputStructEquation("43-5*6")
	calc := NewCalcModel(configCalc)
	os.Remove("history_test.json")
	calc.history.fileName = "history_test.json"
	calc.history.historyData = []d.HistoryItem{}
	calc.GetCalcResult(in)
	os.Remove("history_test.json")

	//Check Not write
	os.Create(TestPath + "/temphistory")
	os.Chmod(TestPath+"/temphistory", 0444)
	tempFileName := calc.history.fileName
	calc.history.fileName = TestPath + "/temphistory"
	if calc.history.writeHistoryJson([]byte("Check write")) == nil {
		t.Errorf("Writing History JSON incorrect - expected: error ; actual: nill")
	}
	os.Chmod(TestPath+"/temphistory", 0644)
	os.Remove(TestPath + "/temphistory")
	calc.history.fileName = tempFileName

	// Check write
	calc.history.writeHistoryJson(calc.history.createHistoryJson())
	calc.history.historyData = []d.HistoryItem{}
	calc.history.historyData = readHistoryJson("history_test.json")
	os.Remove("history_test.json")
	if len(calc.history.historyData) != 1 {
		t.Errorf("Writing History JSON incorrect - expected: %v ; actual: %v", 1, len(calc.history.historyData))
	}
}

// reading entire history from file
func TestReadHistoryJson(t *testing.T) {
	in := createMInputStructEquation("43-5*6")
	calc := NewCalcModel(configCalc)
	os.Remove("history_test.json")
	calc.history.fileName = "history_test.json"
	calc.history.historyData = []d.HistoryItem{}
	calc.GetCalcResult(in)
	calc.history.historyData = []d.HistoryItem{}
	calc.history.historyData = readHistoryJson("history_test.json")
	os.Remove("history_test.json")
	if len(calc.history.historyData) != 1 {
		t.Errorf("Reading History JSON incorrect - expected: %v ; actual: %v", 1, len(calc.history.historyData))
	}

	os.WriteFile(TestPath+"/temphistory", []byte("Check!"), 0111)
	if len(readHistoryJson(TestPath+"/temphistory")) > 0 {
		t.Errorf("Reading History JSON incorrect - expected: empty ; actual: no empty")
	}
	os.Chmod(TestPath+"/temphistory", 0644)
	if len(readHistoryJson(TestPath+"/temphistory")) > 0 {
		t.Errorf("Unmarshall JSON incorrect - expected: empty ; actual: no empty")
	}
	os.Remove(TestPath + "/temphistory")
}

//----------------------------------------Handle history END
