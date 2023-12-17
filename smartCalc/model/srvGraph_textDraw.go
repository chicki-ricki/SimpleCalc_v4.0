package model

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"os"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

// Print Logo at the picture
func (g *graphModel) drawLogo(dst draw.Image, yBase int, Text string) {

	fontBytes, _ := os.ReadFile(g.config.TypePath)
	fnt, _ := truetype.Parse(fontBytes)

	d := font.Drawer{
		Dst: dst,
		Src: image.NewUniform(color.RGBA{G: 0x88, B: 0xAA, A: 0xFF}),
		Face: truetype.NewFace(fnt, &truetype.Options{
			Size: float64(20),
			DPI:  72,
		}),
	}
	d.Dot = fixed.Point26_6{
		X: (fixed.I(int(g.config.XWindowGraph)) - d.MeasureString(Text) - fixed.I(17)),
		Y: fixed.I((yBase + (36 / 2) - 12)),
	}
	d.DrawString(Text)
}

// Creating text with equal data for image
func (g *graphModel) createEqualText() string {
	return fmt.Sprintf("%s  X{%.2f .. %.2f} Y{%.2f .. %.2f}",
		g.equalValue,
		g.xFrom,
		g.xTo,
		g.gRM.yGraphMin,
		g.gRM.yGraphMax)
}

// Print equal data for the graph
func (g *graphModel) drawEqualText(dst draw.Image, xBase, yBase int, Text string) {
	d := font.Drawer{
		Dst:  dst,
		Src:  image.NewUniform(color.Gray{Y: uint8(125)}),
		Face: basicfont.Face7x13,
	}
	d.Dot = fixed.Point26_6{
		X: fixed.I(xBase),
		Y: fixed.I((yBase + (12 / 2) - 12)),
	}
	d.DrawString(Text)
}

// Print Grid value
func (g *graphModel) drawGridText(dst draw.Image, xBase int, yBase int, Text string) {
	d := font.Drawer{
		Dst:  dst,
		Src:  image.NewUniform(color.Gray{Y: uint8(125)}),
		Face: basicfont.Face7x13,
	}
	d.Dot = fixed.Point26_6{
		X: (fixed.I(xBase) - d.MeasureString(Text) - fixed.I(5)),
		Y: fixed.I(int(yBase + 14)),
	}
	d.DrawString(Text)
}
