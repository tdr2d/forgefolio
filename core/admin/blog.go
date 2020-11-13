package admin

import (
	"core/utils"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// BlogPost is the structure of the post
type BlogPost struct {
	Id      string `json:"id" xml:"id" form:"id"`
	Time    uint   `json:"time" xml:"time" form:"time"`
	Version string `json:"version" xml:"version" form:"version"`
	Blocks  string `json:"blocks" xml:"blocks" form:"blocks"`
}

type BlogPostList struct {
	Name string
	Path string
	Date string
}

func nameFromId(id string) string {
	nameRegex := regexp.MustCompile(`BP__[0-9]*__(.*)`)
	return strings.ReplaceAll(string(nameRegex.FindSubmatch([]byte(id))[1]), "_", " ")
}

// BlogController implement blog crud operations
func BlogController(app *fiber.App) {

	app.Get("/blog-posts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "new" {
			return c.Render("admin/blog/form", fiber.Map{"Title": "New Blog Posts", "IsNew": true, "Navigation": Navigation}, "layouts/main")
		}
		bp := new(BlogPost)
		if err := utils.ReadStruct(bp, fmt.Sprintf("%s/%s", BlogDataDir, id)); err != nil {
			log.Error(err)
		}
		return c.Render("admin/blog/form", fiber.Map{"Title": nameFromId(bp.Id), "IsNew": false, "Post": bp, "Navigation": Navigation}, "layouts/main")
	})

	app.Get("/blog-posts", func(c *fiber.Ctx) error {
		files, err := ioutil.ReadDir(BlogDataDir)
		posts := make([]BlogPostList, len(files))
		for i, item := range files {
			posts[i].Name = nameFromId(item.Name())
			posts[i].Path = item.Name()
			posts[i].Date = item.ModTime().Format("2006-01-02 15:04:05")
		}
		if err != nil {
			log.Error(err)
		}
		return c.Render("admin/blog/list", fiber.Map{"Title": "Blog Posts", "Posts": posts, "Navigation": Navigation}, "layouts/main")
	})

	app.Post("/blog-posts", func(c *fiber.Ctx) error {
		bp := new(BlogPost)
		if err := c.BodyParser(bp); err != nil {
			log.Error(err)
			return err
		}
		if err := utils.PersistStruct(bp, fmt.Sprintf("%s/%s", BlogDataDir, bp.Id)); err != nil {
			log.Error(err)
			return err
		}
		return c.SendStatus(200)
	})

	app.Patch("blog-posts/:id", func(c *fiber.Ctx) error {
		oldId := c.Params("id")
		bp := new(BlogPost)
		if err := c.BodyParser(bp); err != nil {
			log.Error(err)
			return err
		}
		if err := os.Rename(fmt.Sprintf("%s/%s", BlogDataDir, oldId), fmt.Sprintf("%s/%s", BlogDataDir, bp.Id)); err != nil {
			log.Error(err)
			return err
		}
		if err := utils.PersistStruct(bp, fmt.Sprintf("%s/%s", BlogDataDir, bp.Id)); err != nil {
			log.Error(err)
			return err
		}
		return c.SendStatus(200)
	})

	app.Delete("blog-posts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if err := os.Remove(fmt.Sprintf("%s/%s", BlogDataDir, id)); err != nil {
			log.Error(err)
			return err
		}
		return c.SendStatus(200)
	})
}
