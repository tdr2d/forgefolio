package admin

import fiber "github.com/gofiber/fiber/v2"

func SettingsController(app fiber.Router) {
	app.Get("/settings", func(c *fiber.Ctx) error {
		data := fiber.Map{
			"Title":     "Settings",
			"Constants": Constants,
			"Settings":  nil,
		}
		return c.Render("views/admin/settings", data, Layout)
	})
}
