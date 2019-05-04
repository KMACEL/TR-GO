package main

import (
	"gocv.io/x/gocv"
)

/*

I'(x,y)=((I(x,y)-min)/(max-min))*255
*/

func main() {
	img := gocv.NewMat()
	img2 := gocv.NewMat()

	defer img.Close()
	defer img2.Close()

	img = gocv.IMRead("../../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)
	img2 = gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	gocv.CvtColor(img, &img, gocv.ColorBGRAToGray)

	for i := 0; i < img.Rows(); i++ {
		for j := 0; j < img.Cols(); j++ {
			p := img.GetUCharAt(i, j)

			if p > 128 {
				p = 255
			} else {
				p = 0
			}

			img2.SetUCharAt(i, j, p)

		}
	}

	window := gocv.NewWindow("Görüntü Karşıtlığının Yayılması")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(img2)
	window.WaitKey(0)

}
