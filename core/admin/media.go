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
func MediaController(app fiber.Router) {
	app.Get("/medias", func(c *fiber.Ctx) error {
		files, err := ioutil.ReadDir(Constants.MediaDir)
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
			"Title":     "Medias",
			"Constants": Constants,
			"Medias":    medias,
		}
		return c.Render("views/admin/media", data, Layout)
	})

	app.Post("/medias", func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			panic(err)
		}
		for _, file := range form.File["medias"] {
			log.Info(file.Filename, string(file.Size), file.Header["Content-Type"][0])
			filepath := fmt.Sprintf("%s/%s", Constants.MediaDir, file.Filename)
			err := c.SaveFile(file, filepath)
			if err != nil {
				log.Error(err)
				return err
			}
			if _, err := utils.Thumbnail(filepath, 250, 0, Constants.MediaThumbnailDir); err != nil {
				log.Error(err)
			}
		}
		if c.Query("redirect") != "" {
			return c.Redirect("/admin/medias")
		}
		return c.SendStatus(201)
	})

	app.Patch("/medias/:name", func(c *fiber.Ctx) error {
		media := c.Params("name")
		newName := c.FormValue("new_name")
		if err := os.Rename(fmt.Sprintf("%s/%s", Constants.MediaDir, media), fmt.Sprintf("%s/%s", Constants.MediaDir, newName)); err != nil {
			log.Error(err)
		}
		_, name, _ := utils.GetDirNameExtension(media)
		_, newThumbName, _ := utils.GetDirNameExtension(newName)
		if err := os.Rename(fmt.Sprintf("%s/%s.jpg", Constants.MediaThumbnailDir, name), fmt.Sprintf("%s/%s.jpg", Constants.MediaThumbnailDir, newThumbName)); err != nil {
			log.Error(err)
		}
		return c.SendStatus(204)
	})

	app.Delete("/medias", func(c *fiber.Ctx) error {
		medias := strings.Split(c.FormValue("medias"), ",")
		log.Info("DELETE /medias ", medias)
		for _, media := range medias {
			filePath := fmt.Sprintf("%s/%s", Constants.MediaDir, media)
			_, name, _ := utils.GetDirNameExtension(filePath)
			thumbnailPath := fmt.Sprintf("%s/%s.jpg", Constants.MediaThumbnailDir, name)
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
