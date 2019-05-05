package main

import (
	"gocv.io/x/gocv"
)

/*
Formül 1 : grayValue=(0.299*r)+(0.587*g)+(0.114*b)
Formül 2 : grayValue=(0.21*r)+(0.71*g)+(0.071xb)
Formül 3 : grayValue = (r + g + b) / 3
*/
func main() {

	img := gocv.NewMat()

	// 1. Yöntem : İlk görüntüyü yüklerken
	//img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadGrayScale)

	// 2. Yöntem : Görüntüyü normal alıp değiştirme
	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)
	//gocv.CvtColor(img, &img, gocv.ColorBGRToGray)

	// 3. Yöntem : Manuel olarak döndürme
	for x := 0; x < img.Rows(); x++ {
		for y := 0; y < img.Cols(); y++ {

			p := img.GetIntAt(x, y)

			//a := (p >> 24) & 0xff
			r := (p >> 16) & 0xff
			g := (p >> 8) & 0xff
			b := (p >> 0) & 0xff

			//Gri-ton formülü
			grayValue := float32(r)*0.299 + float32(g)*0.587 + float32(b)*0.114

			r = int32(grayValue)
			g = int32(grayValue)
			b = int32(grayValue)

			pTotal := (r << 16) | (g << 8) | b
			img.SetIntAt(x, y, pTotal)

		}
	}

	window := gocv.NewWindow("Siyah Beyaz Yapma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(img)
	window.WaitKey(0)
}
