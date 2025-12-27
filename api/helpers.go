package payment_processor

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// GetRandomNumber returns a random number between min and max
func GetRandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

// SendNotification sends a notification to the user
func SendNotification(user string, message string) {
	fmt.Printf("Sending notification to %s: %s\n", user, message)
}

// HandleError handles an error and logs it
func HandleError(err error) {
	log.Println(err)
}

// GetRandomString returns a random string of length n
func GetRandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	const (
		letterIdxBits = 6                    // 6 bits to represent 64 unique letters
		letterIdxMax  = 1 << letterIdxBits - 1
	)
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(letterIdxMax)]
	}
	return string(b)
}