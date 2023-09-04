package models

// Define a struct to match the structure of the API response.
type CRED struct {
	ID         string `json:"id"`
	Categories struct {
		Location string `json:"location"`
	} `json:"categories"`
	Text      string `json:"text"`
	CreatedAt int64  `json:"createdAt"`
}
