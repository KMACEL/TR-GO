package main

import (
	"gocv.io/x/gocv"
)

// f(x;y)=255-g(x;y)

func main() {

	img := gocv.NewMat()
	imgTotal := gocv.NewMat()

	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)
	img.CopyTo(&imgTotal)

	// Yöntem 1 :
	//gocv.BitwiseNot(img, &img)

	// Uzun Yöntem :
	for x := 0; x < img.Rows(); x++ {
		for y := 0; y < img.Cols(); y++ {

			// 1. Uzun yöntem
			//imgTotal.SetIntAt(x, y, 255-img.GetIntAt(x, y))

			//2. Dahada uzun yötem
			p := img.GetIntAt(x, y)

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
			pTotal := (r << 24) | (r << 16) | (g << 8) | b

			imgTotal.SetIntAt(x, y, pTotal)
		}
	}

	window := gocv.NewWindow("Negatif ALma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(imgTotal)
	window.WaitKey(0)

}
