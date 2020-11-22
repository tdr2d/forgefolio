package utils

import (
	"bytes"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

const (
	empty = ""
	tab   = "\t"
)

func PrettyPrintJson(data interface{}) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)
	encoder.Encode(data)
	log.Info(buffer.String())
}
