package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/hbourgeot/blog/helpers"
	"github.com/hbourgeot/blog/views"
	"github.com/hbourgeot/blog/views/layout"
	partial "github.com/hbourgeot/blog/views/partials"
)

func Home(c fiber.Ctx) error {
	return helpers.Render(c, layout.Base(views.Index()))
}

func Foo(c fiber.Ctx) error {
	return helpers.Render(c, layout.Base(partial.Foo()))
}
