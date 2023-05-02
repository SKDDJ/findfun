package main

import (
	"net/http"

	"./handlers"

	"./database"
)

func main() {
	// Connect to the database
	database.ConnectDB()

	// Register handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/publish-gathering", handlers.PublishGatheringHandler)
	http.HandleFunc("/search-gathering", handlers.SearchGatheringHandler)
	http.HandleFunc("/match-gathering", handlers.MatchGatheringHandler)
	http.HandleFunc("/evaluate-system", handlers.EvaluateSystemHandler)

	// Start the web server
	http.ListenAndServe(":8080", nil)
}
