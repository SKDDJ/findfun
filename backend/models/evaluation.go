package models

import (
	"./database"
)

type Evaluation struct {
	ID          int
	Name        string
	Email       string
	Phone       string
	GatheringID int
	Evaluation  string
	Rating      int
}

func CreateEvaluation(e Evaluation) (int, error) {
	id, err := saveEvaluation(e)
	return id, err
}

func saveEvaluation(e Evaluation) (int, error) {
	db := database.GetDB()

	query := `INSERT INTO evaluations (name, email, phone, gathering_id, evaluation, rating) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var id int
	err := db.QueryRow(query, e.Name, e.Email, e.Phone, e.GatheringID, e.Evaluation, e.Rating).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
