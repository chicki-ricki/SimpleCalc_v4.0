package model

import (
	"fmt"
	"image/color"
	"image/draw"
	"math"
	t "smartCalc/tools"
	"strconv"
)

// Draw grid for graph image
func (g *graphModel) graphGridDraw(img draw.Image) {

	// Set variable for print vertical grid lines
	sideModeX := g.checkSideModeArrow(g.yFrom, g.yTo, int(g.config.YWindowGraph))
	masshtabX := g.findMasshtab(g.xFrom, g.xTo)
	arr := g.findGridValue(masshtabX, g.xFrom, g.xTo)

	// Print vertical grid lines
	for _, val := range arr {
		if val == 0 {
			t.Clg.DeepDebug(fmt.Sprint("_graphGridDraw_ g.graphGridFindX(0)=", g.graphGridFindValue(0, "X")))
			g.arrowV(img, g.graphGridFindValue(0, "X"), 7, color.Gray{Y: uint8(0)})
			g.drawVLine(img, 2, int(g.config.YWindowGraph)-20, g.graphGridFindValue(val, "X"), 10, color.Gray{Y: uint8(0)})
		} else {
			g.drawVLine(img, 1, int(g.config.YWindowGraph)-20, g.graphGridFindValue(val, "X"), 10, color.Gray{Y: uint8(125)})
		}
		g.printGridValueX(img, masshtabX, val, sideModeX)
	}

	// Set variable for print horisontal grid lines
	masshtabY := g.findMasshtab(g.yFrom, g.yTo)
	sideModeY := g.checkSideModeArrow(g.xFrom, g.xTo, int(g.config.XWindowGraph))
	arr = g.findGridValue(masshtabY, g.yFrom, g.yTo)

	// Print vertical grid lines
	for _, val := range arr {
		if val == 0 {
			t.Clg.DeepDebug(fmt.Sprint("_graphGridDraw_ g.graphGridFindY(0)=", g.graphGridFindValue(0, "Y")))
			g.drawHLine(img, 2, int(g.config.XWindowGraph)-20, 10, g.graphGridFindValue(val, "Y"), color.Gray{Y: uint8(0)})
			g.arrowH(img, int(g.config.XWindowGraph)-7, g.graphGridFindValue(0, "Y"), color.Gray{Y: uint8(0)})
		} else {
			g.drawHLine(img, 1, int(g.config.XWindowGraph)-20, 10, g.graphGridFindValue(val, "Y"), color.Gray{Y: uint8(125)})
		}
		g.printGridValueY(img, masshtabY, val, sideModeY)
	}

	// Print masshtab
	g.masshtabDraw(img, masshtabX, masshtabY, int(g.config.XWindowGraph)-int(g.config.XWindowGraph)/6, int(g.config.YWindowGraph)-50)
}

// Drawing masshtab
func (g *graphModel) masshtabDraw(img draw.Image, masshtabX, masshtabY float64, x0, y0 int) {
	xLine := g.graphGridFindValue(g.xFrom+masshtabX, "X") - g.graphGridFindValue(g.xFrom, "X")
	t.Clg.DeepDebug(fmt.Sprint("_masshtabDraw_ xLine for masshtab:", xLine))
	g.drawHLine(img, 3, int(xLine), x0-int(xLine)/2, y0, color.Gray{Y: uint8(0)})
	g.drawEqualText(img, x0-20, y0+25, fmt.Sprint("X: ", masshtabX))

	yLine := int(math.Abs(float64(g.graphGridFindValue(g.yFrom+masshtabY, "Y") - g.graphGridFindValue(g.yFrom, "Y"))))
	t.Clg.DeepDebug(fmt.Sprint("_masshtabDraw_ yLine for masshtab:", yLine))
	g.drawVLine(img, 3, int(yLine), x0, y0-int(yLine), color.Gray{Y: uint8(0)})
	g.drawEqualText(img, x0+10, y0+50-int(yLine), fmt.Sprint("Y: ", masshtabY))
}

//---------------------------Calculate Grid

