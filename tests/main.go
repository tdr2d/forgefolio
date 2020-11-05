package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

const mediaDir = "../core/assets/media"
const mediaThumbnailDir = "../core/assets/media"
const mediaFullName = "kubeArchi.png"

func main() {
	dir, name, ext := getDirNameExtension(mediaDir + "/" + mediaFullName)

	thumbnailName := fmt.Sprintf("%s.jpg", name)

	fmt.Println(dir, name, ext, thumbnailName)
}

// Split the path in Dir, name and extension
func getDirNameExtension(path string) (string, string, string) {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return filepath.Dir(path), strings.TrimSuffix(base, ext), ext
}
