package main

import (
	"gocv.io/x/gocv"
)

func main() {
	img := gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadGrayScale)
	imgTotal := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	// Yöntem 1 :
	/*rotationMatrixSettings := gocv.GetRotationMatrix2D(image.Point{0, 0}, 0.0, 2)
	gocv.WarpAffine(img, &imgTotal, rotationMatrixSettings, image.Point{X: 0, Y: 0})*/

	// Uzun Yöntem :
	for x := 0; x <= img.Rows(); x++ {
		for y := 0; y <= img.Cols(); y++ {

			p := img.GetUCharAt(x, y)

			if x*2 >= 0 && x*2 <= img.Rows() && y*2 >= 0 && y*2 <= img.Cols() {
				imgTotal.SetUCharAt(x*2, y*2, p)
			}

			if x*2+1 >= 0 && x*2+1 <= img.Rows() && y*2 >= 0 && y*2 <= img.Cols() {
				imgTotal.SetUCharAt(x*2+1, y*2, p)
			}

			if x*2 >= 0 && x*2 <= img.Rows() && y*2+1 >= 0 && y*2+1 <= img.Cols() {
				imgTotal.SetUCharAt(x*2, y*2+1, p)
			}

			if x*2+1 >= 0 && x*2+1 <= img.Rows() && y*2+1 >= 0 && y*2+1 <= img.Cols() {
				imgTotal.SetUCharAt(x*2+1, y*2+1, p)
			}
		}
	}

	window := gocv.NewWindow("Büyültme-Yakınlaştırma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(imgTotal)
	window.WaitKey(0)
}
