package main

import (
	"gocv.io/x/gocv"
)

/*
I'(x,y) = I(x,y) * K
*/

func main() {

	img := gocv.IMRead("../../MERT_KUBRA_ERDEM.jpg", gocv.IMReadGrayScale)
	imgTotal := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	// Yöntem 1 :
	//gocv.ConvertScaleAbs(img, &imgTotal, 5.0, 1.0)

	// Yöntem 2 :
	value := 5
	for x := 0; x < img.Rows(); x++ {
		for y := 0; y < img.Cols(); y++ {
			p := img.GetUCharAt(x, y)

			if t := int(p) * value; t > 255 {
				imgTotal.SetUCharAt(x, y, 255)
			} else {
				imgTotal.SetUCharAt(x, y, uint8(t))
			}
		}
	}

	window := gocv.NewWindow("Contrast Arttırmak")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(imgTotal)
	window.WaitKey(0)

}
