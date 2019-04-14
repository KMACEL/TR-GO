package main

import (
	"gocv.io/x/gocv"
)

func main() {
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()
	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadColor)

	window.IMShow(img)
	window.WaitKey(0)
}
