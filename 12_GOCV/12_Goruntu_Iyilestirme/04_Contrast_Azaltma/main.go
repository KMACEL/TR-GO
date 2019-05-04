package main

import (
	"gocv.io/x/gocv"
)

/*
I'(x,y) = I(x,y) / K
*/

func main() {
	img := gocv.NewMat()
	img2 := gocv.NewMat()

	defer img.Close()
	defer img2.Close()

	img = gocv.IMRead("../../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)
	img2 = gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	gocv.CvtColor(img, &img, gocv.ColorBGRAToGray)

	// Yöntem 1 :
	//gocv.ConvertScaleAbs(img, &img2, 0.2, 1.0)

	// Uzun Yöntem :
	value := 5
	for i := 0; i < img.Rows(); i++ {
		for j := 0; j < img.Cols(); j++ {
			p := img.GetUCharAt(i, j)

			p = p / uint8(value)
			img2.SetUCharAt(i, j, p)

		}
	}

	window := gocv.NewWindow("Parlaklık Arttırmak")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(img2)
	window.WaitKey(0)

}
