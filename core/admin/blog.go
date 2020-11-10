package admin

import fiber "github.com/gofiber/fiber/v2"

// BlogPost is the structure of the post
type BlogPost struct {
	id int
}

// BlogController implement blog crud operations
func BlogController(app *fiber.App) {
	// app.Get("/blog/:id", func(c *fiber.Ctx) error {
	// })
	app.Get("/blog-posts", func(c *fiber.Ctx) error {
		return c.Render("admin/blog-posts", fiber.Map{"Title": "Blog Posts", "Navigation": Navigation}, "layouts/main")
	})

	// app.Post("/blog-posts", func(c *fiber.Ctx) error {
	// 	return nil
	// })
}
