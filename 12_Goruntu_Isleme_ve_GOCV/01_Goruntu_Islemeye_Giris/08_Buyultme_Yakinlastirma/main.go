package main

import (
	"gocv.io/x/gocv"
)

/*
f(x*2,y*2)=(x,y)
f(x*2+1,y*2)=(x,y)
f(x*2,y*2+1)=(x,y)
f(x*2+1,y*2+1)=(x,y)
*/

func main() {
	img := gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadGrayScale)
	imgTotal := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	// Yöntem 1 :
	/*rotationMatrixSettings := gocv.GetRotationMatrix2D(image.Point{0, 0}, 0.0, 2)
	gocv.WarpAffine(img, &imgTotal, rotationMatrixSettings, image.Point{X: 0, Y: 0})*/

	// Yöntem 2 :
	zoomValue := 2
	for x := 0; x <= img.Rows(); x++ {
		for y := 0; y <= img.Cols(); y++ {

			p := img.GetUCharAt(x, y)

			if x*zoomValue >= 0 && x*zoomValue <= img.Rows() && y*zoomValue >= 0 && y*zoomValue <= img.Cols() {
				imgTotal.SetUCharAt(x*zoomValue, y*zoomValue, p)
			}

			if x*zoomValue+1 >= 0 && x*zoomValue+1 <= img.Rows() && y*zoomValue >= 0 && y*zoomValue <= img.Cols() {
				imgTotal.SetUCharAt(x*zoomValue+1, y*zoomValue, p)
			}

			if x*zoomValue >= 0 && x*zoomValue <= img.Rows() && y*zoomValue+1 >= 0 && y*zoomValue+1 <= img.Cols() {
				imgTotal.SetUCharAt(x*zoomValue, y*zoomValue+1, p)
			}

			if x*zoomValue+1 >= 0 && x*zoomValue+1 <= img.Rows() && y*zoomValue+1 >= 0 && y*zoomValue+1 <= img.Cols() {
				imgTotal.SetUCharAt(x*zoomValue+1, y*zoomValue+1, p)
			}
		}
	}

	window := gocv.NewWindow("Büyültme-Yakınlaştırma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(imgTotal)
	window.WaitKey(0)
}
