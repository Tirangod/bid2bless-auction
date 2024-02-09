package models

type AuctionsCategories struct {
	connectionId int `json:"connection_id"`
	auctionId    int `json:"auction_id"`
	userId       int `json:"user_id"`
}
