package admin

import (
	"fmt"
	"os"
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
		medias := listMedia(MediaDir)
		data := fiber.Map{
			"Title":      "Medias",
			"Navigation": Navigation,
			"Medias":     medias,
			"MediaDir":   MediaDir,
			"BaseUrl":    c.BaseURL(),
		}
		return c.Render("admin/media", data, "layouts/main")
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

	app.Patch("/medias/:name", func(c *fiber.Ctx) error {
		media := c.Params("name")
		newName := c.FormValue("new_name")
		if err := os.Rename(fmt.Sprintf("%s/%s", MediaDir, media), fmt.Sprintf("%s/%s", MediaDir, newName)); err != nil {
			log.Error(err)
		}
		return c.SendStatus(204)
	})

	app.Delete("/medias", func(c *fiber.Ctx) error {
		medias := strings.Split(c.FormValue("medias"), ",")
		log.Info("DELETE /medias ", medias)
		for _, media := range medias {
			if err := os.Remove(fmt.Sprintf("%s/%s", MediaDir, media)); err != nil {
				log.Error(err)
			}
		}
		return c.SendStatus(204)
	})
}

func listMedia(dir string) []mediaList {
	cmd := exec.Command("ls", "-AgohcX", "--color=no", "--time-style=+%s", MediaDir)
	stdouterr, err := cmd.CombinedOutput()
	if err != nil {
		log.Error("MediaController: cmd.Run() failed with ", err)
		log.Error(string(stdouterr))
		return nil
	}

	var data []mediaList = make([]mediaList, 0)
	lines := strings.Split(string(stdouterr), "\n")
	for _, line := range lines[1:] {
		if line != "" && line[0] != 'd' {
			tokens := strings.Split(line, " ")
			dateInt, err := strconv.ParseInt(tokens[3], 10, 0)
			if err != nil {
				log.Error(err)
			}
			mediaListItem := mediaList{Size: tokens[2], UpdatedAt: time.Unix(dateInt, 0), Name: tokens[4]}
			data = append(data, mediaListItem)
		}
	}
	return data
}

// build thumbnail of check if
// func getImageThumbnail(path string, sizeX int, sizeY int) string {
// TODO thumbnail
// }
