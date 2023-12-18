package model

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	d "smartCalc/domains"
	// t "smartCalc/tools"
	"strconv"
	"strings"
)

type SCoordinates struct {
	err bool
	x   float64
	y   float64
}

type graphResultModel struct {
	yGraphMin, yGraphMax float64
	pixelData            []SCoordinates
	graphImage           draw.Image
}

type graphModel struct {
	err error

	equalValue                         string  // string of equal for equal or graph
	xFromStr, xToStr, yFromStr, yToStr string  // raw data
	xFrom, xTo, yFrom, yTo             float64 // border Value for graph
	preparedEquation                   string
	gRM                                graphResultModel
	equal                              equalModel
	config                             *d.Cfg
}

//---------- New object and interface implementing

func (m *CalcModel) NewGraph(in ModelsInput) *graphModel {
	var graph graphModel
	graph.equalValue = in.ModelGraphData.EqualValue
	graph.xFromStr = in.ModelGraphData.XFromStr
	graph.xToStr = in.ModelGraphData.XToStr
	graph.yFromStr = in.ModelGraphData.YFromStr
	graph.yToStr = in.ModelGraphData.YToStr
	if graph.entrysGraphCheck(in) {
		graph.err = errors.New("Invalid graphborder data")
	}

	//update inputData for Creating Equal object
	in.ModelEqualData.EqualValue = in.ModelGraphData.EqualValue
	in.ModelEqualData.XEqualStr = in.ModelGraphData.XFromStr
	graph.equal = *NewEqual(in)
	graph.config = m.Config
	return &graph
}

// Check and copy input data for graph
func (g *graphModel) entrysGraphCheck(in ModelsInput) bool {
	var err error

	if g.xFrom, err = strconv.ParseFloat(in.ModelGraphData.XFromStr, 64); err != nil {
		return true
	} else if g.xTo, err = strconv.ParseFloat(in.ModelGraphData.XToStr, 64); err != nil {
		return true
	} else if g.yFrom, err = strconv.ParseFloat(in.ModelGraphData.YFromStr, 64); err != nil {
		return true
	} else if g.yTo, err = strconv.ParseFloat(in.ModelGraphData.YToStr, 64); err != nil {
		return true
	} else if strings.EqualFold(in.ModelGraphData.EqualValue, "") {
		return true
	}

	positionMinMax(&g.xFrom, &g.xTo)
	positionMinMax(&g.yFrom, &g.yTo)

	if g.yTo == g.yFrom || g.xTo == g.xFrom ||
		g.yFrom < -1000000 || g.xFrom < -1000000 ||
		g.yTo > 1000000 || g.xTo > 1000000 {
		return true
	}
	return false
}

func (g *graphModel) setError(out *ModelsOutput) *ModelsOutput {
	out.Err = true
	out.ModelGraphResult.Err = true
	out.ModelGraphResult.ResultStr = "Error"
	return out
}

// Implementing request interface for graphModel
func (g *graphModel) GetResult() (out ModelsOutput) {
	out.Mode = 2
	out.ModelGraphResult.Mode = 2

	if g.err != nil || g.graphPrepareString() {
		return *g.setError(&out)
	}

	g.calculateData()
	g.graphImageBuild()
	out.ModelGraphResult.ResultStr = fmt.Sprintf("Y {%.2f .. %.2f}", g.gRM.yGraphMin, g.gRM.yGraphMax)

	out.ModelGraphResult.GraphImage = g.gRM.graphImage
	// if t.ExportImageToPng(g.gRM.graphImage, g.config.TempFileDir+g.config.TempGraph) != nil {
	// 	return *g.setError(&out)
	// }

	return
}

// Creating Image with graph
func (g *graphModel) graphImageBuild() {

	// Create New Image RGBA object
	g.gRM.graphImage = image.NewRGBA(image.Rect(0, 0, int(g.config.XWindowGraph), int(g.config.YWindowGraph)))

	// Fill white background
	g.fillBackground(g.gRM.graphImage, color.White)

	// Draw Grid for graph
	g.graphGridDraw(g.gRM.graphImage)

	// Draw graph
	g.graphDraw()

	// Draw additional graph information to image
	g.drawEqualText(g.gRM.graphImage, 20, 30, g.createEqualText())

	// Draw Logo to image
	g.drawLogo(g.gRM.graphImage, 21, "CleverCalc")
	return
}
