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

func FreshWorksScraper() ([]models.Job, error) {
	c := colly.NewCollector()

	url := "https://careers.smartrecruiters.com/Freshworks"

	var postings []models.Job
	// c.Limit(&colly.LimitRule{
	// 	DomainGlob:  "*",
	// 	Parallelism: 2,
	// 	Delay:       2 * time.Second,
	// })

	c.OnError(func(response *colly.Response, err error) {
		log.Fatalf("Error while scraping the URL: %v", err)
	})
	currentTime := time.Now()

	c.OnHTML("section.openings-section", func(e *colly.HTMLElement) {
		posting := models.Job{
			Title: strings.TrimSpace(e.ChildText("h4.details-title")),

			Location: strings.TrimSpace(e.ChildText("h3.opening-title")),
			ApplyURL: e.ChildAttr("a", "href"),

			ID:        utility.GenerateRandomID(),
			Company:   "Freshworks",
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
