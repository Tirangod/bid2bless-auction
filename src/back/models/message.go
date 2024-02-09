package models

import "time"

type Message struct {
	ID        int       `json:"id"`
	AuctionId int       `json:"auction_id"`
	SentAt    time.Time `json:"sent_at"`
	UserId    int       `json:"user_id"`
}
