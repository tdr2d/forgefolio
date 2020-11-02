package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var input2 string = `total 3.9M
-rwxrwxrwx 1 3.8M 1604336772 jachan-devol-xY_6ZENqcfo-unsplash.jpg
-rwxrwxrwx 1 168K 1604336772 kubeArchi.png
`

type mediaList struct {
	Name      string
	Size      string
	UpdatedAt time.Time
}

func parseLs(stdout string) []mediaList {
	var data []mediaList = make([]mediaList, 0)
	lines := strings.Split(stdout, "\n")
	// fmt.Println(lines)
	for _, line := range lines[1:] {
		// fmt.Printf("Line: %s\n", line)
		if line != "" && line[0] != 'd' {
			tokens := strings.Split(line, " ")
			// fmt.Println(tokens)
			dateInt, err := strconv.ParseInt(tokens[3], 10, 0)
			if err != nil {
				fmt.Println(err)
			}
			mediaListItem := mediaList{Size: tokens[2], UpdatedAt: time.Unix(dateInt, 0), Name: tokens[4]}
			data = append(data, mediaListItem)
		}
	}
	return data
}

func main() {
	fmt.Println(parseLs(input2))
}
