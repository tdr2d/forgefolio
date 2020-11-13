package utils

import (
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/nfnt/resize"
	log "github.com/sirupsen/logrus"
)

// Thumbnail create jpeg thumbnail in outpuDir if not exists
func Thumbnail(fromImagePath string, width uint, height uint, outputDir string) (string, error) {
	_, name, _ := GetDirNameExtension(fromImagePath)
	thumbnailPath := fmt.Sprintf("%s/%s.jpg", outputDir, name)

	if _, err := os.Stat(thumbnailPath); os.IsNotExist(err) {

		// Resize image
		file, err := os.Open(fromImagePath)
		defer file.Close()
		if err != nil {
			log.Error(err)
			return fromImagePath, err
		}
		image, _, err := image.Decode(file)
		if err != nil {
			log.Error(fromImagePath, err)
			return fromImagePath, err
		}
		thumbnailImage := resize.Resize(width, height, image, resize.Lanczos3)

		// Write image
		outputFile, err := os.Create(thumbnailPath)
		err = jpeg.Encode(outputFile, thumbnailImage, nil)
		if err != nil {
			log.Error(fromImagePath, err)
			return fromImagePath, err
		}
	}
	return thumbnailPath, nil
}
