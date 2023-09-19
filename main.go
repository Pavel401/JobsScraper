package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"scrapper/handlers"
	"strings"

	"github.com/goccy/go-json"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

// FirebaseConfig represents the structure of the Firebase configuration.
type FirebaseConfig struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
	UniversalDomain         string `json:"universal_domain"` // New field
}

func main() {
	// Initialize a new Gin router.

	viper.SetConfigFile(".env") // Set the name of the configuration file
	viper.ReadInConfig()
	viper.AutomaticEnv()

	// Load environment variables using Viper.
	projectID := viper.GetString("FIREBASE_PROJECT_ID")
	// credentialsFile := viper.GetString("FIREBASE_CREDENTIALS_FILE")
	port := viper.GetString("PORT")
	universalDomain := viper.GetString("UNIVERSAL_DOMAIN")

	privateKey := viper.GetString("private_key")
	stype := viper.GetString("type")

	privateKey = strings.Replace(privateKey, "\\n", "\n", -1)
	// log.Printf("Project ID: %s", projectID)
	// log.Printf("Credentials file: %s", privateKey)
	// log.Printf("Port: %s", port)
	// log.Printf("Universal Domain: %s", universalDomain)
	log.Printf("Type is : %s", stype)

	// Generate the FirebaseConfig struct.
	fireBaseConfig := FirebaseConfig{
		Type:                    stype,
		ProjectID:               viper.GetString("project_id"),
		PrivateKeyID:            viper.GetString("private_key_id"),
		PrivateKey:              privateKey,
		ClientEmail:             viper.GetString("client_email"),
		ClientID:                viper.GetString("client_id"),
		AuthURI:                 viper.GetString("auth_uri"),
		TokenURI:                viper.GetString("token_uri"),
		AuthProviderX509CertURL: viper.GetString("auth_provider_x509_cert_url"),
		ClientX509CertURL:       viper.GetString("client_x509_cert_url"),
		UniversalDomain:         universalDomain,
	}

	// Convert the FirebaseConfig to a JSON string.
	fireBaseConfigJSON, err := json.MarshalIndent(fireBaseConfig, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling FirebaseConfig to JSON: %v", err)
	}

	log.Print(fireBaseConfig.PrivateKey)

	// Write the JSON data to the file.
	filePath := "./firebase-config.json"
	err = ioutil.WriteFile(filePath, fireBaseConfigJSON, 0644)
	if err != nil {
		log.Fatalf("Error writing JSON file: %v", err)
	}

	fmt.Println("JSON file written:", filePath)

	r := gin.Default()

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
	projectID = viper.GetString("FIREBASE_PROJECT_ID")
	// credentialsFile := viper.GetString("FIREBASE_CREDENTIALS_FILE")
	port = viper.GetString("PORT")

	log.Printf("Project ID: %s", projectID)
	// log.Printf("Credentials file: %s", credentialsFile)
	log.Printf("Port: %s", port)

	// Initialize Firestore client using environment variables.
	opt := option.WithCredentialsFile(filePath)
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
