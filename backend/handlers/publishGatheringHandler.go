package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Gathering struct {
	Title            string    `json:"title"`
	Location         string    `json:"location"`
	Date             time.Time `json:"date"`
	Time             time.Time `json:"time"`
	TabletopGameType string    `json:"tabletop_game_type"`
	MaxPlayers       int       `json:"max_players"`
}

func PublishGatheringHandler(w http.ResponseWriter, r *http.Request) {
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

	var gathering Gathering
	err = json.Unmarshal(body, &gathering)
	if err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	// Insert the gathering data into your database here.
	// This example assumes you have a function called `insertGathering` that handles the database operation.

	err = insertGathering(gathering)
	if err != nil {
		http.Error(w, "Error inserting gathering into database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Gathering published successfully")
}

/*This code defines a Gathering struct with fields to store the gathering information.
The publishGatheringHandler function checks if the request method is POST,
reads the request body, unmarshals the JSON data into a Gathering struct,
 and then inserts the data into your database. In this example,
  I've assumed you have a function called insertGathering that handles the database operation
  which you'll need to implement based on your specific database).*/
