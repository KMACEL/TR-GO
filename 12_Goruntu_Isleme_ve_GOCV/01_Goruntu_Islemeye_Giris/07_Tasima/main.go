package main

import (
	"gocv.io/x/gocv"
)

/*
(x,y) = (x+10,y+10)
*/

func main() {
	img := gocv.NewMat()
	img2 := gocv.NewMat()

	defer img.Close()
	defer img2.Close()

	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)
	img2 = gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	gocv.CvtColor(img, &img, gocv.ColorBGRAToGray)

	//img2.SetTo(gocv.Scalar{255, 255, 255, 255})

	// Yöntem 1 :
	//img3 := gocv.NewMat()
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
