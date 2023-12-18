package model

import (
	d "smartCalc/domains"
	tools "smartCalc/tools"

	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"reflect"
	"testing"
)

var (
	configCalc = d.InitConfig("../conf/clevercalcLinuxTest.cfg")
	calcmodel  = NewCalcModel(configCalc)

	TestPath = "../test/"

	testCasesNewGraph = []struct {
		enter  []string
		expect []string
		err    bool
	}{
		{[]string{"34-23x/2", "5", "6", "7", "8"}, []string{"34-23x/2", "5", "6", "7", "8"}, false},
		{[]string{"34-23x/2", "0", "0", "7", "8"}, []string{"34-23x/2", "0", "0", "7", "8"}, true},
		{[]string{"34-23x/2", "0", "r1", "7", "8"}, []string{"34-23x/2", "0", "r1", "7", "8"}, true},
		{[]string{"34-23x/2", "0", "1000009", "-1000001", "8"}, []string{"34-23x/2", "0", "1000009", "-1000001", "8"}, true},
		{[]string{"34x^2-23x/2", "5", "600", "7", "8"}, []string{"34x^2-23x/2", "5", "600", "7", "8"}, false},
	}

	testCasesGraphGetResult = []struct {
		enter  []string
		expect []string
		err    bool
		outErr bool
	}{
		// {[]string{"", "5", "6", "7", "8"}, []string{"", "5", "6", "7", "8"}, false, true},
		{[]string{"34-23x(/2", "5", "6", "7", "8"}, []string{"34-23x(/2", "5", "6", "7", "8"}, false, true},
		{[]string{"34-23x/2", "5", "6", "7", "8"}, []string{"34-23x/2", "5", "6", "7", "8"}, false, false},
		{[]string{"34-23x/2", "0", "0", "7", "8"}, []string{"34-23x/2", "0", "0", "7", "8"}, true, true},
		{[]string{"34xmod3-23x/2", "5", "600", "7", "8"}, []string{"34xmod3-23x/2", "5", "600", "7", "8"}, false, false},
	}

	testCasesEntrysGraphCheck = []struct {
		enter []string
		err   bool
	}{
		{[]string{"34-56x", "5", "6", "7", "8"}, false},
		{[]string{"", "5", "6", "7", "8"}, true},
		{[]string{"34-56x", "5", "5", "7", "8"}, true},
		{[]string{"34-56x", "5", "6", "7", "7"}, true},
		{[]string{"34-56x", "-1000001", "6", "7", "8"}, true},
		{[]string{"34-56x", "5", "1000010", "7", "8"}, true},
		{[]string{"34-56x", "5", "6", "-1000002", "8"}, true},
		{[]string{"34-23x/2", "0", "1", "7", "1004000"}, true},
		{[]string{"34-23x/2", "0", "r1", "7", "8"}, true},
		{[]string{"34-23x/2", "0fg", "1", "7", "8"}, true},
		{[]string{"34-23x/2", "0", "1", "v b", "8"}, true},
		{[]string{"34-23x/2", "0", "1", "7", "*"}, true},
		{[]string{"34-23x/2", "0", "1000009", "-1000001", "8"}, true},
		{[]string{"34x^2-23x/2", "5", "600", "7", "8"}, false},
	}

	testCasesGraphPrepareString = []struct {
		enter []string
		err   bool
	}{
		{[]string{"", "5", "6", "7", "8"}, true},
		{[]string{"34-((56)x", "5", "6", "7", "8"}, true},
		{[]string{"(34)-56x", "5", "6", "7", "8"}, false},
		{[]string{"34-(56x*{34))", "5", "5", "7", "8"}, true},
		{[]string{"34-[56]x", "5", "6", "7", "7"}, false},
	}

	testCasesCalculateData = []struct {
		enter  []string
		expect []float64
	}{
		{[]string{"2x", "-300", "300", "-300", "300"}, []float64{-600.00, 600.00, 601}},
		{[]string{"cos(x)", "-300", "300", "-3", "3"}, []float64{-1.00, 1.00, 601}},
		{[]string{"cos(x)", "-100", "200", "-3", "3"}, []float64{-1.00, 1.00, 601}},
	}

	testCasesGraphDraw = []struct {
		enter  []string
		expect string
		b      bool
	}{
		{[]string{"x", "-300", "300", "-300", "300"}, TestPath + "/graphDraw_1.png", true},
		{[]string{"x^(2)/10", "-300", "300", "-300", "300"}, TestPath + "/graphDraw_2.png", true},
	}

	testCasesDrawEqualText = []struct {
		enter  []string
		expect string
		b      bool
	}{
		{[]string{"x", "-300", "300", "-300", "300"}, TestPath + "/drawEqualText.png", true},
	}

	testCasesDrawLogo = []struct {
		enter  []string
		expect string
		b      bool
	}{
		{[]string{"x", "-300", "300", "-300", "300"}, TestPath + "/drawLogo.png", true},
	}

	testCasesCreateEqualText = []struct {
		enter  string
		float  []float64
		expect string
	}{
		{"2x", []float64{-300, 400, -600, 700}, "2x  X{-300.00 .. 400.00} Y{-600.00 .. 700.00}"},
		{"cos(x)", []float64{-23, 400, -600, 10}, "cos(x)  X{-23.00 .. 400.00} Y{-600.00 .. 10.00}"},
		{"2x-25", []float64{-10, 400, -45, 700}, "2x-25  X{-10.00 .. 400.00} Y{-45.00 .. 700.00}"},
	}

	testCasesFindGridValue = []struct {
		enter  []float64
		expect []float64
	}{
		{[]float64{100, -300, 400}, []float64{0, -300, -200, -100, 0, 100, 200, 300}},
		{[]float64{200, -300, 1400}, []float64{0, -400, -200, 0, 200, 400, 600, 800, 1000, 1200}},
		{[]float64{200, 0, 1400}, []float64{0, 1400, 1200, 1000, 800, 600, 400, 200}},
		{[]float64{10, 1, 80}, []float64{80, 70, 60, 50, 40, 30, 20, 10}},
	}

	testCasesPositionMinMax = []struct {
		enter []float64
	}{
		{[]float64{100, -300}},
		{[]float64{-100, -300}},
		{[]float64{100, 400}},
		{[]float64{0, -300}},
		{[]float64{-300, -300}},
	}
	testCasesPrepareGridValue = []struct {
		enter  []float64
		expect string
	}{
		{[]float64{10, 100}, "100"},
		{[]float64{500, 1500}, "1500"},
		{[]float64{2, 2}, "2"},
		{[]float64{0.1, 0.2}, "0.20"},
		{[]float64{0.01, 0.03}, "0.03"},
	}

	testCasesFindMasshtab = []struct {
		enter  []float64
		expect float64
	}{
		{[]float64{10, 100}, 10},
		{[]float64{500, 1500}, 200},
		{[]float64{2, 3}, 0.2},
		{[]float64{0.1, 0.2}, 0.02},
		{[]float64{0.01, 0.03}, 0.002},
	}

	testCasesCheckSideModeArrowY = []struct {
		enter  []float64
		expect bool
	}{
		{[]float64{10, 100}, true},
		{[]float64{500, 1500}, true},
		{[]float64{2, 3}, true},
		{[]float64{0.1, 0.2}, true},
		{[]float64{-0.01, 0.03}, false},
		{[]float64{-10, 10}, false},
		{[]float64{-3, 1}, false},
		{[]float64{-30, 2}, true},
	}

	testCasesCheckSideModeArrow = []struct {
		enter  []float64
		size   int
		expect bool
	}{
		{[]float64{10, 100}, 600, true},
		{[]float64{500, 1500}, 600, true},
		{[]float64{2, 3}, 600, true},
		{[]float64{0.1, 0.2}, 600, true},
		{[]float64{-0.01, 0.03}, 600, false},
		{[]float64{-10, 10}, 600, false},
		{[]float64{-3, 1}, 600, false},
		{[]float64{-30, 2}, 600, true},
	}

	testCasesCreateArrayValue = []struct {
		enter  []float64
		expect []float64
	}{
		{[]float64{10, 100, 10}, []float64{10, 19, 28, 37, 46, 55, 64, 73, 82, 91}},
		{[]float64{-45, 100, 10}, []float64{-45, -30.5, -16, -1.5, 13, 27.5, 42, 56.5, 71, 85.5}},
	}

	testCasesGraphImageBuild = []struct {
		enter  []string
		expect string
		b      bool
	}{
		{[]string{"x^(2)", "-30", "60", "-30", "60"}, TestPath + "/x^2_-30_60_-30_60.png", true},
		{[]string{"x^(2)", "-30", "60", "-30", "60"}, TestPath + "/Icon.png", false},
		{[]string{"x^(2)", "-30", "60", "-30", "60"}, TestPath + "/x^2_-40_60_-30_60.png", false},
	}

	testCasesGraphGridFindValue = []struct {
		enter  []float64
		mode   string
		expect int
	}{
		{[]float64{10, 100, 45}, "X", 233},
		{[]float64{10, 100, 15}, "X", 33},
		{[]float64{-10, 900, 45}, "X", 36},
	}

	testCasesGraphGridDraw = []struct {
		enter  []string
		expect string
		b      bool
	}{
		{[]string{"x^(2)/500", "-500", "300", "-30", "3000"}, TestPath + "/x^2_500_-500_300_-30_3000 .png", true},
	}

	testCasesDrawHLine = []struct {
		enter  []int
		expect string
		b      bool
	}{
		{[]int{3, 100, 20, 20, 1, 10, 40, 40, 10, 400, 80, 40}, TestPath + "/drawHLine.png", true},
	}

	testCasesDrawVLine = []struct {
		enter  []int
		expect string
		b      bool
	}{
		{[]int{3, 100, 20, 20, 1, 10, 40, 40, 10, 400, 80, 120}, TestPath + "/drawVLine.png", true},
	}

	testCasesMasshtabDraw = []struct {
		enter  []string
		expect string
		b      bool
	}{
		{[]string{"x", "-2", "2", "-2", "2"}, TestPath + "/masshtab_-2_2_-2_2.png", true},
		{[]string{"x", "-2", "200", "-2", "2"}, TestPath + "/masshtab_-2_200_-2_2.png", true},
	}

	testCasesFillBackground = []struct {
		enter  []string
		expect string
		b      bool
	}{
		{[]string{"x", "-2", "2", "-2", "2"}, TestPath + "/background.png", true},
	}

	testCasesArrowH = []struct {
		enter  []int
		expect string
		b      bool
	}{
		{[]int{30, 100, 200, 350, 250, 500}, TestPath + "/arrowH.png", true},
	}

	testCasesArrowV = []struct {
		enter  []int
		expect string
		b      bool
	}{
		{[]int{30, 100, 200, 350, 250, 500}, TestPath + "/arrowV.png", true},
	}

	testCasesDrawGridText = []struct {
		enter  []int
		expect string
		b      bool
	}{
		{[]int{30, 100, 150, 350, 400, 450}, TestPath + "/drawGridText.png", true},
	}

	testCasesPrintGridValueX = []struct {
		enter  []float64
		expect string
		b      bool
	}{
		{[]float64{0.2, 0.4}, TestPath + "/printGridValueX.png", false},
		{[]float64{0.005, 0.002}, TestPath + "/printGridValueX.png", false},
		{[]float64{1, 1}, TestPath + "/printGridValueX.png", false},
		{[]float64{0.2, 0.4}, TestPath + "/printGridValueX.png", true},
		{[]float64{0.005, 0.002}, TestPath + "/printGridValueX.png", true},
		{[]float64{1, 1}, TestPath + "/printGridValueX.png", true},
	}

	testCasesPrintGridValueY = []struct {
		enter  []float64
		expect string
		b      bool
	}{
		{[]float64{0.2, 0.4}, TestPath + "/printGridValueY.png", false},
		{[]float64{0.005, 0.002}, TestPath + "/printGridValueY.png", false},
		{[]float64{1, 1}, TestPath + "/printGridValueY.png", false},
		{[]float64{0.2, 0.4}, TestPath + "/printGridValueY.png", true},
		{[]float64{0.005, 0.002}, TestPath + "/printGridValueY.png", true},
		{[]float64{1, 1}, TestPath + "/printGridValueY.png", true},
	}
)

