package hash

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/corona10/goimagehash"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// GeneratePHash downloads image from specified url and generates the perception hash for the image
func GeneratePHash(url string, fileType string) (uint64, error) {
	tokens := strings.Split(url, "/")
	filePath := "/tmp/" + tokens[len(tokens)-1]
	err := downloadFile(filePath, url)
	if err != nil {
		fmt.Println("error downloading file: ", err)
		return 0, err
	}
	defer os.Remove(filePath)

	file1, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}

	var img1 image.Image
	if fileType == "jpg" || fileType == "jpeg" {
		img1, err = jpeg.Decode(file1)
	}
	if fileType == "png" {
		img1, err = png.Decode(file1)
	}
	if err != nil {
		return 0, err
	}

	hash1, err := goimagehash.PerceptionHash(img1)

	if err != nil {
		return 0, err
	}
	return hash1.GetHash(), nil
}
