// package main

// import (
// 	"image"
// 	"image/color"

// 	"golang.org/x/tour/pic"
// )

// type Image struct {
// 	w, h int
// }

// func (i Image) Bounds() image.Rectangle {
// 	return image.Rect(0, 0, i.w, i.h)
// }

// func (i Image) ColorModel() color.Model {
// 	color := color.RGBAModel
// 	return color
// }

// func (i Image) At(x, y int) color.Color {
// 	r := uint8(x * x + y)
// 	c := color.RGBA{r, 0, 0, 255}
// 	return c
// }

// func main() {
// 	m := Image{250,250}
// 	pic.ShowImage(m)
// }
