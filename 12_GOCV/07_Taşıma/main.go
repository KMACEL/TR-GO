package main

import (
	"gocv.io/x/gocv"
)

func main() {
	img := gocv.NewMat()
	img2 := gocv.NewMat()

	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)
	img2 = gocv.NewMatWithSize(1500, 1500, gocv.MatTypeCV8U)

	gocv.CvtColor(img, &img, gocv.ColorBGRAToGray)

	// Yöntem 1 :
	//img3 := gocv.NewMat()
	//gocv.Rotate(img, &img, gocv.Rotate180Clockwise)
	//img3 = gocv.GetRotationMatrix2D(image.Point{500, 700}, 0.0, 0.8)
	//gocv.WarpAffine(img, &img2, img3, image.Point{img2.Rows(), img2.Cols()})

	// Uzun Yöntem :
	for i := 0; i < img.Rows(); i++ {
		for j := 0; j < img.Cols(); j++ {

			p := img.GetUCharAt(i, j)

			newX := i + 200
			newY := j + 300

			if newX > 0 && newX < img.Rows() && newY >= 0 && newY < img.Cols() {
				img2.SetUCharAt(newX, newY, p)
			}
		}
	}

	window := gocv.NewWindow("Taşıma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(img2)
	window.WaitKey(0)

}