// Compare []float64 false if equal
func compareArrFloat64(a []float64, b []float64) bool {
	if len(a) != len(b) {
		return true
	}
	for i, val := range a {
		if val != b[i] {
			return true
		}
	}
	return false
}

// Compare image true if equal
func CompareImage(img1, img2 image.Image) bool {

	// compare bounds
	if img1.Bounds() != img2.Bounds() {
		return false
	}

	// compare pixels
	size := img1.Bounds().Size()
	for y := 0; y < size.Y; y++ {
		for x := 0; x < size.X; x++ {
			if img1.At(x, y) != img2.At(x, y) {
				return false
			}
		}
	}
	return true
}

func fillBackground(img draw.Image, c color.Color) {
	for x := 0; x < int(configCalc.XWindowGraph); x++ {
		for y := 0; y < int(configCalc.YWindowGraph); y++ {
			img.Set(x, y, c)
		}
	}
}

//---------- New object and interface implementing

func TestNewGraph(t *testing.T) {
	for _, tC := range testCasesNewGraph {

		e := createMInputStructGraph(tC.enter[0], tC.enter[1], tC.enter[2], tC.enter[3], tC.enter[4])
		er := *calcmodel.NewGraph(e)

		if reflect.TypeOf(er).String() != "model.graphModel" {
			t.Errorf("Creating object incorrect - expected: %v; actual: %v", "model.graphModel", reflect.TypeOf(er).String())
		} else if reflect.TypeOf(er.equal).String() != "model.equalModel" {
			t.Errorf("Creating object incorrect - expected: %v; actual: %v", "model.equalModel", reflect.TypeOf(er.equal).String())
		} else if (er.err != nil) != tC.err {
			t.Errorf("Result was incorrect, expected: %v, actual: %v\n %v %v", er.err != nil, tC.err, tC.enter, tC.expect)

		} else if er.equalValue != tC.expect[0] ||
			er.xFromStr != tC.expect[1] || er.xToStr != tC.expect[2] ||
			er.yFromStr != tC.expect[3] || er.yToStr != tC.expect[4] {
			t.Errorf("Copyed data incorrect, expected: %v, actual: %v\n", tC.expect, er)
		}
	}
}

