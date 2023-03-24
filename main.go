package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var Books = []Book{
	{ID: 1, Title: "Book A", Author: "Author A", Desc: "Desc A"},
	{ID: 2, Title: "Book B", Author: "Author B", Desc: "Desc B"},
}

func main() {
	router := gin.Default()

	// GET All Books
	router.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": Books})
	})

	// GET Book By ID
	router.GET("/books/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}

		for _, book := range Books {
			if book.ID == id {
				c.JSON(http.StatusOK, gin.H{"data": book})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	})

	// Add Book
	router.POST("/books", func(c *gin.Context) {
		var book Book
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		book.ID = len(Books) + 1
		Books = append(Books, book)

		c.JSON(http.StatusCreated, gin.H{"data": book})
	})

	// Update Book
	router.PUT("/books/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}

		var updatedBook Book
		if err := c.ShouldBindJSON(&updatedBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for i, book := range Books {
			if book.ID == id {
				Books[i].Title = updatedBook.Title
				Books[i].Author = updatedBook.Author
				Books[i].Desc = updatedBook.Desc

				c.JSON(http.StatusOK, gin.H{"data": Books[i]})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	})

	// Delete Book
	router.DELETE("/books/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
			return
		}

		for i, book := range Books {
			if book.ID == id {
				Books = append(Books[:i], Books[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"data": true})
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
	})

	router.Run(":8080")
}
