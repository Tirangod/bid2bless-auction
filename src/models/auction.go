package models

import (
	"bid2bless/src/database"
	"time"
)

type Auction struct {
	id          int
	name        string
	photo       string
	description string
	startAt     time.Time
	endAt       time.Time
	createdAt   time.Time
	initialBet  int
	bet         int
	userId      int
}

var _ Auction

func (a *Auction) Create(userId int) (*Auction, error) {
	db := database.GetDB()
	var res Auction
	query := "INSERT INTO auctions (name, photo, description, start_at, end_at, created_at, initial_bet, bet, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, name, photo, description, start_at,end_at, created_at, initial_bet, bet, user_id"
	row, err := db.RawQuery(query, a.name, a.photo, a.description, a.startAt, a.endAt, a.createdAt, a.initialBet, a.bet, userId)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&res.id,
		&res.name,
		&res.photo,
		&res.description,
		&res.startAt,
		&res.endAt,
		&res.createdAt,
		&res.initialBet,
		&res.bet,
		&res.userId,
	)
	return &res, err

}
