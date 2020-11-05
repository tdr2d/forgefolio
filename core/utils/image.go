package utils

import (
	"fmt"
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
		thumbnailImage := resize.Resize(width, height, image, resize.NearestNeighbor)

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

// GetDirNameExtension splits the path in Dir, name and extension
func GetDirNameExtension(path string) (string, string, string) {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return filepath.Dir(path), strings.TrimSuffix(base, ext), ext
}
