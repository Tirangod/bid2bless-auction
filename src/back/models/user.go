package models

type Users struct {
	ID             int    `json:"id"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	HashedPassword string `json:"hashedPassword"`
	Balance        int    `json:"balance"`
	InuseBalance   int    `json:"inuse_balance"`
}
