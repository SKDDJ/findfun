package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/skddj/findfun/backend/database"
)

type EvaluationData struct {
	Name        string
	Email       string
	Phone       string
	GatheringID int
	Evaluation  string
	Rating      int
}

func saveEvaluationData(evaluationData EvaluationData) error {
	db := database.GetDB()

	query := `INSERT INTO evaluations (name, email, phone, gathering_id, evaluation, rating) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := db.Exec(query, evaluationData.Name, evaluationData.Email, evaluationData.Phone, evaluationData.GatheringID, evaluationData.Evaluation, evaluationData.Rating)
	return err
}

func EvaluateSystemHandler(w http.ResponseWriter, r *http.Request) {
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

	var evaluationData EvaluationData
	err = json.Unmarshal(body, &evaluationData)
	if err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	err = saveEvaluationData(evaluationData)
	if err != nil {
		http.Error(w, "Error saving evaluation data to database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
