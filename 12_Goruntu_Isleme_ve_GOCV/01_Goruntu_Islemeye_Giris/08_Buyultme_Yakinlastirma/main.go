package main

import (
	"gocv.io/x/gocv"
)

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
	img3 = gocv.GetRotationMatrix2D(image.Point{0, 0}, 0.0, 2)

	gocv.WarpAffine(img, &img2, img3, image.Point{X: 0, Y: 0})*/

	// Uzun Yöntem :
	for i := 0; i < img.Rows(); i++ {
		for j := 0; j < img.Cols(); j++ {

			p := img.GetUCharAt(i, j)

			if i*2 > 0 && i*2 < img.Rows() && j*2 >= 0 && j*2 < img.Cols() {
				img2.SetUCharAt(i*2, j*2, p)
			}

			if i*2+1 > 0 && i*2+1 < img.Rows() && j*2 >= 0 && j*2 < img.Cols() {
				img2.SetUCharAt(i*2+1, j*2, p)
			}

			if i*2 > 0 && i*2 < img.Rows() && j*2+1 >= 0 && j*2+1 < img.Cols() {
				img2.SetUCharAt(i*2, j*2+1, p)
			}

			if i*2+1 > 0 && i*2+1 < img.Rows() && j*2+1 >= 0 && j*2+1 < img.Cols() {
				img2.SetUCharAt(i*2+1, j*2+1, p)
			}

		}
	}

	window := gocv.NewWindow("Büyültme-Yakınlaştırma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(img2)
	window.WaitKey(0)

}
