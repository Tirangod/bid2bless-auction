package models

import "time"

type BetHistory struct {
	ID         int       `json:"id"`
	AuctionId  int       `json:"auction_id"`
	UpdatedBet int       `json:"updated_bet"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserId     int       `json:"user_id"`
}
