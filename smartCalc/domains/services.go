package domains

import "image/draw"

type EqualModel struct {
	err         error
	EqualString string
	Prepared    string
	XEqualStr   string
	ResultStr   string
	XEqual      float64
	Result      float64
	Equation    EquationModel
}

type EquationModel struct {
	Err       error
	Equation  string
	Checked   string
	Prepared  string
	ResultStr string
	Result    float64
}

type SCoordinates struct {
	err bool
	x   float64
	y   float64
}

type GraphResultModel struct {
	yGraphMin, yGraphMax float64
	pixelData            []SCoordinates
	graphImage           draw.Image
}

type GraphModel struct {
	err                                error
	equalValue                         string  // string of equal for equal or graph
	xFromStr, xToStr, yFromStr, yToStr string  // raw data
	xFrom, xTo, yFrom, yTo             float64 // border Value for graph
	preparedEquation                   string
	gRM                                GraphResultModel
	Equal                              EqualModel
}
