package admin

import fiber "github.com/gofiber/fiber/v2"

type Struct Post {
	id int
	
}


// BlogController implement blog crud operations
func BlogController(app *fiber.App) {
	// app.Get("/blog/:id", func(c *fiber.Ctx) error {
	// })

	app.Post("/blog", func(c *fiber.Ctx) error {

	})
}
