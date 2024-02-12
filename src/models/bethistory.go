package models

import (
	"bid2bless/src/database"
	"time"
)

type BetHistory struct {
	ID         int       `json:"id"`
	AuctionId  int       `json:"auction_id"`
	UpdatedBet int       `json:"updated_bet"`
	UpdatedAt  time.Time `json:"updated_at"`
	UserId     int       `json:"user_id"`
}

func (a *BetHistory) Create(userId string) (*BetHistory, error) {
	db := database.GetDB()
	var res BetHistory
	query := "INSERT INTO bet_history (auction_id, updated_bet, user_id) VALUES ($1, $2, $3) RETURNING id, auction_id, category_id"
	row, err := db.RawQuery(query, a.AuctionId, a.UpdatedBet, userId)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&res.ID,
		&res.AuctionId,
		&res.UpdatedBet,
		&res.UpdatedAt,
		&res.UserId,
	)
	return &res, err

}

func (a *BetHistory) Get() (*BetHistory, error) {
	db := database.GetDB()
	var res BetHistory
	query := "SELECT auction_id, updated_bet, updated_at, user_id FROM bet_history WHERE id=$1 LIMIT 1"
	row, err := db.RawQuery(query, a.ID)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&res.ID,
		&res.AuctionId,
		&res.UpdatedBet,
		&res.UpdatedAt,
		&res.UserId,
	)
	return &res, err
}

func (a *BetHistory) ListByAuction() ([]BetHistory, error) {
	db := database.GetDB()

	query := "SELECT id, auction_id, updated_bet, updated_at, user_id FROM bet_history WHERE auction_id=$1"
	rows, err := db.RawQuery(query, a.AuctionId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	res := []BetHistory{}
	for rows.Next() {
		var i BetHistory
		err = rows.Scan(
			&i.ID,
			&i.AuctionId,
			&i.UpdatedBet,
			&i.UpdatedAt,
			&i.UserId,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, i)
	}

	return res, err
}

func (a *BetHistory) Delete() error {
	db := database.GetDB()
	query := "DELETE FROM bet_history WHERE id=$1"
	_, err := db.RawQuery(query, a.ID)
	return err
}
