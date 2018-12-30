package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Bird struct {
	Species string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		fmt.Println("Error: can't converse birds to json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, writeErr := w.Write(birdListBytes)

	if writeErr != nil {
		fmt.Println("Error: can't write to response")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func createBirdHandler(w http.ResponseWriter, r * http.Request) {
	bird := Bird{}

	err := r.ParseForm()

	if err != nil {
		fmt.Println("Error: can't parse request form")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	birds = append(birds, bird)

	http.Redirect(w, r, "/assets/", http.StatusFound)
}


