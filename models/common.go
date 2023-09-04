package models

// Define a struct to represent the JSON response from the API.
type Job struct {
	ID        string `json:"id"`
	Title     string `json:"text"`
	Location  string `json:"location"`
	CreatedAt int64  `json:"createdAt"`
	Company   string `json:"company"`
	ApplyURL  string `json:"applyUrl"`
}
