package main

import (
	"core/admin"
	"core/utils"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet"
	log "github.com/sirupsen/logrus"
)

var port string = ":8080"

func init() {
	utils.CheckDir(admin.MediaDir)
	utils.CheckDir(admin.MediaThumbnailDir)
	utils.CheckDir(admin.PluginDir)
	// log.SetReportCaller(true)
}

func main() {
	engine := jet.New("./views", ".jet")
	engine.Reload(true)

	app := fiber.New(fiber.Config{Views: engine, BodyLimit: admin.BodyLimit})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("admin/home", fiber.Map{"Title": "Home", "Navigation": admin.Navigation}, "layouts/main")
	})
	admin.MediaController(app)
	admin.PluginController(app)
	app.Get("/settings", func(c *fiber.Ctx) error {
		return c.Render("admin/settings", fiber.Map{"Title": "Settings", "Navigation": admin.Navigation}, "layouts/main")
	})
	app.Get("/blog-posts", func(c *fiber.Ctx) error {
		return c.Render("admin/blog-posts", fiber.Map{"PluginDir": admin.PluginDir, "Title": "Blog Posts", "Navigation": admin.Navigation}, "layouts/main")
	})

	app.Static("/assets/", "./assets")
	// app.Use(recover.New())

	// data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	// log.Info(string(data))
	log.Fatal(app.Listen(port))
}
