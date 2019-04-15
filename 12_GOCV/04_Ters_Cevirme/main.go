package main

import (
	"gocv.io/x/gocv"
)

func main() {

	img := gocv.NewMat()
	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)

	// Yöntem 1 :
	gocv.Rotate(img, &img, gocv.Rotate180Clockwise)

	// Uzun Yöntem :
	/*
		var angle float64 = 180 * 2 * math.Pi / 360
		var x0 float64 = float64(img.Rows()) / 2
		var y0 float64 = float64(img.Cols()) / 2

		for i := 0; i < img.Rows(); i++ {
			for j := 0; j < img.Cols(); j++ {
				p := img.GetIntAt(i, j)
				x2 := int(math.Cos(angle)*(float64(i)-x0) - math.Sin(angle)*(float64(j)-y0) + x0)
				y2 := int(math.Sin(angle)*(float64(i)-x0) + math.Cos(angle)*(float64(j)-y0) + y0)
			}
		}*/

	window := gocv.NewWindow("Ters Çevirme")
	window.IMShow(img)
	window.WaitKey(0)

}
