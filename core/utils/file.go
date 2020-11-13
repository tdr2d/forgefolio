package utils

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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

// GetDirNameExtension splits the path in Dir, name and extension
func GetDirNameExtension(path string) (string, string, string) {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return filepath.Dir(path), strings.TrimSuffix(base, ext), ext
}

// Serialize and save struct in file
func PersistStruct(data interface{}, path string) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return err
	}
	return ioutil.WriteFile(path, buf.Bytes(), 0644)
}

// Read file and deserialize data into struct
func ReadStruct(output interface{}, path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(file)
	dec := gob.NewDecoder(reader)
	return dec.Decode(output)
}
