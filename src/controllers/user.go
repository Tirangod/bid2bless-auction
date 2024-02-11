package controllers

import (
	"bid2bless/src/models"
	"encoding/json"
)

// User registration request only created when user tries to register
type userSignupData struct {
	Name      string `json:"username"`
	Email     string `json:"email"`
	LoginHash int64  `json:"login_hash"`
}

type userLoginData struct {
	LoginHash int64 `json:"login_hash"`
}

func UserSignup(rawData []byte) error {
	var userData userSignupData
	err := json.Unmarshal(rawData, &userData)
	if err != nil {
		return err
	}
	err = (&models.User{
		Name:         userData.Name,
		Email:        userData.Email,
		LoginHash:    userData.LoginHash,
	}).CreateUser()
	return err
}
