package main

import (
	"gocv.io/x/gocv"
)

func main() {

	img := gocv.NewMat()
	img2 := gocv.NewMat()

	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)

	// Yöntem 1 :
	//gocv.Rotate(img, &img, gocv.Rotate180Clockwise)

	// Uzun Yöntem :

	gocv.CvtColor(img, &img, gocv.ColorBGRAToGray)
	img.CopyTo(&img2)

	for i := 0; i < img.Rows(); i++ {
		for j := 0; j < img.Cols(); j++ {

			p := img.GetUCharAt(i, j)

			newX := i
			newY := -j + img.Cols()

			if newX > 0 && newX < img.Rows() && newY >= 0 && newY < img.Cols() {
				img2.SetUCharAt(newX, newY, p)
			}
		}
	}

	window := gocv.NewWindow("Ters Çevirme")
	window.IMShow(img2)
	window.WaitKey(0)

}
