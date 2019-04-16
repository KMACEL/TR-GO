package main

import (
	"gocv.io/x/gocv"
)

// f(x;y)=255-g(x;y)

func main() {

	img := gocv.NewMat()
	img = gocv.IMRead("../MERT_KUBRA_ERDEM.jpg", gocv.IMReadAnyColor)

	// Yöntem 1 :

	//gocv.BitwiseNot(img, &img)

	// Uzun Yöntem :
	/*
		neden img.Cols()*3 yaptığımız konusunda;
			https://stackoverflow.com/questions/40848254/opencv-videowriter-size-issue
			   May be I'm too late. But this solution works for me. Resize the Mat by using cv::resize(finalOutputRGB,finalOutputRGB, cv::size(width, height*3))**Use height of the image*3 for the height. ** That's it. It solved my problem. Hope that would help.
	*/
	for i := 0; i < img.Rows(); i++ {
		for j := 0; j < img.Cols()*3; j++ {

			// 1. Uzun yöntem
			//	img.SetUCharAt(i, j, 255-img.GetUCharAt(i, j))

			//2. Dahada uzun yötem
			p := img.GetUCharAt(i, j)

			r := (p >> 16) & 0xff
			g := (p >> 8) & 0xff
			b := p & 0xff

			// 255 ten pikseldeki renkleri çıkartıyoruz
			r = 255 - r
			g = 255 - g
			b = 255 - b

			//Yeni değerleri piksele giriyoruz
			p = (r << 16) | (g << 8) | b
			img.SetUCharAt(i, j, p)
		}
	}

	window := gocv.NewWindow("Negatif ALma")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(img)
	window.WaitKey(0)

}
