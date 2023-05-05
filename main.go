package main

import (
	"log"
	"user"

	"github.com/codewithed/example/go-rest-api/user"
	"github.com/gofiber/fiber/v2"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func AppIsLive(c *fiber.Ctx) error {
	return c.SendString("App is running live and coloured")
}

func main() {
	user.InitialMigration()
	// create a new instance of fiber
	app := fiber.New()

	app.Get("/", AppIsLive)

	log.Fatal(app.Listen(":8080"))

}