// Implementing request interface for graphModel
func TestGetResultGraph(t *testing.T) {
	for _, tC := range testCasesGraphGetResult {
		e := createMInputStructGraph(tC.enter[0], tC.enter[1], tC.enter[2], tC.enter[3], tC.enter[4])
		c := *calcmodel
		gm := c.NewGraph(e)
		er := gm.GetResult()

		if reflect.TypeOf(er).String() != "model.ModelsOutput" {
			t.Errorf("Creating object incorrect - expected: %v; actual: %v", "model.ModelsOutput", reflect.TypeOf(er).String())
		} else if er.Err != tC.outErr || er.ModelGraphResult.Err != tC.outErr {
			t.Errorf("Result was incorrect, expected: %v, actual: %v\n %v %v", tC.outErr, tC.outErr, er.Err, er.ModelGraphResult.Err)

		} else if er.ModelGraphResult.ResultStr == "" && !tC.outErr {
			t.Errorf("ResultStr incorrect: empty sring %v", er)
		}
	}
}

//--------------Calculate and preparing data START

// Check and copy input data for graph
func TestEntrysGraphCheck(t *testing.T) {
	for _, tC := range testCasesEntrysGraphCheck {
		var er graphModel
		e := createMInputStructGraph(tC.enter[0], tC.enter[1], tC.enter[2], tC.enter[3], tC.enter[4])
		if er.entrysGraphCheck(e) != tC.err {
			t.Errorf("Copyed data incorrect, expected: %v, actual: %v\n, %v", tC.err, !(tC.err), tC)
		}
	}
}

