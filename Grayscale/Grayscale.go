// Bryce Koch
package Grayscale

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"os"
)

func Gray(input string) {
	// Requesting the image file name.
	fmt.Println("Enter the name of the .jpg image file you are using: ")

	// Storing the input from the console into input.
	fmt.Scanln(input)

	fmt.Println("Opening " + input)

	// Open the original image
	reader, err := os.Open(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer reader.Close()

	// Decode the image
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Get image bounds
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Create a new grayscale image
	grayImg := image.NewGray(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Get the color at pixel (x, y)
			oldColor := img.At(x, y)
			r, g, b, _ := oldColor.RGBA()

			// Convert to gray using the formula
			gray := uint8((0.3*float64(r) + 0.59*float64(g) + 0.11*float64(b)) / 256.0)

			// Set the gray color
			grayColor := color.Gray{Y: gray}
			grayImg.Set(x, y, grayColor)
		}
	}

	// Save the grayscale image
	grayFile, err := os.Create("gray_image.png")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer grayFile.Close()
	png.Encode(grayFile, grayImg)

	fmt.Println("Grayscale image saved.")
}
