package admin

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// PluginController controller for the plugin screen and plugin assets
func PluginController(app *fiber.App) {
	app.Get("/plugins/:pluginName/*", func(c *fiber.Ctx) error {
		path := fmt.Sprintf("%s/%s/%s", PluginDir, c.Params("pluginName"), c.Params("*1"))
		log.Info(path)
		return c.SendFile(path)
	})
}