// graph prepared string - check and stapples(without replace X)
func TestGraphPrepareString(t *testing.T) {
	for _, tC := range testCasesGraphPrepareString {
		e := createMInputStructGraph(tC.enter[0], tC.enter[1], tC.enter[2], tC.enter[3], tC.enter[4])
		er := *calcmodel.NewGraph(e)
		if er.graphPrepareString() != tC.err {
			t.Errorf("Graph prepare string is incorrect, expected: %v, actual: %v\n, %v", tC.err, !(tC.err), tC)
		}
	}
}

// calculate pixels data for graph
func TestCalculateData(t *testing.T) {
	for _, tC := range testCasesCalculateData {
		e := createMInputStructGraph(tC.enter[0], tC.enter[1], tC.enter[2], tC.enter[3], tC.enter[4])
		er := *calcmodel.NewGraph(e)
		er.graphPrepareString()
		er.calculateData()
		var deltaPixel float64 = math.Abs(er.xFrom-er.xTo) / float64(int(er.config.XWindowGraph))
		if er.gRM.yGraphMin != tC.expect[0] || er.gRM.yGraphMax != tC.expect[1] ||
			len(er.gRM.pixelData) != int(tC.expect[2]) || (er.gRM.pixelData[2].x-er.gRM.pixelData[1].x) != deltaPixel {
			t.Errorf("Graph calculate is incorrect, expected: %v, %v\n actual: %v, %v , %v, %v", tC.expect, deltaPixel,
				er.gRM.yGraphMin, er.gRM.yGraphMax, len(er.gRM.pixelData), er.gRM.pixelData[2].x-er.gRM.pixelData[1].x)
		}
	}
}

