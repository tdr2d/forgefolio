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
var bodyLimit int = 4 * 1024 * 1024

func init() {
	utils.CheckDir(admin.Constants.MediaDir)
	utils.CheckDir(admin.Constants.MediaThumbnailDir)
	utils.CheckDir(admin.DataDir.Blog)
	utils.CheckDir(admin.DataDir.Page)
	utils.CheckDir(admin.DataDir.Themes)
	utils.CheckDir(admin.DataDir.ThemeData)
	log.SetReportCaller(true)
}

func main() {
	engine := jet.New("./", ".jet")
	engine.Reload(true)
	app := fiber.New(fiber.Config{Views: engine, BodyLimit: bodyLimit})
	app.Use(logger.New())
	app.Use(recover.New())

	adminApiGroup := app.Group("/admin")
	admin.HomeController(adminApiGroup)
	admin.MediaController(adminApiGroup)
	admin.BlogController(adminApiGroup)
	admin.ThemeController(adminApiGroup)
	admin.SettingsController(adminApiGroup)
	admin.PageController(app)

	app.Static("/assets/", "./assets")
	// data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	// log.Info(string(data))
	log.Fatal(app.Listen(port))
}
