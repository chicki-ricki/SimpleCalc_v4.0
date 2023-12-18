package model

import (
	"math"
)

// graph prepared string - check and stapples(without replace X)
func (g *graphModel) graphPrepareString() bool {
	g.equal.equation.equation = g.equalValue
	if str, err := g.equal.equation.onlyCheck(); err != nil {
		return true
	} else {
		// g.preparedEquation = str
		g.preparedEquation = g.equal.addStaplesForX(str)
	}
	return false
}

// calculate pixels data for graph (requared g.preparedEquation)
func (g *graphModel) calculateData() {

	// find gap between pixels
	var deltaPixel float64 = math.Abs(g.xFrom-g.xTo) / float64(g.config.XWindowGraph)
	var xr float64

	// filling pixels array with calculated data
	for x := math.Min(g.xFrom, g.xTo); x <= math.Max(g.xFrom, g.xTo); x += deltaPixel {
		xr = x
		if (x < 0.00001 && x > 0) || (x > 0-0.00001 && x < 0) {
			xr = 0
		}
		g.gRM.pixelData = append(g.gRM.pixelData, g.equal.calculate(g.preparedEquation, xr))
	}

	// find MIN and MAX in Y
	for _, val := range g.gRM.pixelData {
		if !val.err {
			g.gRM.yGraphMin = math.Min(g.gRM.yGraphMin, val.y)
			g.gRM.yGraphMax = math.Max(g.gRM.yGraphMax, val.y)

		}
	}
}
