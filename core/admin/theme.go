package admin

import fiber "github.com/gofiber/fiber/v2"

type Theme struct {
	name    string
	version string
	author  string
}

// PluginController controller for the plugin screen and plugin assets
func ThemeController(app fiber.Router) {
	app.Get("/theme", func(c *fiber.Ctx) error {
		data := fiber.Map{
			"Title":     "Theme",
			"Constants": Constants,
			"Themes":    nil,
		}
		return c.Render("admin/theme", data, "layouts/main")
	})
}
