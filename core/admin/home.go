package admin

import fiber "github.com/gofiber/fiber/v2"

func HomeController(app fiber.Router) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("admin/home", fiber.Map{"Title": "Home", "Constants": Constants}, "layouts/main")
	})
}