func TestGraphImageBuild(t *testing.T) {
	for _, tC := range testCasesGraphImageBuild {
		e := createMInputStructGraph(tC.enter[0], tC.enter[1], tC.enter[2], tC.enter[3], tC.enter[4])
		er := *calcmodel.NewGraph(e)
		er.graphPrepareString()
		er.calculateData()
		er.graphImageBuild()
		if im, err := tools.LoadImage(tC.expect); err == nil {
			if CompareImage(im, er.gRM.graphImage) != tC.b {
				t.Errorf("GraphImageBuild's print is incorrect: %v", tC)
			}
		} else {
			fmt.Println("Can't load test image")
		}
	}
}

// Creating text with equal data for image
func TestCreateEqualText(t *testing.T) {
	for _, tC := range testCasesCreateEqualText {
		var er graphModel
		er.equalValue = tC.enter
		er.xFrom = tC.float[0]
		er.xTo = tC.float[1]
		er.gRM.yGraphMin = tC.float[2]
		er.gRM.yGraphMax = tC.float[3]
		actual := er.createEqualText()
		if actual != tC.expect {
			t.Errorf("CreateEqualText is incorrect, expected: %v, actual: %v, ", tC.expect, actual)
		}
	}
}

func TestCreateArrayValue(t *testing.T) {
	for _, tC := range testCasesCreateArrayValue {
		var er graphModel
		arr := er.createArrayValue(tC.enter[0], tC.enter[1], tC.enter[2])
		if compareArrFloat64(arr, tC.expect) {
			t.Errorf("CreateArrayValue is incorrect, expected: %v, actual: %v", tC.expect, arr)
		}
	}
}

// Check sidemode
func TestCheckSideModeArrow(t *testing.T) {
	for _, tC := range testCasesCheckSideModeArrow {
		var er graphModel
		if er.checkSideModeArrow(tC.enter[0], tC.enter[1], tC.size) != tC.expect {
			t.Errorf("CheckSideModeArrow is incorrect, expected: %v, actual: %v, %v", tC.expect, !(tC.expect), tC)
		}
	}
}

// Place min to min, max to max
func TestPositionMinMax(t *testing.T) {
	for _, tC := range testCasesPositionMinMax {
		first, second := tC.enter[0], tC.enter[1]
		positionMinMax(&first, &second)

		if second < first {
			t.Errorf("PositionMinMax is incorrect, expected the order: %v,%v actual: %v, %v", first, second, second, first)
		}
	}
}

func TestFindGridValue(t *testing.T) {
	for _, tC := range testCasesFindGridValue {
		var er graphModel
		actual := er.findGridValue(tC.enter[0], tC.enter[1], tC.enter[2])
		if compareArrFloat64(tC.expect, actual) {
			t.Errorf("Find Grid Value is incorrect, expected len: %v, actual len: %v", len(tC.expect), len(actual))
		}
	}
}

func TestFindMasshtab(t *testing.T) {
	var er graphModel
	for _, tC := range testCasesFindMasshtab {
		actual := er.findMasshtab(tC.enter[0], tC.enter[1])
		if actual != tC.expect {
			t.Errorf("FindMasshtab is incorrect, expected: %v, actual: %v", tC.expect, actual)
		}
	}
}

