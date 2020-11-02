package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet"
	log "github.com/sirupsen/logrus"
)

func init() {
	// configure logrus
}

func main() {
	engine := jet.New("./views", ".jet")
	engine.Reload(true)

	app := fiber.New(fiber.Config{Views: engine})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{"Title": "Hello, World!"}, "layouts/main")
	})

	app.Get("/admin", func(c *fiber.Ctx) error {
		return c.Render("admin", fiber.Map{"Title": "Hello, World!"})
	})

	app.Static("/assets/", "./assets")
	log.Fatal(app.Listen(":8080"))
}
