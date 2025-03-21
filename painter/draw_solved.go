package painter

import (
	_ "embed"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"
)

//go:embed ShadeBlue-2OozX.ttf
var fontBytes []byte

const (
	cellSize = 100
	imgSize  = 700
)

var img = image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))

func Draw(field [][]string) {
	customFont := loadFont()
	x1 := 0
	y1 := 0
	x2 := cellSize
	y2 := cellSize

	for _, row := range field {
		for _, cell := range row {
			drawRect(x1, y1, x2, y2, changeColor(cell))
			if !strings.Contains(cell, "{") && !strings.Contains(cell, "[x]") {
				addLabel(x1, y1, cell, customFont)
				drawBorder(x1-1, y1-1, x2-1, y2-1)
			}
			x1 += cellSize
			x2 += cellSize
		}
		x1 = 0
		x2 = cellSize
		y1 += cellSize
		y2 += cellSize

	}

	f, err := os.Create("calendar.png") // Change extension to match format
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}

func drawBorder(x1 int, y1 int, x2 int, y2 int) {
	for y := y1; y <= y2; y++ {
		img.Set(x1, y, color.Black)
	}
	for y := y1; y <= y2; y++ {
		img.Set(x2, y, color.Black)
	}
	for x := x1; x <= x2; x++ {
		img.Set(x, y1, color.Black)
	}
	for x := x1; x <= x2; x++ {
		img.Set(x, y2, color.Black)
	}
}

func changeColor(cell string) color.Color {
	switch cell {
	case "{1}":
		return setColor(255, 0, 0) // RED
	case "{2}":
		return setColor(0, 0, 255) // BLUE
	case "{3}":
		return setColor(105, 215, 0) // ORANGE
	case "{4}":
		return setColor(255, 150, 230) //PINK
	case "{5}":
		return setColor(0, 115, 0) //DARK GREEN
	case "{6}":
		return setColor(0, 0, 110) // DEEP BLUE
	case "{7}":
		return setColor(120, 0, 0) // DARK RED
	case "{8}":
		return setColor(0, 255, 255) //CYAN
	case "[x]":
		return setColor(180, 130, 0) // BROWNY
	default:
		return setColor(180, 130, 0) // lite BROWN
	}
}

func setColor(r uint8, g uint8, b uint8) color.Color {
	return color.RGBA{R: r, G: g, B: b, A: 255}
}

func drawRect(x1, y1, x2, y2 int, col color.Color) {
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			img.Set(x, y, col)
		}
	}
}

func addLabel(x, y int, label string, customFont font.Face) {
	point := fixed.Point26_6{
		X: fixed.I(x + 20),
		Y: fixed.I(y + 75),
	}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.Black),
		Face: customFont,
		Dot:  point,
	}
	d.DrawString(strings.Trim(label, "[]"))
}

func loadFont() font.Face {
	parsedFont, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}

	myFont, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
		Size:    50,
		DPI:     80,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create font face: %v", err)
	}

	return myFont
}
