package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})
	app.Get("/about", func(c *fiber.Ctx) error {
		// return c.SendString("about.html")
		return c.SendString("Simple golang fiber app. No contacts of dev\n links such as http://127.0.0.1:3000/ are unavaliable, dumbass")
	})

	log.Fatal(app.Listen(":3000"))
}
