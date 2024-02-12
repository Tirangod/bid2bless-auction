package models

import "bid2bless/src/database"

type AuctionCategory struct {
	ConnectionId int `json:"connection_id"`
	AuctionId    int `json:"auction_id"`
	CategoryId   int `json:"category_id"`
}

func (a *AuctionCategory) Create() (*AuctionCategory, error) {
	db := database.GetDB()
	var res AuctionCategory
	query := "INSERT INTO auctions_categories (auction_id, category_id) VALUES ($1, $2) RETURNING id, auction_id, category_id"
	row, err := db.RawQuery(query, a.AuctionId, a.CategoryId)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&res.ConnectionId,
		&res.AuctionId,
		&res.CategoryId,
	)
	return &res, err

}

func (a *AuctionCategory) Get() (*AuctionCategory, error) {
	db := database.GetDB()
	var res AuctionCategory
	query := "SELECT auction_id, category_id FROM auctions_categories WHERE id=$1 LIMIT 1"
	row, err := db.RawQuery(query, a.ConnectionId)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&res.ConnectionId,
		&res.AuctionId,
		&res.CategoryId,
	)
	return &res, err
}

func (a *AuctionCategory) Update() (*AuctionCategory, error) {
	db := database.GetDB()
	var res AuctionCategory
	query := "UPDATE auctions_categories SET auction_id=$2, category_id=$3 WHERE id=$1 RETURNING id, auction_id, category_id"
	row, err := db.RawQuery(query, a.ConnectionId, a.AuctionId, a.CategoryId)
	if err != nil {
		return nil, err
	}

	err = row.Scan(
		&res.ConnectionId,
		&res.AuctionId,
		&res.CategoryId,
	)
	return &res, err
}

func (a *AuctionCategory) ListByCategory() ([]AuctionCategory, error) {
	db := database.GetDB()

	query := "SELECT id, auction_id, category_id FROM auctions_categories WHERE category_id=$1"
	rows, err := db.RawQuery(query, a.CategoryId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	res := []AuctionCategory{}
	for rows.Next() {
		var i AuctionCategory
		err = rows.Scan(
			&i.ConnectionId,
			&i.AuctionId,
			&i.CategoryId,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, i)
	}

	return res, err
}

func (a *AuctionCategory) ListByAuction() ([]AuctionCategory, error) {
	db := database.GetDB()

	query := "SELECT id, auction_id, category_id FROM auctions_categories WHERE auction_id=$1"
	rows, err := db.RawQuery(query, a.AuctionId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	res := []AuctionCategory{}
	for rows.Next() {
		var i AuctionCategory
		err = rows.Scan(
			&i.ConnectionId,
			&i.AuctionId,
			&i.CategoryId,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, i)
	}

	return res, err
}

func (a *AuctionCategory) Delete() error {
	db := database.GetDB()
	query := "DELETE FROM auctions_categories WHERE id=$1"
	_, err := db.RawQuery(query, a.ConnectionId)
	return err
}
