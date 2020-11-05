package admin

import (
	"core/utils"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type mediaList struct {
	Name string
	Ext  string
}

// MediaController implement media crud operations
func MediaController(app *fiber.App) {
	app.Get("/medias", func(c *fiber.Ctx) error {
		files, err := ioutil.ReadDir(MediaDir)
		if err != nil {
			log.Error(err)
		}
		medias := make([]mediaList, len(files))
		for i, file := range files {
			_, name, ext := utils.GetDirNameExtension(file.Name())
			medias[i].Name = name
			medias[i].Ext = ext
		}
		data := fiber.Map{
			"Title":       "Medias",
			"Navigation":  Navigation,
			"Medias":      medias,
			"MediaDir":    MediaDir,
			"ThumnailDir": MediaThumbnailDir,
			"BaseUrl":     c.BaseURL(),
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
			filepath := fmt.Sprintf("%s/%s", MediaDir, file.Filename)
			err := c.SaveFile(file, filepath)
			if err != nil {
				log.Error(err)
				return err
			}
			if _, err := utils.Thumbnail(filepath, 250, 0, MediaThumbnailDir); err != nil {
				log.Error(err)
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
			filePath := fmt.Sprintf("%s/%s", MediaDir, media)
			_, name, _ := utils.GetDirNameExtension(filePath)
			thumbnailPath := fmt.Sprintf("%s/%s.jpg", MediaThumbnailDir, name)
			if err := os.Remove(filePath); err != nil {
				log.Error(err)
			}
			if err := os.Remove(thumbnailPath); err != nil {
				log.Error(err)
			}
		}
		return c.SendStatus(204)
	})
}
