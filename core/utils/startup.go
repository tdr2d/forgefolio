package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func CheckDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Error(err)
		}
	}
}
