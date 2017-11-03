package main

import (
	"fmt"
	"github.com/desintegration/imaging"
	"image"
	"image/color"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://elelur.com/data_images/mammals/lion/lion-05.jpg")
	if err != nil {
		print("you need try hard")
	}

	//open a file for writing
	file, err := os.Create("/tmp/test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	src, err := imaging.Open("/tmp/test.jpg")
	if err != nil {
		log.Fatalf("Open failed: %v", err)
	}
	// Crop the original image to 350x350px size using the center anchor.
	src = imaging.CropAnchor(src, 1500, 1900, imaging.Center)

	// Resize the cropped image to width = 256px preserving the aspect ratio.
	src = imaging.Resize(src, 256, 0, imaging.Lanczos)

	// Create a blurred version of the image.
	img1 := imaging.Blur(src, 2)

	// Create an inverted version of the image.
	img3 := imaging.Invert(src)

	img2 := imaging.Grayscale(src)
	img2 = imaging.AdjustContrast(img2, 20)
	img2 = imaging.Sharpen(img2, 2)

	// Create an embossed version of the image using a convolution filter.
	img4 := imaging.Convolve3x3(
		src,
		[9]float64{
			-1, -1, 0,
			-1, 1, 1,
			0, 1, 1,
		},
		nil,
	)

	dst := imaging.New(512, 512, color.NRGBA{0, 0, 0, 0})

	dst = imaging.Paste(dst, img1, image.Pt(0, 0))
	dst = imaging.Paste(dst, img2, image.Pt(0, 256))
	dst = imaging.Paste(dst, img3, image.Pt(256, 0))
	dst = imaging.Paste(dst, img4, image.Pt(256, 256))

	// Save the resulting image using JPEG format.
	err = imaging.Save(dst, "/tmp/test.jpg")

	file.Close()
	fmt.Println("Success!")
}
