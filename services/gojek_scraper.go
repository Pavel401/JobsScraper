package services

import (
	"fmt"
	"log"
	"scrapper/models"
	"scrapper/utility"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func GojekScraper() ([]models.Job, error) {
	c := colly.NewCollector()

	url := "https://jobs.lever.co/GoToGroup"

	var postings []models.Job

	c.OnError(func(response *colly.Response, err error) {
		log.Fatalf("Error while scraping the URL: %v", err)
	})
	currentTime := time.Now()

	c.OnHTML("div.posting", func(e *colly.HTMLElement) {
		location := e.ChildText("span.sort-by-location")
		applyURL := e.ChildAttr("a", "href")
		jobTitle := e.DOM.Find("h5[data-qa='posting-name']").Text()
		jobTitle = strings.TrimSpace(jobTitle)
		// Create a new job posting
		posting := models.Job{
			Title:     jobTitle,
			Location:  location,
			ApplyURL:  applyURL,
			ID:        utility.GenerateRandomID(),
			Company:   "GoJek",
			CreatedAt: currentTime.Unix(),
			ImageUrl:  "https://w7.pngwing.com/pngs/898/376/png-transparent-united-states-coursera-massive-open-online-course-education-united-states-blue-text-trademark.png",
		}
		fmt.Print(posting)
		postings = append(postings, posting)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatalf("Error visiting the URL: %v", err)
	}

	// // Print the scraped job postings
	// for i, posting := range postings {
	// 	fmt.Printf("Job #%d\n", i+1)
	// 	fmt.Printf("Title: %s\n", posting.Title)
	// 	fmt.Printf("Location: %s\n", posting.Location)
	// 	fmt.Printf("URL: %s\n", posting.ApplyURL)
	// 	fmt.Println()
	// }

	return postings, nil
}
