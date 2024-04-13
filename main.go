package main

import (
	"github.com/hbourgeot/blog/handlers"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		AppName:       "henrrius.tech Blog, a blog for devs",
	})

	app.Get("/", handlers.Home)

	log.Fatal(app.Listen(":3000"))

}
