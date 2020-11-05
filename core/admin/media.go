package admin

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// MediaDir directory of medias
const MediaDir string = "assets/media"
const MediaThumbnailDir string = "assets/media/thumbnail"

// MediaController implement media crud operations
func MediaController(app *fiber.App) {
	app.Get("/medias", func(c *fiber.Ctx) error {
		files, err := ioutil.ReadDir(MediaDir)
		if err != nil {
			log.Error(err)
		}
		data := fiber.Map{
			"Title":      "Medias",
			"Navigation": Navigation,
			"Files":      files,
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

// build thumbnail of check if
// func getImageThumbnail(path string, sizeX int, sizeY int) string {
// TODO thumbnail
// }
