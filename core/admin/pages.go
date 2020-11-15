package admin

import (
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
)

// PluginController controller for the plugin screen and plugin assets
func PageController(app *fiber.App) {
	app.Get("/admin/pages", func(c *fiber.Ctx) error {
		data := fiber.Map{
			"Title":     "Pages",
			"Constants": Constants,
			"Pages":     nil,
		}
		return c.Render("views/admin/pages", data, Layout)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		currentTheme := "blog-theme"
		homePage := "pages/post"
		return c.Render(fmt.Sprintf("themes/%s/%s", currentTheme, homePage), nil)
	})
}
