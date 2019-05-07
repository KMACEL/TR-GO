package main

import (
	"gocv.io/x/gocv"
)

/*
W = Toplam Genişlik
f(x,y)=(x,W - y)
*/

func main() {

	img := gocv.NewMat()
	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadGrayScale)

	// Yöntem 1 :
	//gocv.Flip(img, &img, 1)

	// Yöntem 2 :
	imgTotal := gocv.NewMat()
	imgTotal = gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)
	w := img.Cols()
	for x := 0; x < img.Rows(); x++ {
		for y := 0; y < img.Cols(); y++ {

			p := img.GetUCharAt(x, y)

			newX := x
			newY := w - y

			imgTotal.SetUCharAt(newX, newY, p)

		}
	}

	window := gocv.NewWindow("Aynalama")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(imgTotal)
	window.WaitKey(0)

}
