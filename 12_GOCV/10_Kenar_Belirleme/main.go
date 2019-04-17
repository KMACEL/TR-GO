package main

import (
	"math"

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
	//gocv.Canny(img, &img2, 200, 200)

	// Uzun Yöntem :
	for x := 0; x < img.Rows(); x++ {
		for y := 0; y < img.Cols(); y++ {

			p := img.GetIntAt(x-1, y-1)
			c1 := (p >> 16) & 0xff

			p = img.GetIntAt(x, y-1)
			c2 := (p >> 16) & 0xff

			p = img.GetIntAt(x+1, y-1)
			c3 := (p >> 16) & 0xff

			p = img.GetIntAt(x-1, y)
			c4 := (p >> 16) & 0xff

			//p = img.GetUCharAt(x, y)
			//c5 := (p >> 16) & 0xff

			p = img.GetIntAt(x+1, y)
			c6 := (p >> 16) & 0xff

			p = img.GetIntAt(x-1, y+1)
			c7 := (p >> 16) & 0xff

			p = img.GetIntAt(x, y+1)
			c8 := (p >> 16) & 0xff

			p = img.GetIntAt(x+1, y+1)
			c9 := (p >> 16) & 0xff

			pValue := math.Abs(float64(c1+2*c2+c3)-float64(c7+2*c8+c9)) + math.Abs(float64(c3+2*c6+c9)-float64(c1+2*c4+c7))

			if pValue > 255 {
				pValue = 255
				img2.SetIntAt(x, y, int32(pValue))
			}
		}
	}

	window := gocv.NewWindow("Kenar Belirleme")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)

	window.IMShow(img2)
	window.WaitKey(0)

}
