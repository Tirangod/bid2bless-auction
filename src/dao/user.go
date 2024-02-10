package dao

import "errors"

type UserDAO struct {
	email          string
	name           string
	hashedPassword string
	balance        int
	inuseBalance   int
}

func isEmailValid(email string) bool {
	// TODO
	return true
}

func isNameValid(email string) bool {
	// TODO
	return true
}

func isHashedPasswordValid(email string) bool {
	// TODO
	return true
}

func CreateUser(email string, name string, hashedPassword string) (*UserDAO, error) {
	if isEmailValid(email) &&
		isNameValid(name) &&
		isHashedPasswordValid(hashedPassword) {
		return nil, errors.New("email, name or password are invalid")
	}
	return &UserDAO{
		email,
		name,
		hashedPassword,
		0,
		0,
	}, nil
}
