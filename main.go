package main

import (
	"log"

	"github.com/codewithed/go-rest-api/book"
	"github.com/gofiber/fiber/v2"
)

func AppIsLive(c *fiber.Ctx) error {
	return c.SendString("App is running live and coloured")
}

func main() {
	book.InitialMigration()
	// create a new instance of fiber
	app := fiber.New()

	app.Get("/", AppIsLive)

	log.Fatal(app.Listen(":8080"))

}
