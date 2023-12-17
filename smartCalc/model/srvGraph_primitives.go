package model

import (
	"image/color"
	"image/draw"
)

func (g *graphModel) fillBackground(img draw.Image, c color.Color) {
	for x := 0; x < int(g.config.XWindowGraph); x++ {
		for y := 0; y < int(g.config.YWindowGraph); y++ {
			img.Set(x, y, c)
		}
	}
}

// Draw vertical Up arrow
func (g *graphModel) arrowV(img draw.Image, x0, y0 int, c color.Color) {

	for j := 0; j < 20; j++ {
		for i := -j / 3; i <= j/3; i++ {
			img.Set(x0+i, y0+j, c)
		}
	}
}

// Draw horisontal right arrow
func (g *graphModel) arrowH(img draw.Image, x0, y0 int, c color.Color) {
	for j := 0; j < 20; j++ {
		for i := -j / 3; i <= j/3; i++ {
			img.Set(x0-j, y0+i, c)
		}
	}
}

// Draw vertical line
func (g *graphModel) drawVLine(img draw.Image, width, heigth, x0, y0 int, c color.Color) {
	for y := y0; y < int(g.config.YWindowGraph) && y < y0+heigth; y++ {
		for i := 0; i < width; i++ {
			img.Set(x0+i, y, c)
		}
	}
}

// Draw horisontal line
func (g *graphModel) drawHLine(img draw.Image, width, lenght, x0, y0 int, c color.Color) {
	for x := x0; x < int(g.config.XWindowGraph) && x < x0+lenght; x++ {
		for i := 0; i < width; i++ {
			img.Set(x, y0+i, c)
		}
	}
}
