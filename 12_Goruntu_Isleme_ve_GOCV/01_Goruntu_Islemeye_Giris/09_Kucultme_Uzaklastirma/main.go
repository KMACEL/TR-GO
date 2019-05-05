package main

import (
	"gocv.io/x/gocv"
)

/*
	-----------------
	|p1 |p2 |p3 |p4 |    	-----------------------------------------
	-----------------  		|(p1+p2+p5+p6)/4    |   (p3+p4+p7+p8)/4  |
	|p5 |p6 |p7 |p8 | =		|(p9+p10+p13+p16)/4 |(p11+p12+p15+p16)/4 |
	-----------------       -----------------------------------------
	|p9 |p10|p11|p12|
	-----------------
	|p13|p14|p15|p16|
	-----------------
*/

func main() {
	img := gocv.NewMat()
	img2 := gocv.NewMat()

	defer img.Close()
	defer img2.Close()

	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)
	img2 = gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	gocv.CvtColor(img, &img, gocv.ColorBGRAToGray)

	// Yöntem 1 :
	/*img3 := gocv.NewMat()
	img3 = gocv.GetRotationMatrix2D(image.Point{0, 0}, 0.0, 0.5)

	gocv.WarpAffine(img, &img2, img3, image.Point{X: 0, Y: 0})*/

	// Uzun Yöntem :
	var x, y int
	for i := 0; i < img.Rows(); i = i + 2 {
		y = 0
		for j := 0; j < img.Cols(); j = j + 2 {

			p := img.GetUCharAt(i, j)
			p2 := img.GetUCharAt(i+1, j)
			p3 := img.GetUCharAt(i, j+1)
			p4 := img.GetUCharAt(i+1, j+1)
			t := (int(p) + int(p2) + int(p3) + int(p4)) / 4

			img2.SetUCharAt(x, y, uint8(t))
			y++
		}
		x++
	}

	window := gocv.NewWindow("Küçültme - Uzaklaştırma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(img2)
	window.WaitKey(0)

}
