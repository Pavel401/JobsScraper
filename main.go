package main

import (
	"net/http"
	"scrapper/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize a new Gin router.
	r := gin.Default()

	// Define a route to fetch and return the list of Posting structs as JSON.
	r.GET("/cred", handlers.GetPostingsHandler)
	r.GET("/atlassian", handlers.AtlassianHandler)
	r.GET("/amazon", handlers.Amazonhandler)

	// Define a route for the root path.
	r.GET("/", RootHandler)

	// Start the Gin server on port 8080.
	r.Run(":3003")
}

// RootHandler handles the root route (/).
func RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"hello": "Mabud"})
}
