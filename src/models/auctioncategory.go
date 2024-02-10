package models

type AuctionCategory struct {
	connectionId int
	auctionId    int
	userId       int
}

var _ AuctionCategory
