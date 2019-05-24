package main

import (
	"fmt"
	"io"
	"os/exec"
	"strconv"

	"gocv.io/x/gocv"
)

const (
	frameX      = 960
	frameY      = 720
	frameSize   = frameX * frameY * 3
	minimumArea = 3000
)

func main() {
	ffmpeg := exec.Command("ffmpeg", "-i", "http://localhost:8089/movie", "-pix_fmt", "bgr24", "-s", strconv.Itoa(frameX)+"x"+strconv.Itoa(frameY), "-f", "rawvideo", "pipe:1")

	ffmpegOut, _ := ffmpeg.StdoutPipe()

	if err := ffmpeg.Start(); err != nil {
		panic(err)
	}

	window := gocv.NewWindow("Video Stream")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	for {
		buf := make([]byte, frameSize)
		if _, err := io.ReadFull(ffmpegOut, buf); err != nil {
			fmt.Println(err)
			continue
		}
		img, _ := gocv.NewMatFromBytes(frameY, frameX, gocv.MatTypeCV8UC3, buf)
		if img.Empty() {
			continue
		}

		window.IMShow(img)
		if window.WaitKey(1) == 27 {
			break
		}
	}
}
