package main

import (
	"log"

	"github.com/codewithed/go-rest-api/book"
	"github.com/gofiber/fiber/v2"
)

func AppIsLive(c *fiber.Ctx) error {
	return c.SendString("App is running live and coloured")
}

func Routers(app *fiber.App) {
	app.Get("/books", book.GetBooks)
	app.Get("/books/:id", book.GetBook)
	app.Post("/books/:id", book.AddBook)
	app.Put("/books/:id", book.UpdateBook)
	app.Delete("/books/:id", book.DeleteBook)
}

func main() {
	book.InitialMigration()

	// create a new instance of fiber
	app := fiber.New()
	app.Get("/", AppIsLive)
	Routers(app)
	log.Fatal(app.Listen(":8080"))

}
