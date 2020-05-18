package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type Request struct {
	Question string `json:"question"`
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

	var req Request
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	answer := Response{Answer: generateAnswer()}
	var buffer bytes.Buffer
	err = json.NewEncoder(&buffer).Encode(&answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write([]byte(buffer.String()))
}

func main() {
	http.HandleFunc("/", askMagicConch)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
