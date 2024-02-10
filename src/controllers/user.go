package controllers

import (
	"encoding/json"
	"fmt"
)

// User registration request only created when user tries to register
type userSignupData struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	LoginHash int64  `json:"login_hash"`
}

// type userLoginData struct {
// 	LoginHash int64 `json:"login_hash"`
// }

func UserSignup(rawData []byte) error {
	var u userSignupData
	err := json.Unmarshal(rawData, &u)
	if err != nil {
		return err
	}
	fmt.Printf("Login: %s, %s, %d;\n", u.Username, u.Email, u.LoginHash)
	return nil
}
