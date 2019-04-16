package main

import (
	"gocv.io/x/gocv"
)

func main() {
	img := gocv.NewMat()
	img2 := gocv.NewMat()

	defer img.Close()
	defer img2.Close()

	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)
	img2 = gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	gocv.CvtColor(img, &img, gocv.ColorBGRAToGray)

	// Yöntem 1 :
	/*img3 := gocv.NewMat()
	img3 = gocv.GetRotationMatrix2D(image.Point{0, 0}, 0.0, 0.5)

	gocv.WarpAffine(img, &img2, img3, image.Point{X: 0, Y: 0})*/

	// Uzun Yöntem :
	var x, y int
	for i := 0; i < img.Rows(); i = i + 2 {
		y = 0
		for j := 0; j < img.Cols(); j = j + 2 {

			p := img.GetUCharAt(i, j)

			img2.SetUCharAt(x, y, p)
			y++
		}
		x++
	}

	window := gocv.NewWindow("Küçültme - Uzaklaştırma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(img2)
	window.WaitKey(0)

}
