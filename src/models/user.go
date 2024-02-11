package models

type User struct {
	ID           int    `validate:"unique"`
	Email        string `validate:"email" json:"email"`
	Name         string `validate:"required, unique" json:"username"`
	LoginHash    int64  `validate:"required" json:"login_hash"`
	Balance      int    `validate:"gte:0"`
	InuseBalance int    `validate:"gte:0"`
}

func (u *User) Create() error {
	// db query
	return nil
}

func (u *User) Exists() bool {
	return false
}

func GetUserByID(id int64) User {
	return User{}
}

func GetUserByName(name string) User {
	return User{}
}
