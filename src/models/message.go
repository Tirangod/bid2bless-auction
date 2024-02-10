package models

import "time"

type Message struct {
	id        int
	auctionId int
	sentAt    time.Time
	userId    int
}

var _ Message
