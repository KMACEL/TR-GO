package main

import (
	"gocv.io/x/gocv"
)

func main() {

	img := gocv.IMRead("../../MERT_KUBRA_ERDEM.jpg", gocv.IMReadGrayScale)
	imgTotal := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	for x := 0; x < img.Rows(); x++ {
		for y := 0; y < img.Cols(); y++ {
			p := img.GetUCharAt(x, y)

			if p > 127 {
				p = 255
			} else {
				p = 0
			}

			imgTotal.SetUCharAt(x, y, p)
		}
	}

	window := gocv.NewWindow("Görüntü Eşitleme")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(imgTotal)
	window.WaitKey(0)

}
