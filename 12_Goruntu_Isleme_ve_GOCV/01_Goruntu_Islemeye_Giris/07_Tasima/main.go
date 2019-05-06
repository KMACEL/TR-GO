package main

import (
	"gocv.io/x/gocv"
)

/*
(x,y) = (x+A,y+B)
*/

func main() {

	img := gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadGrayScale)
	imgTotal := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	a := 200
	b := 300

	for x := 0; x < img.Rows(); x++ {
		for y := 0; y < img.Cols(); y++ {

			p := img.GetUCharAt(x, y)

			newX := x + a
			newY := y + b

			if newX >= 0 && newX <= imgTotal.Rows() && newY >= 0 && newY <= imgTotal.Cols() {
				imgTotal.SetUCharAt(newX, newY, p)
			}
		}
	}

	window := gocv.NewWindow("Taşıma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(imgTotal)
	window.WaitKey(0)
}
