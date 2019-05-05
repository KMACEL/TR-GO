package main

import (
	"gocv.io/x/gocv"
)

/*
h = Toplam Yükseklik
f(x,y)=(H - x,y)
*/

func main() {

	img := gocv.NewMat()
	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)

	// Yöntem 1 :
	//gocv.Rotate(img, &img, gocv.Rotate180Clockwise)

	// Uzun Yöntem :
	imgTotal := gocv.NewMat()
	imgTotal = gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8UC3)

	h := img.Rows()
	for x := 0; x < img.Rows(); x++ {
		for y := 0; y < img.Cols(); y++ {

			p := img.GetIntAt(x, y)

			newX := h - x
			newY := y

			imgTotal.SetIntAt(newX, newY, p)

		}
	}

	window := gocv.NewWindow("Ters Çevirme")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(imgTotal)
	window.WaitKey(0)

}
