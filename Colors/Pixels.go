// Package decleration
package Pixels

// Required packages being imported.
import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"
)

func Pixels(input string) {
	// Requesting the image file name.
	fmt.Println("Enter the name of the .jpg image file you are using: ")

	// Storing the input from the console into input.
	fmt.Scanln(input)

	fmt.Println("Opening " + input)

	// Opening the image file.
	reader, err := os.Open(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer reader.Close()

	// Breaking down the pixel details.
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Prints out the pixel details.
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			color := img.At(x, y)
			r, g, b, _ := color.RGBA()
			fmt.Printf("Pixel at (%d, %d) - R: %d, G: %d, B: %d\n", x, y, r>>8, g>>8, b>>8)
		}
	}
}
