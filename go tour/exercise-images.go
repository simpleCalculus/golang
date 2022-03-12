package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)


type Image struct{
	Height, Width int
	colr uint8
}

func (im Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, im.Width, im.Height)
}

func  (im Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (im Image) At(x, y int) color.Color {
	return color.RGBA{im.colr+uint8(x), im.colr+uint8(y), 255, 255}
}


func main() {
	m := Image{200, 200, 135}
	pic.ShowImage(m)
}
