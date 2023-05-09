package main

import (
	"errors" // Importing package for handling errors
	"net/http"

	"github.com/gin-gonic/gin"
)

// Defining a struct for a book
type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// Creating a slice of books
var books = []book{
	{ID: "1", Title: "Cvetje v jeseni", Author: "Ivan Tavčar", Quantity: 2},
	{ID: "2", Title: "Deseti brat", Author: "Ivan Sivec", Quantity: 5},
	{ID: "3", Title: "Martin Krpan", Author: "Fran Levstik", Quantity: 3},
}

// Handler function to get all books
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books) // Returns all books
}

// Handler function to get a book by its ID
func bookById(c *gin.Context) {
	id := c.Param("id")          // Get the id of the book from the URL parameter
	book, err := getBookById(id) // Get the book from the ID

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."}) // If book is not found, return an error message
		return
	}

	c.IndentedJSON(http.StatusOK, book) // Return the book
}

// Handler function to checkout a book
func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id") // Get the id of the book from the query parameter

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."}) // If id query parameter is missing, return an error message
		return
	}

	book, err := getBookById(id) // Get the book from the ID

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."}) // If book is not found, return an error message
		return
	}

	if book.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book not available."}) // If book is not available, return an error message
		return
	}

	book.Quantity -= 1                  // Decrement the book quantity by 1
	c.IndentedJSON(http.StatusOK, book) // Return the book
}

// Handler function to return a book
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id") // Get the id of the book from the query parameter

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter."}) // If id query parameter is missing, return an error message
		return
	}

	book, err := getBookById(id) // Get the book from the ID

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."}) // If book is not found, return an error message
		return
	}

	book.Quantity += 1                  // Increment the book quantity by 1
	c.IndentedJSON(http.StatusOK, book) // Return the book
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func removeBook(c *gin.Context) {
	id := c.Param("id")

	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	var updatedBook book

	if err := c.BindJSON(&updatedBook); err != nil {
		return
	}

	for i, b := range books {
		if b.ID == id {
			books[i] = updatedBook
			c.IndentedJSON(http.StatusOK, updatedBook)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", bookById)
	router.POST("/books", createBook)
	router.PATCH("/books/:id", updateBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
	router.DELETE("/books/:id", removeBook)
	router.Run("localhost:8080")
}
