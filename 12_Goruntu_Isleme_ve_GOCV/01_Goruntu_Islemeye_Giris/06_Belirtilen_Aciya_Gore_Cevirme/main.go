package main

import (
	"math"

	"gocv.io/x/gocv"
)

/*
  x1 = cosθ(x - x0) - sinθ(y-y0) + x0
  y1 = sinθ(x - x0) - cosθ(y-y0) + y0
*/

func main() {
	img := gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadGrayScale)

	imgTotal := gocv.NewMatWithSize(img.Rows(), img.Cols(), gocv.MatTypeCV8U)
	// Yöntem 1 :
	// Açı + olursa saat yönünün tersi, - olursa saat yönünde döner
	/*rotationMatrixSettings := gocv.GetRotationMatrix2D(image.Point{X: img.Cols() / 2, Y: img.Rows() / 2}, -120, 1)
	gocv.WarpAffine(img, &imgTotal, rotationMatrixSettings, image.Point{X: 0, Y: 0})*/

	// Yöntem 2 :

	x0 := img.Rows() / 2
	y0 := img.Cols() / 2

	angleValue := 100.0
	angle := angleValue * 2 * math.Pi / 360

	for i := 0; i < img.Rows(); i++ {
		for j := 0; j < img.Cols(); j++ {

			p := img.GetUCharAt(i, j)

			newX := int(math.Cos(angle)*float64(i-x0) - math.Sin(angle)*float64(j-y0) + float64(x0))
			newY := int(math.Sin(angle)*float64(i-x0) - math.Cos(angle)*float64(j-y0) + float64(y0))

			if newX >= 0 && newX < imgTotal.Rows() && newY >= 0 && newY < imgTotal.Cols() {
				imgTotal.SetUCharAt(newX, newY, p)
			}
		}
	}

	window := gocv.NewWindow("Belirtilen Açıya Göre Çevirme")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(imgTotal)
	window.WaitKey(0)

}
