package main

import (
	"go_prac/accounts"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	accDB := accounts.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Main page")
	})
	app.Get("/account", accDB.GetAccount)
	app.Post("/account/create", accDB.CreateAccount)
	app.Delete("/account/delete/:name", accDB.DeleteAccount)
	log.Fatal(app.Listen(":3000"))
}
