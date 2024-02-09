package models

type AuctionCategory struct {
	ConnectionId int `json:"connection_id"`
	AuctionId    int `json:"auction_id"`
	UserId       int `json:"user_id"`
}
