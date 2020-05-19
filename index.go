package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Error struct {
	ErrorStatus  int
	ErrorMessage string
}

type Request struct {
	Question string `json:"question"`
}

func (r *Request) validate() []Error {
	var errors []Error

	if strings.TrimSpace(r.Question) == "" {
		errors = append(errors, Error{ErrorStatus: 400, ErrorMessage: "You did not ask any Question"})
		return errors
	}

	return errors
}

type Response struct {
	Answer string `json:"answer"`
}

func generateAnswer() string {
	answers := []string{
		"Yes",
		"No",
	}

	seed := time.Now().UnixNano()
	rand.Seed(seed)

	return answers[rand.Intn(len(answers))]
}

func askMagicConch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("content-type", "application/json")

	var req Request
	var buffer bytes.Buffer

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if validationErr := req.validate(); len(validationErr) > 0 {
		http.Error(w, validationErr[0].ErrorMessage, validationErr[0].ErrorStatus)
		return
	}

	answer := Response{Answer: generateAnswer()}
	if err := json.NewEncoder(&buffer).Encode(&answer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(buffer.String()))
}

func main() {
	http.HandleFunc("/", askMagicConch)
	log.Print("server is up on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
