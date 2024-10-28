package main

import (
	"littlelink/backend"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Path for database file
var database_file = "test.db"

// Database connection
var db, _ = gorm.Open(sqlite.Open(database_file), &gorm.Config{})

type Link struct {
	url string
}

// Path function to retrieve a url
func getLongUrl(c *gin.Context) {
	shorturl := c.Param("url")
	c.Redirect(http.StatusMovedPermanently, "http://"+backend.RetrieveUrl(db, shorturl))
}

// Path function to log a url and get a shortened url
func putShorten(c *gin.Context) {
	var link Link
	if err := c.BindJSON(&link); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Invalid Request")
		return
	}

	shortenedUrl := backend.LogUrl(db, link.url)

	c.IndentedJSON(http.StatusOK, shortenedUrl)
}

// Function to start the server
func start() {

	router := gin.Default()

	router.GET("/:url", getLongUrl)
	router.PUT("/shorten", putShorten)

	router.Run("localhost:8080")
}

// Main function
func main() {
	start()
}
