package models

type AuctionCategory struct {
	ID        int `json:"id"`
	AuctionId int `json:"auction_id"`
	UserId    int `json:"user_id"`
}

var _ AuctionCategory