// Print Logo at the picture
func TestDrawLogo(t *testing.T) {
	for _, tC := range testCasesDrawLogo {
		e := createMInputStructGraph(tC.enter[0], tC.enter[1], tC.enter[2], tC.enter[3], tC.enter[4])
		er := *calcmodel.NewGraph(e)
		er.gRM.graphImage = image.NewRGBA(image.Rect(0, 0, int(int(er.config.XWindowGraph)), int(er.config.YWindowGraph)))
		er.fillBackground(er.gRM.graphImage, color.White)
		er.drawLogo(er.gRM.graphImage, 21, "CleverCalc")
		if im, err := tools.LoadImage(tC.expect); err == nil {
			if CompareImage(im, er.gRM.graphImage) != tC.b {
				t.Errorf("DrawLogo's print is incorrect: %v", tC)
			}
		} else {
			fmt.Println("Can't load test image")
		}
	}
}

// Print equal data for the graph
func TestDrawEqualText(t *testing.T) {
	for _, tC := range testCasesDrawEqualText {
		e := createMInputStructGraph(tC.enter[0], tC.enter[1], tC.enter[2], tC.enter[3], tC.enter[4])
		er := *calcmodel.NewGraph(e)
		er.graphPrepareString()
		er.calculateData()
		er.gRM.graphImage = image.NewRGBA(image.Rect(0, 0, int(er.config.XWindowGraph), int(er.config.YWindowGraph)))
		er.fillBackground(er.gRM.graphImage, color.White)
		er.drawEqualText(er.gRM.graphImage, 20, 30, er.createEqualText())
		if im, err := tools.LoadImage(tC.expect); err == nil {
			if CompareImage(im, er.gRM.graphImage) != tC.b {
				t.Errorf("DrawEqualText's print is incorrect: %v", tC)
			}
		} else {
			fmt.Println("Can't load test image")
		}
	}
}

func TestDrawGridText(t *testing.T) {
	for _, tC := range testCasesDrawGridText {
		var er graphModel
		er.config = configCalc
		er.gRM.graphImage = image.NewRGBA(image.Rect(0, 0, int(er.config.XWindowGraph), int(er.config.YWindowGraph)))
		er.fillBackground(er.gRM.graphImage, color.White)
		for i := 0; i < 5; i += 2 {
			er.drawGridText(er.gRM.graphImage, tC.enter[i], tC.enter[i+1], "!0123456789")
		}

		if im, err := tools.LoadImage(tC.expect); err == nil {
			if CompareImage(im, er.gRM.graphImage) != tC.b {
				t.Errorf("DrawGridText's print is incorrect: %v", tC)
			}
		} else {
			fmt.Println("Can't load test image")
		}
	}
}

// find X for printing gridVLine
func TestGraphGridFindValue(t *testing.T) {
	var er graphModel
	er.config = configCalc
	for _, tC := range testCasesGraphGridFindValue {
		er.xFrom, er.xTo = tC.enter[0], tC.enter[1]
		// int(er.config.XWindowGraph) and int(er.config.YWindowGraph) from global (=600)
		actual := er.graphGridFindValue(tC.enter[2], tC.mode)
		if actual != tC.expect {
			t.Errorf("GraphGridFindValue is incorrect, expected: %v, actual: %v", tC.expect, actual)
		}
	}
}

func TestPrepareGridValue(t *testing.T) {
	var er graphModel
	for _, tC := range testCasesPrepareGridValue {
		actual := er.prepareGridValue(tC.enter[0], tC.enter[1])
		if actual != tC.expect {
			t.Errorf("PrepareGridValue is incorrect, expected: %v, actual: %v", tC.expect, actual)
		}
	}

}

func TestPrintGridValueY(t *testing.T) {
	var er graphModel
	er.config = configCalc
	var cmpFile string
	er.gRM.graphImage = image.NewRGBA(image.Rect(0, 0, int(er.config.XWindowGraph), int(er.config.YWindowGraph)))
	er.fillBackground(er.gRM.graphImage, color.White)

	for _, tC := range testCasesPrintGridValueY {
		cmpFile = tC.expect
		er.printGridValueY(er.gRM.graphImage, tC.enter[0], tC.enter[1], tC.b)
	}

	if im, err := tools.LoadImage(cmpFile); err == nil {
		if CompareImage(im, er.gRM.graphImage) {
			t.Errorf("PrintGridValueY's print is incorrect:")
		}
	} else {
		fmt.Println("Can't load test image")
	}
}

