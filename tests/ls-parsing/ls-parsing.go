package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

var stdout string = `.:
total 3.9M
drwxrwxrwx 1 4.0K 2020-11-02 20:38:25.571559100 +0100 nested
-rwxrwxrwx 1 3.8M 2020-11-02 18:06:12.643041200 +0100 jachan-devol-xY_6ZENqcfo-unsplash.jpg
-rwxrwxrwx 1 168K 2020-11-02 18:06:12.648040900 +0100 kubeArchi.png

./nested:
total 0
drwxrwxrwx 1 4.0K 2020-11-02 20:38:25.571559100 +0100 nested_2

./nested/nested_2:
total 0`

type mediaList struct {
	Dir       string
	Name      string
	Size      string
	UpdatedAt time.Time
}

func testParsing() {
	directoriesRegex := regexp.MustCompile(`\n\n`)
	directories := directoriesRegex.Split(stdout, -1)
	rDir := regexp.MustCompile(`[./]*(.+):`)
	var data []mediaList = make([]mediaList, 0)

	fmt.Printf("Found %d directories\n\n", len(directories))
	for _, dir := range directories {
		lines := strings.Split(dir, "\n")
		mediaListItem := mediaList{Dir: rDir.FindStringSubmatch(lines[0])[1]}

		for _, l := range lines[2:] {
			if l[0] != 'd' {
				tokens := strings.Split(l, " ")
				date := fmt.Sprintf("%s %s %s", tokens[3], tokens[4], tokens[5])
				parsedData, err := time.Parse(time.RFC1123Z, date)
				if err != nil {
					fmt.Printf("ERROR: %s\n", err)
				}
				mediaListItem.Size = tokens[2]
				mediaListItem.UpdatedAt = parsedData
				mediaListItem.Name = tokens[6]
			}
		}
		data = append(data, mediaListItem)
		fmt.Println(mediaListItem)
	}
}

func main() {
	testParsing()
}
