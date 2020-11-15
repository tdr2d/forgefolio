package admin

import fiber "github.com/gofiber/fiber/v2"

func HomeController(app fiber.Router) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("views/admin/home", fiber.Map{"Title": "Home", "Constants": Constants}, Layout)
	})
}
