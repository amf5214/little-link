package main

import (
	"fmt"
	"littlelink/backend"
	"littlelink/frontend"
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Path for database file
var database_file = "test.db"

// Database connection
var db, _ = gorm.Open(sqlite.Open(database_file), &gorm.Config{})

type LinkObj struct {
	url string
}

// Path function to retrieve a url
func getLongUrl(c *gin.Context) {
	shorturl := c.Param("url")
	c.Redirect(http.StatusMovedPermanently, "http://"+backend.RetrieveUrl(db, shorturl))
}

// Path function to log a url and get a shortened url
func putShorten(c *gin.Context) {
	var link LinkObj
	if err := c.BindJSON(&link); err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Invalid Request")
		return
	}

	shortenedUrl := backend.LogUrl(db, link.url)

	c.IndentedJSON(http.StatusOK, shortenedUrl)
}

// Path for frontend to shorten url
func frontentShorten(c *gin.Context) {
	path := c.PostForm("urlpath")

	shortenedUrl := backend.LogUrl(db, path)

	c.Redirect(http.StatusMovedPermanently, "/home/completed/"+shortenedUrl)
}

// Function to start the server
func start() {

	var tcpport string

	if len(os.Args) > 1 {
		tcpport = os.Args[1]
		fmt.Print("\n\n\nPort Selected = " + tcpport + "\n\n\n")
	} else {
		tcpport = "8080"
	}

	router := gin.Default()

	router.GET("/:url", getLongUrl)
	router.PUT("/shorten", putShorten)
	router.POST("/tinylink", frontentShorten)
	router.GET("/home", func(c *gin.Context) {
		// Wraper for Gomponents handler
		frontend.GetHomePage(c.Writer, c.Request)
	})
	router.GET("/home/completed/:shorten", func(c *gin.Context) {
		// Wraper for Gomponents handler
		frontend.GetCompletedPage(c.Writer, c.Request, c.Param("shorten"))
	})

	router.Run("localhost:" + tcpport)

}

// Main function
func main() {
	start()
}