// Create Array of X or Y
func (g *graphModel) createArrayValue(min, max, size float64) (arr []float64) {
	positionMinMax(&min, &max)
	var deltaPixel float64 = (max - min) / (size)
	for i := min; i < max; i += deltaPixel {
		arr = append(arr, i)
	}

	return
}

// Check sidemode (when line with arrow close to border)
func (g *graphModel) checkSideModeArrow(min, max float64, size int) bool {
	var tempDelta float64 = 2000000
	var Finded int

	arr := g.createArrayValue(min, max, float64(size-1))

	for i, val := range arr {
		if math.Abs(0-val) < tempDelta {
			tempDelta = math.Abs(0 - val)
			Finded = i
		}
	}

	if Finded < 100 || Finded > 500 {
		return true
	}

	return false
}

// Place min to min, max to max
func positionMinMax(min, max *float64) {
	if *max < *min {
		temp := *min
		*min = *max
		*max = temp
	}
}

// finding X or Y Values for grid lines
func (g *graphModel) findGridValue(masshtab, from, to float64) (arr []float64) {

	positionMinMax(&from, &to)

	// add zero arrow
	if to == 0 || from == 0 || (from < 0 && to > 0) {
		arr = append(arr, 0)
	}

	// find value from - to +
	if (from < 0 && to > 0) || to <= 0 {
		for x := masshtab * math.Round(from/masshtab); x < to; x += masshtab {
			arr = append(arr, x)
		}

		// find value from + to -
	} else if from >= 0 {
		for x := masshtab * math.Round(to/masshtab); x > from; x -= masshtab {
			arr = append(arr, x)
		}
	}

	// fmt.Println("arr", arr, "masshtab:", masshtab)
	return
}

// finding masshtab for grid
func (g *graphModel) findMasshtab(min, max float64) float64 {
	volume := math.Abs(max - min)
	for x := 0.01; x <= 10000000; x *= 10 {
		if volume/x < 10 && volume/x >= 5 {
			return x
		} else if volume/x < 5 && volume/x >= 2 {
			return x / 2
		} else if volume/x < 2 {
			return x / 5
		}
	}

	return 10000000
}

