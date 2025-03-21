package main

import (
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
	"strings"
)

const (
	cellSize = 100
	imgSize  = 700
)

var img = image.NewRGBA(image.Rect(0, 0, imgSize, imgSize))

func Draw(field [][]string, font font.Face) {
	x1 := 0
	y1 := 0
	x2 := cellSize
	y2 := cellSize

	for _, row := range field {
		for _, cell := range row {
			drawRect(x1, y1, x2, y2, changeColor(cell))
			if !strings.Contains(cell, "{") && !strings.Contains(cell, "[x]") {
				addLabel(x1, y1, cell, font)
			}
			x1 += cellSize
			x2 += cellSize
		}
		x1 = 0
		x2 = cellSize
		y1 += cellSize
		y2 += cellSize

	}

	f, err := os.Create("draw.jpg") // Change extension to match format
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = jpeg.Encode(f, img, nil)
	if err != nil {
		panic(err)
	}
}

func LoadFont() font.Face {
	fontBytes, err := os.ReadFile("./ShadeBlue-2OozX.ttf")
	if err != nil {
		log.Fatalf("failed to read font file: %v", err)
	}

	parsedFont, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("failed to parse font: %v", err)
	}

	myFont, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
		Size:    70,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("failed to create font face: %v", err)
	}

	return myFont
}

func changeColor(cell string) color.Color {
	switch cell {
	case "{1}":
		return color.RGBA{R: 255, A: 255}
	case "{2}":
		return color.RGBA{B: 255, A: 255}
	case "{3}":
		return color.RGBA{G: 255, A: 255}
	case "{4}":
		return color.RGBA{R: 255, G: 255, A: 255}
	case "{5}":
		return color.RGBA{R: 255, B: 255, A: 255}
	case "{6}":
		return color.RGBA{G: 255, B: 255, A: 255}
	case "{7}":
		return color.RGBA{R: 255, G: 100, B: 200, A: 255}
	case "{8}":
		return color.RGBA{R: 10, G: 100, B: 255, A: 255}
	case "[x]":
		return color.RGBA{R: 150, G: 70, A: 255}
	default:
		return color.White
	}
}

func fillImage(c color.Color) {
	for y := 0; y < img.Bounds().Dy(); y++ {
		for x := 0; x < img.Bounds().Dx(); x++ {
			img.Set(x, y, c)
		}
	}
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
		X: fixed.I(x + 10),
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
