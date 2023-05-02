package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type MatchParams struct {
	PreferredGameType string `json:"preferred_game_type"`
	PreferredLocation string `json:"preferred_location"`
}

type MatchResult struct {
	MatchedGatherings []Gathering `json:"matched_gatherings"`
}

func MatchGatheringHandler(w http.ResponseWriter, r *http.Request) {
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

	var matchParams MatchParams
	err = json.Unmarshal(body, &matchParams)
	if err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	// Query the gatherings from your database based on matchParams.
	// This example assumes you have a function called `findMatchingGatherings` that handles the database operation.

	matchedGatherings, err := findMatchingGatherings(matchParams)
	if err != nil {
		http.Error(w, "Error finding matching gatherings from database", http.StatusInternalServerError)
		return
	}

	matchResult := MatchResult{MatchedGatherings: matchedGatherings}
	matchResultJSON, err := json.Marshal(matchResult)
	if err != nil {
		http.Error(w, "Error marshalling match result to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(matchResultJSON)
}

/*This code defines the MatchParams struct to hold the matching parameters, the MatchResult struct to hold the matched gatherings, and the MatchGatheringHandler function to handle the matching request. The handler reads the request body, unmarshals it into the MatchParams struct, queries the database using the provided matching parameters, and returns the matched gatherings as a JSON response.

You should replace the findMatchingGatherings function with your actual database query function.*/