// finding pixel place for value X or Y
func (g *graphModel) graphGridFindValue(v0 float64, mode string) (Finded int) {
	var arr []float64
	var modeValue int //-1 for to < 0, +1 for from > 0, 0 for to >  && from < 0

	// Creating array of X or Y according mode
	switch mode {
	case "X":
		arr = g.createArrayValue(g.xFrom, g.xTo, float64(g.config.XWindowGraph))
	case "Y":
		arr = g.createArrayValue(g.yFrom, g.yTo, float64(g.config.YWindowGraph))
	default:
		return 0
	}

	// Gap calculation beetwin the closest points in array
	deltaValue := math.Abs(arr[1] - arr[2])
	t.Clg.DeepDebug(fmt.Sprintf("_graphGridFindValue_ mode: |%s|, v0: |%.3f|, deltaValue: |%.5f|", mode, v0, deltaValue))

	// Detection modeValue
	if arr[0] > 0 {
		modeValue = 1
	} else if arr[len(arr)-1] < 0 {
		modeValue = -1
	}

	// Find
	t.Clg.DeepDebug(fmt.Sprintf("_graphGridFindValue_ mode: |%s|, modeValue: |%d|", mode, modeValue))
	if v0 == 0 && modeValue != 0 {
		switch modeValue {
		case -1:
			Finded = 0 - int((0-arr[len(arr)-1])/deltaValue)
		case 1:
			Finded = int(g.config.YWindowGraph) - int(0-arr[0]/deltaValue)
		}
	} else {
		var tempDelta float64 = 2000000
		for i, val := range arr {
			if math.Abs(v0-val) < tempDelta {
				tempDelta = math.Abs(v0 - val)
				Finded = i
			}
		}

		if mode == "Y" {
			Finded = int(g.config.YWindowGraph) - Finded
		}
	}
	t.Clg.DeepDebug(fmt.Sprintf("_graphGridFindValue_ mode: |%s|, v0: |%.3f|, Finded: |%d|", mode, v0, Finded))
	return

	// switch mode {
	// case "X":
	// 	t.Clg.DeepDebug(fmt.Sprintf("_graphGridFindValue_ mode: |%s|, v0: |%.3f|, Finded: |%d|", mode, v0, Finded))
	// 	return
	// case "Y":
	// 	if v0 == 0 && modeValue == 1 {
	// 		t.Clg.DeepDebug(fmt.Sprintf("_graphGridFindValue_ mode: |%s|, v0: |%.3f|, Finded: |%d|", mode, v0, int(g.config.YWindowGraph)-Finded))
	// 		return int(g.config.YWindowGraph) - Finded
	// 	} else if v0 == 0 && modeValue == -1 {
	// 		t.Clg.DeepDebug(fmt.Sprintf("_graphGridFindValue_ mode: |%s|, v0: |%.3f|, Finded: |%d|", mode, v0, 0-Finded))
	// 		return 0 - Finded
	// 	}
	// 	return int(g.config.YWindowGraph) - Finded
	// default:
	// 	return 0
	// }

	// t.Clg.DeepDebug(fmt.Sprintf("_graphGridFindValue_ mode: |%s|, modeValue: |%d|", mode, modeValue))
	// if v0 == 0 && modeValue != 0 {
	// 	switch modeValue {
	// 	case -1:
	// 		Finded = int((0 - arr[len(arr)-1]) / deltaValue)
	// 	case 1:
	// 		Finded = int(0 - arr[0]/deltaValue)
	// 	}
	// } else {
	// 	var tempDelta float64 = 2000000
	// 	for i, val := range arr {
	// 		if math.Abs(v0-val) < tempDelta {
	// 			tempDelta = math.Abs(v0 - val)
	// 			Finded = i
	// 		}
	// 	}
	// }

	// switch mode {
	// case "X":
	// 	t.Clg.DeepDebug(fmt.Sprintf("_graphGridFindValue_ mode: |%s|, v0: |%.3f|, Finded: |%d|", mode, v0, Finded))
	// 	return
	// case "Y":
	// 	if v0 == 0 && modeValue == 1 {
	// 		t.Clg.DeepDebug(fmt.Sprintf("_graphGridFindValue_ mode: |%s|, v0: |%.3f|, Finded: |%d|", mode, v0, int(g.config.YWindowGraph)-Finded))
	// 		return int(g.config.YWindowGraph) - Finded
	// 	} else if v0 == 0 && modeValue == -1 {
	// 		t.Clg.DeepDebug(fmt.Sprintf("_graphGridFindValue_ mode: |%s|, v0: |%.3f|, Finded: |%d|", mode, v0, 0-Finded))
	// 		return 0 - Finded
	// 	}
	// 	return int(g.config.YWindowGraph) - Finded
	// default:
	// 	return 0
	// }
}

// formatting Grid value for print
func (g *graphModel) prepareGridValue(masshtab, val float64) (printNumber string) {
	if masshtab < 0.01 {
		printNumber = strconv.FormatFloat(val, 'f', 3, 64)
	} else if masshtab < 1 {
		printNumber = strconv.FormatFloat(val, 'f', 2, 64)
	} else {
		printNumber = strconv.FormatFloat(val, 'f', 0, 64)
	}
	return printNumber
}

// Printing Grid value for Y
func (g *graphModel) printGridValueY(img draw.Image, masshtabY, val float64, sideModeY bool) {

	printNumber := g.prepareGridValue(masshtabY, val)
	if sideModeY {
		g.drawGridText(img, 50, g.graphGridFindValue(val, "Y"), printNumber)
	} else {
		g.drawGridText(img, g.graphGridFindValue(0, "X"), g.graphGridFindValue(val, "Y"), printNumber)
	}
}

// Printing Grid value for X
func (g *graphModel) printGridValueX(img draw.Image, masshtabX, val float64, sideModeX bool) {
	printNumber := g.prepareGridValue(masshtabX, val)
	if sideModeX {
		g.drawGridText(img, g.graphGridFindValue(val, "X"), 580, printNumber)
	} else {
		g.drawGridText(img, g.graphGridFindValue(val, "X"), g.graphGridFindValue(0, "Y"), printNumber)
	}
}
