package utility

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateRandomID generates an 8-digit random ID.
func GenerateRandomID() string {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate an 8-digit random number
	id := rand.Intn(100000000)

	// Format the number as an 8-digit string with leading zeros
	idStr := fmt.Sprintf("%08d", id)

	return idStr
}
