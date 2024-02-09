package models

import "time"

type Messages struct {
	ID        int       `json:"id"`
	AuctionId int       `json:"auction_id"`
	SentAt    time.Time `json:"sent_at"`
	UserId    int       `json:"user_id"`
}
