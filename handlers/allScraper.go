package handlers

import (
	"net/http"
	"scrapper/models"
	"scrapper/services"

	"github.com/gin-gonic/gin"
)

// Defined a list of scraper functions
var scrapers = []struct {
	name    string
	scraper func() ([]models.Job, error)
}{
	{
		name:    "Amazon",
		scraper: services.AmazonScrapper,
	},
	{
		name:    "Atlassian",
		scraper: services.AtlassianScrapper,
	},
	{
		name:    "Coursera",
		scraper: services.CourseraScraper,
	},
	{
		name:    "CRED",
		scraper: services.CredScraper,
	}, {
		name:    "FreshWorks",
		scraper: services.FreshWorksScraper,
	},
	{
		name:    "Gojek",
		scraper: services.GojekScraper,
	},
	{
		name:    "MPL",
		scraper: services.MplScrapper,
	},
}

func AllScrapersHandler(c *gin.Context) {
	var allPostings []models.Job

	// Loop through the list of scrapers and call each one.
	for _, scraper := range scrapers {
		postings, err := scraper.scraper()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		allPostings = append(allPostings, postings...)
	}

	c.JSON(http.StatusOK, allPostings)
}
