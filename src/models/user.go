package models

import (
	"bid2bless/src/database"
	"errors"
	"strconv"
	"strings"
)

var db *database.DB = database.GetDB()

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"username"`
	LoginHash    int64  `json:"login_hash"`
	Balance      int    `json:"balance"`
	InuseBalance int    `json:"inuse_balance"`
}

func (u *User) validate() error {
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

func (u *User) Create() error {
	err := u.validate()

	if err != nil {
		return err
	}

	query := `
		insert into 
			users (email, name, login_hash, balance, inuse_balance)
		values ($1, $2, $3, $4, $5);
	`
	db.RawQuery(query, u.Email, u.Name, u.LoginHash, u.Balance, u.InuseBalance)
	return nil
}

func (u *User) Exists() bool {
	n, err := db.CountWhere("users", "users.login_hash == "+strconv.FormatInt(u.LoginHash, 10))
	if err != nil {
		return false
	}
	return n > 0
}

func GetUserByID(id int64) User {
	return User{}
}

func GetUserByName(name string) User {
	return User{}
}
