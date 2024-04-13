package main

import (
	"github.com/a-h/templ"
	"github.com/hbourgeot/blog/handlers"
	"github.com/hbourgeot/blog/helpers"
	view "github.com/hbourgeot/blog/views"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "henrrius.tech Blog, a blog for devs",
	})

	MountRoutes(app)

	log.Fatal(app.Listen(":3000"))

}

func NotFoundMiddleware(c fiber.Ctx) error {
	return helpers.Render(c, view.NotFound(), templ.WithStatus(http.StatusNotFound))
}

func MountRoutes(app *fiber.App) {
	// static files
	app.Static("/static/", "./static", fiber.Static{
		Browse: true,
	})

	// routes
	app.Get("/", handlers.Home)
	app.Get("/foo", handlers.Foo)

	// middlewares
	app.Use(NotFoundMiddleware)
}
