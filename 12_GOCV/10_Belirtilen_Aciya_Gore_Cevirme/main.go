package main

import (
	"fmt"
	"math"

	"gocv.io/x/gocv"
)

/*
  x1 = cosθ(x - x0) - sinθ(y-y0) + x0
  y1 = sinθ(x - x0) - cosθ(y-y0) + y0
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
	/*
		img3 := gocv.NewMat()
		img3 = gocv.GetRotationMatrix2D(image.Point{img.Cols() / 2, img.Rows() / 2}, 45, 1)

		gocv.WarpAffine(img, &img2, img3, image.Point{X: 0, Y: 0})*/

	// Uzun Yöntem :

	x0 := img.Rows() / 2
	y0 := img.Cols() / 2
	angle := 90 * 2 * math.Pi / 360

	for i := 0; i < img.Rows(); i = i + 2 {
		for j := 0; j < img.Cols(); j = j + 2 {

			p := img.GetSCharAt(i, j)

			newX := int(math.Cos(angle)*float64(i-x0) - math.Sin(angle)*float64(j-y0) + float64(x0))
			newY := int(math.Sin(angle)*float64(i-x0) - math.Cos(angle)*float64(j-y0) + float64(y0))

			if newX > 0 && newX < img.Rows() && newY >= 0 && newY < img.Cols() {
				img2.SetUCharAt(newX, newY, uint8(p))
			}
		}
	}
	fmt.Println(img.Rows(), img.Cols())
	fmt.Println(img2.Rows(), img2.Cols())
	window := gocv.NewWindow("Küçültme - Uzaklaştırma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(img2)
	window.WaitKey(0)

}
