package models

import "github.com/go-playground/validator/v10"

type User struct {
	ID           int    `validate:"unique"`
	Email        string `validate:"email"`
	Name         string `validate:"required"`
	LoginHash    int64  `validate:"required"`
	Balance      int    `validate:"gte:0"`
	InuseBalance int    `validate:"gte:0"`
}

<<<<<<< HEAD
var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

=======
>>>>>>> 4c848bcf28dae2fe33355c44852d0d3d9b50d0bb
func (u *User) CreateUser() error {
	// db query
	return nil
}

func (u *User) GetUserByName(name string) {

}
