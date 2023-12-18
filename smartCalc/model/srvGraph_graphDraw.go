package model

import (
	"fmt"
	"image/color"
	"math"
	d "smartCalc/domains"
)

// Draw graph line at the sCoordinates Array
func (g *graphModel) graphDraw() {

	// Init variable
	var (
		oldErr    bool                                                               // exist old pixel flag
		pixelOldY int                                                                // Y of old pixel
		sdvig     int     = g.graphGridFindValue(0, "Y")                             // find place of ZeroLine
		deltaY    float64 = float64(g.config.YWindowGraph) / math.Abs(g.yFrom-g.yTo) // quantity of pixel to 1 of Y
	)

	// draw graph with circle
	for i, value := range g.gRM.pixelData {

		// if pixel have err==true, skip drawing
		if value.err {
			oldErr = true
			continue
		}
		d.DbgPrint(fmt.Sprint("GraphDraw:", i, value))
		// find y coordinate for draw graph
		y := int(float64(g.config.YWindowGraph) - (value.y * deltaY) - (float64(g.config.YWindowGraph) - float64(sdvig)))

		// cut going beyond y
		if y < 0 {
			y = -1
		} else if y >= int(g.config.YWindowGraph) {
			y = int(g.config.YWindowGraph) + 1
		}

		// if y in windowgraph - draw pixel
		if y > 0-10 && y < int(g.config.YWindowGraph)+10 {
			g.gRM.graphImage.Set(i, y, color.Black)

			// if nessesary, draw vertical line
			if i != 0 {
				if math.Abs(float64(pixelOldY-y)) >= 1.5 && !oldErr {
					g.drawVLine(g.gRM.graphImage, 1, int(math.Abs(float64(pixelOldY-y)))-1, i-1, int(math.Min(float64(pixelOldY), float64(y))), color.Black)
				}
			}
		}

		// remember pixel for next vertical line
		pixelOldY = y
		oldErr = false
	}
}
