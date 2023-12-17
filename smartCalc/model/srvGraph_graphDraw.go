package model

import (
	"image/color"
	"math"
)

// Draw graph line at the sCoordinates Array
func (g *graphModel) graphDraw() {
	var oldErr bool
	var pixelOldY int
	var sdvig int = g.graphGridFindValue(0, "Y")
	var deltaY float64 = float64(g.config.YWindowGraph) / math.Abs(g.yFrom-g.yTo)

	for i, value := range g.gRM.pixelData {
		if value.err {
			oldErr = true
			continue
		}
		y := int(float64(g.config.YWindowGraph) - (value.y * deltaY) - (float64(g.config.YWindowGraph) - float64(sdvig)))
		g.gRM.graphImage.Set(i, y, color.Black)
		if i != 0 {
			if math.Abs(float64(pixelOldY-y)) >= 1.5 && !oldErr {
				g.drawVLine(g.gRM.graphImage, 1, int(math.Abs(float64(pixelOldY-y)))-1, i-1, int(math.Min(float64(pixelOldY), float64(y))), color.Black)
				pixelOldY = y
				oldErr = false
			}
		} else {
			pixelOldY = y
			oldErr = false
		}
	}
}
