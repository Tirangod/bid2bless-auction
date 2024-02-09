package models

type User struct {
	id             int
	email          string
	name           string
	hashedPassword string
	balance        int
	inuseBalance   int
}

var _ User
