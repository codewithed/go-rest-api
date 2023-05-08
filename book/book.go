package book

import (
	"fmt"

	"github.com/gofiber/fiber"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database configuration
var DB *gorm.DB
var err error

const dsn = "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"

// create book struct
type Book struct {
	gorm.Model
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

func InitialMigration() {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf(err.Error())
		panic("cannot connect to database")
	}
	DB.AutoMigrate(&Book{})
}

func GetBooks() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var books []Book
		DB.Find(&books)
		return c.JSON(&books)
	}
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params(":id")
	var book Book
	DB.Find(&book, id)
	return c.JSON(&book)
}

func AddBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(500).SendString(err.Error())
	}

	DB.Create(&book)
	return c.JSON(&book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	book := new(Book)
	DB.First(&book, id)
	if book.Title == "" {
		c.Status(500).SendString("Book not found")
	}

	if err := c.BodyParser(&book); err != nil {
		c.Status(500).SendString(err.Error())
	}
	DB.Save(&book)
	return c.JSON(&book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book Book
	DB.First(&book, id)
	if book.Title == "" {
		c.Status(500).SendString("Book not found")
	}

	DB.Delete(&book)
	return nil
}
