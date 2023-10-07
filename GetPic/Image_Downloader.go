package Image_Downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Download(downloaded_Image string) {
	// Variable for storing the URL of the image.
	var input string

	// Asking for the image URL
	fmt.Println("Input the URL of the .jpg image file you want downloaded: ")

	// Storing the console input
	fmt.Scanln(input)

	// URL of the image you want to download
	//imageUrl := "https://images.pexels.com/photos/17929271/pexels-photo-17929271/free-photo-of-woman-standing-on-vineyard.jpeg"
	imageUrl := input

	fmt.Println("Locating the file specified...")

	// Create an HTTP GET request
	response, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", response.StatusCode)
		return
	}

	fmt.Println("Image found")

	fmt.Println("What do you want the downloaded image file name to be?:")

	fmt.Scanln(downloaded_Image)

	// Create a new file to save the image
	outputFile, err := os.Create(downloaded_Image)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer outputFile.Close()

	// Copy the HTTP response body to the file
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		fmt.Println("Error saving the image:", err)
		return
	}

	fmt.Println("Image downloaded and saved as " + downloaded_Image)
}