func TestPrintGridValueX(t *testing.T) {
	var er graphModel
	er.config = configCalc
	var cmpFile string
	er.gRM.graphImage = image.NewRGBA(image.Rect(0, 0, int(er.config.XWindowGraph), int(er.config.YWindowGraph)))
	er.fillBackground(er.gRM.graphImage, color.White)

	for _, tC := range testCasesPrintGridValueX {
		cmpFile = tC.expect
		er.printGridValueX(er.gRM.graphImage, tC.enter[0], tC.enter[1], tC.b)
	}

	if im, err := tools.LoadImage(cmpFile); err == nil {
		if CompareImage(im, er.gRM.graphImage) {
			t.Errorf("PrintGridValueX's print is incorrect:")
		}
	} else {
		fmt.Println("Can't load test image")
	}
}

// Draw grid for graph image
func TestGraphGridDraw(t *testing.T) {
	for _, tC := range testCasesMasshtabDraw {
		e := createMInputStructGraph(tC.enter[0], tC.enter[1], tC.enter[2], tC.enter[3], tC.enter[4])
		er := *calcmodel.NewGraph(e)
		er.gRM.graphImage = image.NewRGBA(image.Rect(0, 0, int(er.config.XWindowGraph), int(er.config.YWindowGraph)))
		er.fillBackground(er.gRM.graphImage, color.White)

		er.masshtabDraw(er.gRM.graphImage, er.findMasshtab(er.xFrom, er.xTo),
			er.findMasshtab(er.yFrom, er.yTo), 500, 550)

		if im, err := tools.LoadImage(tC.expect); err == nil {
			if CompareImage(im, er.gRM.graphImage) != tC.b {
				t.Errorf("GraphGridDraw's print is incorrect: %v", tC)
			}
		} else {
			fmt.Println("Can't load test image")
		}
	}
}

// Draw graph line at the sCoordinates Array
func TestGraphDraw(t *testing.T) {
	for _, tC := range testCasesGraphDraw {
		e := createMInputStructGraph(tC.enter[0], tC.enter[1], tC.enter[2], tC.enter[3], tC.enter[4])
		er := *calcmodel.NewGraph(e)
		er.graphPrepareString()
		er.calculateData()
		er.gRM.graphImage = image.NewRGBA(image.Rect(0, 0, int(er.config.XWindowGraph), int(er.config.YWindowGraph)))
		er.fillBackground(er.gRM.graphImage, color.White)
		er.graphDraw()
		if im, err := tools.LoadImage(tC.expect); err == nil {
			if CompareImage(im, er.gRM.graphImage) != tC.b {
				t.Errorf("GraphDraw's print is incorrect: %v", tC)
			}
		} else {
			fmt.Println("Can't load test image")
		}
	}
}

func TestFillBackground(t *testing.T) {
	for _, tC := range testCasesFillBackground {
		e := createMInputStructGraph(tC.enter[0], tC.enter[1], tC.enter[2], tC.enter[3], tC.enter[4])
		er := *calcmodel.NewGraph(e)
		er.gRM.graphImage = image.NewRGBA(image.Rect(0, 0, int(er.config.XWindowGraph), int(er.config.YWindowGraph)))
		er.fillBackground(er.gRM.graphImage, color.White)

		if im, err := tools.LoadImage(tC.expect); err == nil {
			if CompareImage(im, er.gRM.graphImage) != tC.b {
				t.Errorf("FillBackground's print is incorrect: %v", tC)
			}
		} else {
			fmt.Println("Can't load test image")
		}
	}
}

