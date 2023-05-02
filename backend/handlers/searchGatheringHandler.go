package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type SearchParams struct {
	Location         string    `json:"location"`
	Date             time.Time `json:"date"`
	TabletopGameType string    `json:"tabletop_game_type"`
}

func SearchGatheringHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var searchParams SearchParams
	err = json.Unmarshal(body, &searchParams)
	if err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	// Query the gatherings from your database based on searchParams.
	// This example assumes you have a function called `queryGatherings` that handles the database operation.

	gatherings, err := queryGatherings(searchParams)
	if err != nil {
		http.Error(w, "Error querying gatherings from database", http.StatusInternalServerError)
		return
	}

	gatheringsJSON, err := json.Marshal(gatherings)
	if err != nil {
		http.Error(w, "Error marshalling gatherings to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(gatheringsJSON)
}
