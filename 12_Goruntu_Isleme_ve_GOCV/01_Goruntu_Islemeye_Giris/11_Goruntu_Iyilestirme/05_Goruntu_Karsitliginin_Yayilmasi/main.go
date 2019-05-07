package main

import (
	"gocv.io/x/gocv"
)

/*

I'(x,y)=((I(x,y)-min)/(max-min))*255
*/

func main() {

	img := gocv.IMRead("../../MERT_KUBRA_ERDEM.jpg", gocv.IMReadGrayScale)
	imgTotal := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	min, max, _, _ := gocv.MinMaxLoc(img)
	for x := 0; x < img.Rows(); x++ {
		for y := 0; y < img.Cols(); y++ {
			p := img.GetUCharAt(x, y)

			p = uint8(((float32(p) - min) / (max - min)) * 255)
			imgTotal.SetUCharAt(x, y, p)

		}
	}

	window := gocv.NewWindow("Görüntü Karşıtlığını Alma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(imgTotal)
	window.WaitKey(0)

}
