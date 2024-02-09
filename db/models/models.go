package db

import "time"

type Auctions struct {
	ID          int       `json:"id"`
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
type BetHistory struct {
	ID         int       `json:"id"`
	AuctionId  int       `json:"auction_id"`
	UpdatedBet int       `json:"updated_bet"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserId     int       `json:"user_id"`
}
