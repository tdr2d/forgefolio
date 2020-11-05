package utils

import (
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
	log "github.com/sirupsen/logrus"
)

// Thmbnail create thumbnail if not exists
func Thumbnail(path string, width uint, height uint, outputDir string) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Error(err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Error(path, err)
	}
	// fmt.Println("Width:", image.Width, "Height:", image.Height)

	name, ext := getNameAndExtension(path)
	out, err := os.Create("test_resized.jpg")
	thumbnail := resize.Resize(width, height, image, resize.NearestNeighbor)
	err = jpeg.Encode(someWriter, newImage, nil)
	if err != nil {
		log.Error(path, err)
	}
}

func getNameAndExtension(path string) (string, string) {
	ext := filepath.Ext(path)
	return strings.TrimSuffix(path, ext), ext
}
