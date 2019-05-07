package main

import (
	"gocv.io/x/gocv"
)

/*
f(x,y)=((x,y)+(x+1,y)+(x,y+1)+(x+1,y+1))/4
*/

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

	img := gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadGrayScale)
	imgTotal := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)

	// Yöntem 1 :
	/*rotationMatrixSettings := gocv.GetRotationMatrix2D(image.Point{0, 0}, 0.0, 0.5)
	gocv.WarpAffine(img, &imgTotal, rotationMatrixSettings, image.Point{X: 0, Y: 0})*/

	// Yöntem 2 :
	for x := 0; x < img.Rows(); x++ {
		for y := 0; y < img.Cols(); y++ {

			var p, p2, p3, p4 uint8
			if x*2 >= 0 && x*2 <= img.Rows() && y*2 >= 0 && y*2 <= img.Cols() {
				p = img.GetUCharAt(x*2, y*2)
			}

			if x*2+1 >= 0 && x*2+1 <= img.Rows() && y*2 >= 0 && y*2 <= img.Cols() {
				p2 = img.GetUCharAt(x*2+1, y*2)
			}

			if x*2 >= 0 && x*2 <= img.Rows() && y*2+1 >= 0 && y*2+1 <= img.Cols() {
				p3 = img.GetUCharAt(x*2, y*2+1)
			}

			if x*2+1 >= 0 && x*2+1 <= img.Rows() && y*2+1 >= 0 && y*2+1 <= img.Cols() {
				p4 = img.GetUCharAt(x*2+1, y*2+1)
			}
			total := (int(p) + int(p2) + int(p3) + int(p4)) / 4
			imgTotal.SetUCharAt(x, y, uint8(total))
		}
	}
	window := gocv.NewWindow("Küçültme - Uzaklaştırma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(imgTotal)
	window.WaitKey(0)
}
