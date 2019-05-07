package main

import (
	"gocv.io/x/gocv"
)

/*
I'(x,y) = I(x,y) + K
*/

func main() {

	img := gocv.IMRead("../../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)
	imgTotal := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	gocv.CvtColor(img, &img, gocv.ColorBGRAToGray)

	// Yöntem 1 :
	// Beta değeri parlaklık, Alpha değeri contrast
	//gocv.ConvertScaleAbs(img, &imgTotal, 1.0, 100.0)

	// Yöntem 2 :
	value := 100
	for x := 0; x < img.Rows(); x++ {
		for y := 0; y < img.Cols(); y++ {
			p := img.GetUCharAt(x, y)

			if t := int(p) + value; t > 255 {
				imgTotal.SetUCharAt(x, y, 255)
			} else {
				imgTotal.SetUCharAt(x, y, uint8(t))
			}
		}
	}

	window := gocv.NewWindow("Parlaklık Arttırmak")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(imgTotal)
	window.WaitKey(0)

}
