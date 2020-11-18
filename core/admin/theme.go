package admin

import (
	"core/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type Theme struct {
	Path     string
	Name     string `json:"name"`
	Version  string `json:"version"`
	Author   string `json:"author"`
	Styles   string `json:"styles"`
	PostPage string `json:"postpage"`
}

const currentThemeKey = "current_theme"

func openThemeConfig(basePath string, theme *Theme) error {
	themeConfPath := fmt.Sprintf("%s/%s/%s", DataDir.Themes, basePath, ThemeIndexConfigFile)
	themeConfig, err := ioutil.ReadFile(themeConfPath)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(themeConfig, theme); err != nil {
		return err
	}
	return nil
}

func GetCurrentTheme() *Theme {
	currentTheme := new(Theme)
	utils.ReadStruct(currentTheme, fmt.Sprintf("%s/%s", DataDir.ThemeData, currentThemeKey))
	return currentTheme
}

// PluginController controller for the plugin screen and plugin assets
func ThemeController(app fiber.Router) {
	app.Get("/theme", func(c *fiber.Ctx) error {
		dirs, _ := ioutil.ReadDir(DataDir.Themes)
		themes := make([]Theme, len(dirs))
		for index, dir := range dirs {
			if dir.IsDir() {
				if err := openThemeConfig(dir.Name(), &themes[index]); err != nil {
					log.Error(err)
				} else {
					themes[index].Path = dir.Name()
				}
			}
		}
		data := fiber.Map{
			"Title":        "Theme",
			"Constants":    Constants,
			"Themes":       themes,
			"CurrentTheme": GetCurrentTheme(),
		}
		return c.Render("views/admin/theme", data, Layout)
	})

	app.Post("/theme", func(c *fiber.Ctx) error {
		theme := new(Theme)
		if err := openThemeConfig(c.FormValue("theme_path"), theme); err != nil {
			log.Error(err)
		} else {
			err = utils.PersistStruct(theme, fmt.Sprintf("%s/%s", DataDir.ThemeData, currentThemeKey))
			if err != nil {
				log.Error(err)
			}
		}
		return c.Redirect(fmt.Sprintf("%s/theme", Constants.BaseUrl))
	})
}
