package main

import (
	"gocv.io/x/gocv"
)

// f(x;y)=255-g(x;y)

func main() {

	img := gocv.NewMat()
	img2 := gocv.NewMat()

	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)
	img.CopyTo(&img2)

	// Yöntem 1 :

	//gocv.BitwiseNot(img, &img)

	// Uzun Yöntem :
	for i := 0; i < img.Rows(); i++ {
		for j := 0; j < img.Cols(); j++ {

			// 1. Uzun yöntem
			//img2.SetIntAt(i, j, 255-img.GetIntAt(i, j))

			//2. Dahada uzun yötem
			p := img.GetIntAt(i, j)

			a := (p >> 24) & 0xff
			r := (p >> 16) & 0xff
			g := (p >> 8) & 0xff
			b := p & 0xff

			// 255 ten pikseldeki renkleri çıkartıyoruz
			a = 255 - a
			r = 255 - r
			g = 255 - g
			b = 255 - b

			//Yeni değerleri piksele giriyoruz
			p = (r << 24) | (r << 16) | (g << 8) | b

			img2.SetIntAt(i, j, p)
		}
	}

	window := gocv.NewWindow("Negatif ALma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(img2)
	window.WaitKey(0)

}
