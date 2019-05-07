package main

import (
	"gocv.io/x/gocv"
)

/*
I'(x,y) = I(x,y) / K
*/

func main() {

	img := gocv.IMRead("../../MERT_KUBRA_ERDEM.jpg", gocv.IMReadGrayScale)
	imgTotal := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	// Yöntem 1 :
	//gocv.ConvertScaleAbs(img, &imgTotal, 0.2, 1.0)

	// Yöntem 2 :
	value := 5
	for x := 0; x < img.Rows(); x++ {
		for y := 0; y < img.Cols(); y++ {
			p := img.GetUCharAt(x, y)

			p = p / uint8(value)
			imgTotal.SetUCharAt(x, y, p)
		}
	}

	window := gocv.NewWindow("Contrast Azaltma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(imgTotal)
	window.WaitKey(0)
}
