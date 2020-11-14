package admin

import (
	fiber "github.com/gofiber/fiber/v2"
)

// PluginController controller for the plugin screen and plugin assets
func PageController(app fiber.Router) {
	app.Get("/pages", func(c *fiber.Ctx) error {
		data := fiber.Map{
			"Title":     "Pages",
			"Constants": Constants,
			"Pages":     nil,
		}
		return c.Render("admin/pages", data, "layouts/main")
	})
}
