package models

import (
	"errors"
	"strings"
)

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"username"`
	LoginHash    int64  `json:"login_hash"`
	Balance      int    `json:"balance"`
	InuseBalance int    `json:"inuse_balance"`
}

func (u *User) Create() error {
	// db query
	return nil
}

func (u *User) Exists() bool {
	return false
}

func (u *User) Validate() error {
	var errs []error
	if u.ID < 0 {
		errs = append(errs, errors.New("invalid ID in user model"))
	}
	if len(u.Email) < 5 {
		errs = append(errs, errors.New("invalid Email in user model"))
	}
	if !strings.Contains(u.Email, "@") {
		errs = append(errs, errors.New("invalid Email in user model"))
	}
	if u.LoginHash == 0 {
		errs = append(errs, errors.New("invalid LoginHash in user model"))
	}
	if u.Balance < 0 {
		errs = append(errs, errors.New("invalid Balance in user model"))
	}
	if u.InuseBalance < 0 {
		errs = append(errs, errors.New("invalid InuseBalance in user model"))
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func GetUserByID(id int64) User {
	return User{}
}

func GetUserByName(name string) User {
	return User{}
}
