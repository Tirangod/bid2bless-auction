package models

type User struct {
	Email          string `validate:"email"`
	Name           string `validate:"required"`
	HashedPassword int64  `validate:"required"`
	Balance        int    `validate:"gte:0"`
	InuseBalance   int    `validate:"gte:0"`
}

var _ User
