package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// CheckDir creates the directory if it does't exist
func CheckDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Error(err)
		}
	}
}

// list plugin dir
// read plugin.json

// create endpoints
