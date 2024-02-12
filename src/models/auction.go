package models

import (
	"bid2bless/src/database"
	"time"
)

type Auction struct {
	ID          int       `json:"ID"`
	Name        string    `json:"name"`
	Photo       string    `json:"photo"`
	Description string    `json:"description"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
	CreatedAt   time.Time `json:"created_at"`
	InitialBet  int       `json:"initial_bet"`
	Bet         int       `json:"bet"`
	UserId      int       `json:"user_id"`
}

func (a *Auction) Create(UserId int) (*Auction, error) {
	db := database.GetDB()
	var res Auction
	query := "INSERT INTO auctions (name, photo, description, start_at, end_at, created_at, initial_bet, bet, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING ID, name, photo, description, start_at,end_at, created_at, initial_bet, bet, user_id"
	row, err := db.RawQuery(query, a.Name, a.Photo, a.Description, a.StartAt, a.EndAt, a.CreatedAt, a.InitialBet, a.Bet, UserId)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&res.ID,
		&res.Name,
		&res.Photo,
		&res.Description,
		&res.StartAt,
		&res.EndAt,
		&res.CreatedAt,
		&res.InitialBet,
		&res.Bet,
		&res.UserId,
	)
	return &res, err

}

func (a *Auction) Get() (*Auction, error) {
	db := database.GetDB()
	var res Auction
	query := "SELECT name, photo, description, start_at, end_at, created_at, initial_bet, bet, user_id FROM auctions WHERE id=$1 LIMIT 1"
	row, err := db.RawQuery(query, a.ID)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&res.ID,
		&res.Name,
		&res.Photo,
		&res.Description,
		&res.StartAt,
		&res.EndAt,
		&res.CreatedAt,
		&res.InitialBet,
		&res.Bet,
		&res.UserId,
	)
	return &res, err
}

func (a *Auction) Update() (*Auction, error) {
	db := database.GetDB()
	var res Auction
	query := "UPDATE auctions SET name=$2, photo=$3, description=$4, start_at=$5, end_at=$6, initial_bet=$7, bet=$8 WHERE id=$1 RETURNING ID, name, photo, description, start_at,end_at, created_at, initial_bet, bet, user_id"
	row, err := db.RawQuery(query, a.ID, a.Name, a.Photo, a.Description, a.StartAt, a.EndAt, a.InitialBet, a.Bet)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&res.ID,
		&res.Name,
		&res.Photo,
		&res.Description,
		&res.StartAt,
		&res.EndAt,
		&res.CreatedAt,
		&res.InitialBet,
		&res.Bet,
		&res.UserId,
	)
	return &res, err
}

func (a *Auction) List() ([]Auction, error) {
	db := database.GetDB()

	query := "SELECT name, photo, description, start_at, end_at, created_at, initial_bet, bet, user_id FROM auctions"
	rows, err := db.RawQuery(query, a.ID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	res := []Auction{}
	for rows.Next() {
		var i Auction
		err = rows.Scan(
			&i.ID,
			&i.Name,
			&i.Photo,
			&i.Description,
			&i.StartAt,
			&i.EndAt,
			&i.CreatedAt,
			&i.InitialBet,
			&i.Bet,
			&i.UserId,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, i)
	}

	return res, err
}

func (a *Auction) Delete() error {
	db := database.GetDB()
	query := "DELETE FROM auctions WHERE id=$1"
	_, err := db.RawQuery(query, a.ID)
	return err
}
