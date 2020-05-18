package main

import (
	"math/rand"
	"time"
)

func generateAnswer() string {
	answers := []string{
		"Yes",
		"No",
	}

	seed := time.Now().UnixNano()
	rand.Seed(seed)

	return answers[rand.Intn(len(answers))]
}

func main() {
	println(generateAnswer())
}
