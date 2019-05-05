package main

import (
	"gocv.io/x/gocv"
)

func main() {
	window := gocv.NewWindow("Merhaba")
	img := gocv.NewMat()
	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadColor)

	window.IMShow(img)
	window.WaitKey(0)
}
