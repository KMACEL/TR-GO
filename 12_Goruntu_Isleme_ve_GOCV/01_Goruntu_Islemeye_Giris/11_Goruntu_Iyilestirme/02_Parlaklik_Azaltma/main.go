package main

import (
	"gocv.io/x/gocv"
)

/*
I'(x,y) = I(x,y) - K
*/

func main() {

	img := gocv.IMRead("../../MERT_KUBRA_ERDEM.jpg", gocv.IMReadGrayScale)
	imgTotal := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	// Yöntem 1 :
	//gocv.ConvertScaleAbs(img, &imgTotal, 1.0, -100.0)

	// Yöntem 2 :
	value := 100
	for i := 0; i < img.Rows(); i++ {
		for j := 0; j < img.Cols(); j++ {
			p := img.GetUCharAt(i, j)

			if t := int(p) - value; t < 0 {
				imgTotal.SetUCharAt(i, j, 0)
			} else {
				imgTotal.SetUCharAt(i, j, uint8(t))
			}
		}
	}

	window := gocv.NewWindow("Parlaklık Azaltma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(imgTotal)
	window.WaitKey(0)
}
