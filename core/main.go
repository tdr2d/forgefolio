package main

import (
	"core/admin"
	"core/utils"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/jet"
	log "github.com/sirupsen/logrus"
)

var port string = ":8080"

func init() {
	utils.CheckDir(admin.Constants.MediaDir)
	utils.CheckDir(admin.Constants.MediaThumbnailDir)
	utils.CheckDir(admin.Constants.BlogDataDir)
	// log.SetReportCaller(true)
}

func main() {
	engine := jet.New("./views", ".jet")
	engine.Reload(true)
	app := fiber.New(fiber.Config{Views: engine, BodyLimit: admin.Constants.BodyLimit})
	app.Use(logger.New())
	app.Use(recover.New())

	adminApiGroup := app.Group("/admin")
	adminApiGroup.Get("/", func(c *fiber.Ctx) error {
		return c.Render("admin/home", fiber.Map{"Title": "Home", "Constants": admin.Constants}, "layouts/main")
	})
	admin.MediaController(adminApiGroup)
	admin.BlogController(adminApiGroup)
	adminApiGroup.Get("/settings", func(c *fiber.Ctx) error {
		return c.Render("admin/settings", fiber.Map{"Title": "Settings", "Constants": admin.Constants}, "layouts/main")
	})
	app.Static("/assets/", "./assets")
	// data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	// log.Info(string(data))
	log.Fatal(app.Listen(port))
}
