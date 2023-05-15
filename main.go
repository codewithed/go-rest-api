package main

import (
	"net/http"

	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func createBook(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func RemoveIndex(s []book, index int) ([]book, error) {
	if index < 0 || index >= len(s) {
		return nil, errors.New("index out of range")
	}
	return append(s[:index], s[index+1:]...), nil
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, "Book not found")
		return
	}

	c.IndentedJSON(http.StatusFound, book)
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, "Book not found")
		return
	}
	index, err := strconv.Atoi(id)
	books, err = RemoveIndex(books, index-1)
	c.JSON(http.StatusOK, book.Title+" deleted")

}

func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, "missing id query parameter")
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, "Book not found")
		return
	}

	if book.Quantity <= 0 {
		c.JSON(http.StatusNoContent, book.Title+" is out of stock")
		return
	}

	book.Quantity = book.Quantity - 1
	c.JSON(http.StatusOK, book.Title+" has been successfully checked out")
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, "missing id query parameter")
		return
	}

	book, err := getBookById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, "Book not found")
		return
	}

	book.Quantity = book.Quantity + 1
	c.JSON(http.StatusOK, book.Title+" has been has returned successfully")
	return
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "App is running live")
	})
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBook)
	router.POST("/books", createBook)
	router.DELETE("/books/:id", deleteBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.Run("localhost:8080")
}
