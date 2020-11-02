package admin

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// MediaDir directory of medias
const MediaDir string = "assets/media"

type mediaList struct {
	Name      string
	Size      string
	UpdatedAt time.Time
}

// MediaController implement media crud operations
func MediaController(app *fiber.App) {
	app.Get("/medias", func(c *fiber.Ctx) error {
		cmd := exec.Command("ls", "-AgohcX", "--color=no", "--time-style=+%s", MediaDir)
		stdouterr, err := cmd.CombinedOutput()
		if err != nil {
			log.Error("MediaController: cmd.Run() failed with ", err)
			log.Error(string(stdouterr))
			return err
		}
		medias := parseLs(string(stdouterr))
		return c.Render("admin/media", fiber.Map{"Title": "Medias", "Navigation": Navigation, "Medias": medias, "MediaDir": MediaDir}, "layouts/main")
	})

	app.Post("/medias", func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		for _, file := range form.File["medias"] {
			log.Info(file.Filename, string(file.Size), file.Header["Content-Type"][0])
			err := c.SaveFile(file, fmt.Sprintf("%s/%s", MediaDir, file.Filename))

			if err != nil {
				return err
			}
		}
		return c.Redirect("/medias")
	})
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
