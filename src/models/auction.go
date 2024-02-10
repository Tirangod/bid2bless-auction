package models

import "time"

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
	u1serId     int
}

var _ Auction
