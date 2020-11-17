package admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type Theme struct {
	Error   error
	Path    string
	Name    string `json:"name"`
	Version string `json:"version"`
	Author  string `json:"author"`
}

func openThemeConfig(baseName string, theme *Theme) error {
	themeConfPath := fmt.Sprintf("%s/%s/%s", DataDir.Theme, baseName, ThemeIndexConfigFile)
	themeConfig, err := ioutil.ReadFile(themeConfPath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(themeConfig, theme); err != nil {
		return err
	}
	return nil
}

// PluginController controller for the plugin screen and plugin assets
func ThemeController(app fiber.Router) {
	app.Get("/theme", func(c *fiber.Ctx) error {
		dirs, _ := ioutil.ReadDir(DataDir.Theme)
		themes := make([]Theme, len(dirs))
		for index, dir := range dirs {
			if err := openThemeConfig(dir.Name(), &themes[index]); err != nil {
				log.Error(err)
				themes[index].Error = err
			} else {
				themes[index].Path = dir.Name()
			}
		}
		log.Info(themes)
		data := fiber.Map{
			"Title":     "Theme",
			"Constants": Constants,
			"Themes":    themes,
		}
		return c.Render("views/admin/theme", data, Layout)
	})

	app.Post("/theme", func(c *fiber.Ctx) error {
		theme := new(Theme)
		if err := openThemeConfig(c.FormValue("theme"), theme); err != nil {
			log.Error(err)
		}
		log.Info(theme)
		return c.Redirect(Constants.BaseUrl + "/theme")
	})
}
