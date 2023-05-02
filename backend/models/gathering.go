package models

import (
	"time"

	"./database"
)

type Gathering struct {
	ID         int
	Title      string
	Location   string
	Date       time.Time
	Time       time.Time
	GameType   string
	MaxPlayers int
}

func CreateGathering(g Gathering) (int, error) {
	db := database.GetDB()

	query := `INSERT INTO gatherings (title, location, date, time, game_type, max_players) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var id int
	err := db.QueryRow(query, g.Title, g.Location, g.Date, g.Time, g.GameType, g.MaxPlayers).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func SearchGatherings(location string, date time.Time, gameType string) ([]Gathering, error) {
	db := database.GetDB()

	query := `SELECT id, title, location, date, time, game_type, max_players FROM gatherings WHERE location=$1 AND date=$2 AND game_type=$3`

	rows, err := db.Query(query, location, date, gameType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	gatherings := make([]Gathering, 0)

	for rows.Next() {
		var g Gathering
		err := rows.Scan(&g.ID, &g.Title, &g.Location, &g.Date, &g.Time, &g.GameType, &g.MaxPlayers)
		if err != nil {
			return nil, err
		}
		gatherings = append(gatherings, g)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return gatherings, nil
}
