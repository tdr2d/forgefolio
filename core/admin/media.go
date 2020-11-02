package admin

import (
	"fmt"
	"os/exec"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// MediaDir directory of medias
const MediaDir string = "assets/media"

type mediaList struct {
	Dir       string
	Name      string
	Size      string
	UpdatedAt time.Time
}

// MediaController implement media crud operations
func MediaController(app *fiber.App) {
	app.Get("/medias", func(c *fiber.Ctx) error {
		cmd := exec.Command("ls", "-ARgohcX", "--color=no", "--time-style=full-iso", MediaDir)
		stdout, err := cmd.Output()
		if err != nil {
			log.Error("cmd.Run() failed with", err)
			return err
		}
		log.Info(string(stdout))

		return c.Render("admin/media", fiber.Map{"Title": "Medias", "Navigation": Navigation}, "layouts/main")
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

func lsParsing(stdout string) {

}
