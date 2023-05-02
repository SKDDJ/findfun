package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/skddj/findfun/backend/models"
)

func InsertGathering(db *sql.DB, gathering *models.Gathering) error {
	sqlStatement := `
	INSERT INTO gatherings (title, description, location, date, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, gathering.Title, gathering.Description, gathering.Location, gathering.Date, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return err
	}
	gathering.ID = id
	return nil
}

func GetGathering(db *sql.DB, id int) (*models.Gathering, error) {
	sqlStatement := `SELECT * FROM gatherings WHERE id=$1`
	gathering := &models.Gathering{}
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&gathering.ID, &gathering.Title, &gathering.Description, &gathering.Location, &gathering.Date, &gathering.CreatedAt, &gathering.UpdatedAt)
	switch err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("Gathering not found")
	case nil:
		return gathering, nil
	default:
		return nil, err
	}
}

func UpdateGathering(db *sql.DB, gathering *models.Gathering) error {
	sqlStatement := `
	UPDATE gatherings
	SET title=$2, description=$3, location=$4, date=$5, updated_at=$6
	WHERE id=$1`
	_, err := db.Exec(sqlStatement, gathering.ID, gathering.Title, gathering.Description, gathering.Location, gathering.Date, time.Now())
	return err
}

func DeleteGathering(db *sql.DB, id int) error {
	sqlStatement := `DELETE FROM gatherings WHERE id=$1`
	_, err := db.Exec(sqlStatement, id)
	return err
}

func InsertEvaluation(db *sql.DB, evaluation *models.Evaluation) error {
	sqlStatement := `
	INSERT INTO evaluations (gathering_id, user_id, rating, comment, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`
	id := 0
	err := db.QueryRow(sqlStatement, evaluation.GatheringID, evaluation.UserID, evaluation.Rating, evaluation.Comment, time.Now(), time.Now()).Scan(&id)
	if err != nil {
		return err
	}
	evaluation.ID = id
	return nil
}

func GetEvaluation(db *sql.DB, id int) (*models.Evaluation, error) {
	sqlStatement := `SELECT * FROM evaluations WHERE id=$1`
	evaluation := &models.Evaluation{}
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&evaluation.ID, &evaluation.GatheringID, &evaluation.UserID, &evaluation.Rating, &evaluation.Comment, &evaluation.CreatedAt, &evaluation.UpdatedAt)
	switch err {
	case sql.ErrNoRows:
		return nil, fmt.Errorf("Evaluation not found")
	case nil:
		return evaluation, nil
	default:
		return nil, err
	}
}

func UpdateEvaluation(db *sql.DB, evaluation *models.Evaluation) error {
	sqlStatement := `
	UPDATE evaluations
	SET gathering_id=$2, user_id=$3, rating=$4, comment=$5, updated_at=$6
	WHERE id=$1`
	_, err := db.Exec(sqlStatement, evaluation.ID, evaluation.GatheringID, evaluation.UserID, evaluation.Rating, evaluation.Comment, time.Now())
	return err
}

func DeleteEvaluation(db *sql.DB, id int) error {
	sqlStatement := `DELETE FROM evaluations WHERE id=$1`
	_, err := db.Exec(sqlStatement, id)
	return err
}

func GetEvaluationsForGathering(db *sql.DB, gatheringID int) ([]*models.Evaluation, error) {
	sqlStatement := `SELECT * FROM evaluations WHERE gathering_id=$1`
	rows, err := db.Query(sqlStatement, gatheringID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	evaluations := []*models.Evaluation{}
	for rows.Next() {
		evaluation := &models.Evaluation{}
		err := rows.Scan(&evaluation.ID, &evaluation.GatheringID, &evaluation.UserID, &evaluation.Rating, &evaluation.Comment, &evaluation.CreatedAt, &evaluation.UpdatedAt)
		if err != nil {
			return nil, err
		}
		evaluations = append(evaluations, evaluation)
	}
	return evaluations, nil
}
