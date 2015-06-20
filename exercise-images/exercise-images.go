package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	Height, Width int
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.Height, m.Width)
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) At(x, y int) color.Color {
	v := uint8(x ^ y)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{255, 255}
	pic.ShowImage(m)
}
