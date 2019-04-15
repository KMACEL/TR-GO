package main

import (
	"gocv.io/x/gocv"
)

func main() {

	img := gocv.NewMat()

	// 1. Yöntem : İlk görüntüyü yüklerken
	//img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadGrayScale)

	// 2. Yöntem : Görüntüyü normal alıp değiştirme
	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)
	//gocv.CvtColor(img, &img, gocv.ColorBGRToGray)

	// 3. Yöntem : Manuel olarak döndürme
	for i := 0; i < img.Rows(); i++ {
		for j := 0; j < img.Cols(); j++ {

			p := img.GetIntAt(i, j)

			r := (p >> 16) & 0xff
			g := (p >> 8) & 0xff
			b := (p >> 0) & 0xFF

			//Gri-ton formülü
			grayValue := float32(r)*0.299 + float32(g)*0.587 + float32(b)*0.114

			r = int32(grayValue)
			g = int32(grayValue)
			b = int32(grayValue)

			p = (r << 16) | (g << 8) | b
			img.SetIntAt(i, j, p)

		}
	}

	window := gocv.NewWindow("Siyah Beyaz Yapma")
	window.IMShow(img)
	window.WaitKey(0)
}