// Drowing masshtab
func TestMasshtabDraw(t *testing.T) {
	for _, tC := range testCasesMasshtabDraw {
		e := createMInputStructGraph(tC.enter[0], tC.enter[1], tC.enter[2], tC.enter[3], tC.enter[4])
		er := *calcmodel.NewGraph(e)
		er.gRM.graphImage = image.NewRGBA(image.Rect(0, 0, int(er.config.XWindowGraph), int(er.config.YWindowGraph)))
		er.fillBackground(er.gRM.graphImage, color.White)
		er.masshtabDraw(er.gRM.graphImage, er.findMasshtab(er.xFrom, er.xTo),
			er.findMasshtab(er.yFrom, er.yTo), 500, 550)
		if im, err := tools.LoadImage(tC.expect); err == nil {
			if CompareImage(im, er.gRM.graphImage) != tC.b {
				t.Errorf("MasshtabDraw's print is incorrect: %v", tC)
			}
		} else {
			fmt.Println("Can't load test image")
		}
	}
}

// Draw vertical Up arrow
func TestArrowV(t *testing.T) {
	for _, tC := range testCasesArrowV {
		var er graphModel
		er.config = configCalc
		er.gRM.graphImage = image.NewRGBA(image.Rect(0, 0, int(er.config.XWindowGraph), int(er.config.YWindowGraph)))
		er.fillBackground(er.gRM.graphImage, color.White)
		for i := 0; i < 5; i += 2 {
			er.arrowV(er.gRM.graphImage, tC.enter[i], tC.enter[i+1], color.Black)
		}

		if im, err := tools.LoadImage(tC.expect); err == nil {
			if CompareImage(im, er.gRM.graphImage) != tC.b {
				t.Errorf("ArrowV's print is incorrect: %v", tC)
			}
		} else {
			fmt.Println("Can't load test image")
		}
	}
}

// Draw horisontal right arrow
func TestArrowH(t *testing.T) {
	for _, tC := range testCasesArrowH {
		var er graphModel
		er.config = configCalc
		er.gRM.graphImage = image.NewRGBA(image.Rect(0, 0, int(er.config.XWindowGraph), int(er.config.YWindowGraph)))
		er.fillBackground(er.gRM.graphImage, color.White)
		for i := 0; i < 5; i += 2 {
			er.arrowH(er.gRM.graphImage, tC.enter[i], tC.enter[i+1], color.Black)
		}

		if im, err := tools.LoadImage(tC.expect); err == nil {
			if CompareImage(im, er.gRM.graphImage) != tC.b {
				t.Errorf("ArrowH's print is incorrect: %v", tC)
			}
		} else {
			fmt.Println("Can't load test image")
		}
	}
}

// Draw vertical line
func TestDrawVLine(t *testing.T) {
	for _, tC := range testCasesDrawVLine {
		var er graphModel
		er.config = configCalc
		er.gRM.graphImage = image.NewRGBA(image.Rect(0, 0, int(er.config.XWindowGraph), int(er.config.YWindowGraph)))
		er.fillBackground(er.gRM.graphImage, color.White)
		for i := 0; i < 10; i += 4 {
			er.drawVLine(er.gRM.graphImage, tC.enter[i], tC.enter[i+1], tC.enter[i+2], tC.enter[i+3], color.Black)
		}

		if im, err := tools.LoadImage(tC.expect); err == nil {
			if CompareImage(im, er.gRM.graphImage) != tC.b {
				t.Errorf("DrawVLine's print is incorrect: %v", tC)
			}
		} else {
			fmt.Println("Can't load test image")
		}
	}
}

// Draw horisontal line
func TestDrawHLine(t *testing.T) {
	for _, tC := range testCasesDrawHLine {
		var er graphModel
		er.config = configCalc

		er.gRM.graphImage = image.NewRGBA(image.Rect(0, 0, int(er.config.XWindowGraph), int(er.config.YWindowGraph)))
		er.fillBackground(er.gRM.graphImage, color.White)
		for i := 0; i < 10; i += 4 {
			er.drawHLine(er.gRM.graphImage, tC.enter[i], tC.enter[i+1], tC.enter[i+2], tC.enter[i+3], color.Black)
		}

		if im, err := tools.LoadImage(tC.expect); err == nil {
			if CompareImage(im, er.gRM.graphImage) != tC.b {
				t.Errorf("DrawHLine's print is incorrect: %v", tC)
			}
		} else {
			fmt.Println("Can't load test image")
		}
	}
}
