package main

import (
	"math"

	"gocv.io/x/gocv"
)

/*
	----------
	|p1|p2|p3|
	----------
	|p4|p5|p6|
	----------
	|p7|p8|p9|
	----------
*/
// Formül : |G| =|((p1+2*p2+p3)-(p7+2*p8+p9))| + |((p3+2*p6+p9)-(p1+2*p4+p7))|

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
			p1 := (p >> 16) & 0xff

			p = img.GetIntAt(x, y-1)
			p2 := (p >> 16) & 0xff

			p = img.GetIntAt(x+1, y-1)
			p3 := (p >> 16) & 0xff

			p = img.GetIntAt(x-1, y)
			p4 := (p >> 16) & 0xff

			//p = img.GetUCharAt(x, y)
			//p5 := (p >> 16) & 0xff

			p = img.GetIntAt(x+1, y)
			p6 := (p >> 16) & 0xff

			p = img.GetIntAt(x-1, y+1)
			p7 := (p >> 16) & 0xff

			p = img.GetIntAt(x, y+1)
			p8 := (p >> 16) & 0xff

			p = img.GetIntAt(x+1, y+1)
			p9 := (p >> 16) & 0xff

			pValue := math.Abs(float64((p1+2*p2+p3)-(p7+2*p8+p9))) + math.Abs(float64((p3+2*p6+p9)-(p1+2*p4+p7)))

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
