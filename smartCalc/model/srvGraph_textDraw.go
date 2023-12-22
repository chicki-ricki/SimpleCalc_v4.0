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

	t "smartCalc/tools"
)

// Print Logo at the picture
func (g *graphModel) drawLogo(dst draw.Image, yBase int, Text string) {

	// Read font from config path
	fontBytes, err := os.ReadFile(g.config.TypePath)
	if err != nil {
		t.Clg.Warning(fmt.Sprintf("_drawLogo_ cannot read logo font from path: %s", g.config.TypePath))
		return
	}

	// Parce font
	fnt, err := truetype.Parse(fontBytes)
	if err != nil {
		t.Clg.Warning(fmt.Sprintf("_drawLogo_ cannot parce font from path: %s: %v", g.config.TypePath, err))
		return
	}

	// Create font drawer
	d := font.Drawer{
		Dst: dst,
		Src: image.NewUniform(color.RGBA{G: 0x88, B: 0xAA, A: 0xFF}),
		Face: truetype.NewFace(fnt, &truetype.Options{
			Size: float64(20),
			DPI:  72,
		}),
	}

	// Calculate draw pisition
	d.Dot = fixed.Point26_6{
		X: (fixed.I(int(g.config.XWindowGraph)) - d.MeasureString(Text) - fixed.I(17)),
		Y: fixed.I((yBase + (36 / 2) - 12)),
	}

	// Drawing text string
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
