package main

import (
	"context"
	"log"
	"net/http"
	"scrapper/handlers"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"

	// "github.com/joho/godotenv"
	// "google.golang.org/api/option"
	"github.com/spf13/viper"
)

//func init() {
//	viper.SetConfigFile("ENV") // Set the name of the configuration file
//	viper.ReadInConfig()
//	viper.AutomaticEnv()
//
//	// Optionally, set a default value for an environment variable if it's not set.
//	// This can be useful to avoid panics when trying to read unset variables.
//	viper.SetDefault("PORT", "8080")
//}

func main() {
	// Initialize a new Gin router.
	r := gin.Default()

	viper.SetConfigFile("ENV") // Set the name of the configuration file
	viper.ReadInConfig()
	viper.AutomaticEnv()

	// Optionally, set a default value for an environment variable if it's not set.
	// This can be useful to avoid panics when trying to read unset variables.
	//viper.SetDefault("PORT", "8080")

	// Define a route to fetch and return the list of Posting structs as JSON.
	r.GET("/cred", handlers.GetPostingsHandler)
	r.GET("/atlassian", handlers.AtlassianHandler)
	r.GET("/amazon", handlers.Amazonhandler)
	r.GET("/coursera", handlers.CourseraHandler)
	r.GET("/freshworks", handlers.FreshWorksHandler)
	r.GET("/gojek", handlers.Gojekhandler)
	r.GET("/mpl", handlers.MplHandler)

	r.GET("/all", handlers.AllScrapersHandler)
	r.GET("/sync", handlers.SyncAll)

	r.GET("/getallJobs", handlers.GetJobsFromDB)

	// Read environment variables using Viper.
	projectID := viper.GetString("FIREBASE_PROJECT_ID")
	credentialsFile := viper.GetString("/app/firebase-config.json")
	port := viper.GetString("PORT")

	log.Printf("Project ID: %s", projectID)
	log.Printf("Credentials file: %s", credentialsFile)
	log.Printf("Port: %s", port)

	// Initialize Firestore client using environment variables.
	opt := option.WithCredentialsFile(credentialsFile)
	firestoreClient, err := firestore.NewClient(context.Background(), projectID, opt)
	if err != nil {
		log.Fatalf("Error initializing Firestore client: %v", err)
	}

	r.GET("/syncFirestore", func(c *gin.Context) {
		handlers.SyncWithFireBase(c, firestoreClient)
	})

	r.GET("/getJobsFromFirestore", func(c *gin.Context) {
		handlers.GetJobsFromFirestore(c, firestoreClient)
	})

	// Define a route for the root path.
	r.GET("/", RootHandler)
	r.Use(CORSMiddleware())

	// Start the Gin server on the specified port.
	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// RootHandler handles the root route (/).
func RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"hello": "Mabud"})
}
