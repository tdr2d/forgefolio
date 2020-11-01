package main

import (
	"core/templates"
	"fmt"
	"math/rand"
	"strings"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
)

var log = logrus.New()

func main() {
	app := fiber.New()

	// GET /john
	app.Get("/:name", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("Hello, %s 👋!", c.Params("name"))
		return c.SendString(msg) // => Hello john 👋!
	})

	// GET /john/75
	app.Get("/:name/:age/:gender?", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("👴 %s is %s years old", c.Params("name"), c.Params("age"))
		return c.SendString(msg) // => 👴 john is 75 years old
	})

	// GET /dictionary.txt
	app.Get("/:file.:ext", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("📃 %s.%s", c.Params("file"), c.Params("ext"))
		return c.SendString(msg) // => 📃 dictionary.txt
	})

	// GET /flights/LAX-SFO
	app.Get("/flights/:from-:to", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("💸 From: %s, To: %s", c.Params("from"), c.Params("to"))
		return c.SendString(msg) // => 💸 From: LAX, To: SFO
	})

	// GET /api/register
	app.Get("/api/*", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", c.Params("*"))
		return c.SendString(msg) // => ✋ register
	})

	log.Fatal(app.Listen(":3000"))
}

func main() {
	log.Printf("starting the server at http://localhost:8080 ...")
	err := fasthttp.ListenAndServe("localhost:8080", requestHandler)
	if err != nil {
		log.Fatalf("unexpected error in server: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	var path string = string(ctx.Path())
	logger.Info("access %s", path))

	if strings.HasPrefix(path, "assets/") {
		fasthttp.ServeFile(ctx, path)
	} else {
		switch path {
		case "/":
			mainPageHandler(ctx)
		case "/table":
			tablePageHandler(ctx)
		default:
			errorPageHandler(ctx)
		}
		ctx.SetContentType("text/html; charset=utf-8")
	}
}

func mainPageHandler(ctx *fasthttp.RequestCtx) {
	p := &templates.MainPage{
		CTX: ctx,
	}
	templates.WritePageTemplate(ctx, p)
}

func tablePageHandler(ctx *fasthttp.RequestCtx) {
	rowsCount := ctx.QueryArgs().GetUintOrZero("rowsCount")
	if rowsCount == 0 {
		rowsCount = 10
	}
	p := &templates.TablePage{
		Rows: generateRows(rowsCount),
	}
	templates.WritePageTemplate(ctx, p)
}

func errorPageHandler(ctx *fasthttp.RequestCtx) {
	p := &templates.ErrorPage{
		Path: ctx.Path(),
	}
	templates.WritePageTemplate(ctx, p)
	ctx.SetStatusCode(fasthttp.StatusBadRequest)
}

func generateRows(rowsCount int) []string {
	var rows []string
	for i := 0; i < rowsCount; i++ {
		r := fmt.Sprintf("row %d", i)
		if rand.Intn(20) == 0 {
			r = "bingo"
		}
		rows = append(rows, r)
	}
	return rows
}
