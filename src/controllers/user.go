package controllers

import (
	"bid2bless/src/models"
	"encoding/json"
	"os"

	"github.com/charmbracelet/log"
	"github.com/go-playground/validator/v10"
)

var logger *log.Logger = log.New(os.Stdout)
var validate *validator.Validate = validator.New(validator.WithRequiredStructEnabled())

// User registration request only created when user tries to register
type userSignupData struct {
	Name      string 
	Email     string `json:"email"`
	LoginHash int64  `json:"login_hash"`
}

func UserSignup(rawData []byte) error {
	var userData userSignupData
	err := json.Unmarshal(rawData, &userData)
	if err != nil {
		return err
	}
	userModel := &models.User{
		Name:      userData.Name,
		Email:     userData.Email,
		LoginHash: userData.LoginHash,
	}

	err = validate.Struct(userModel)
	if err != nil {
		logger.Error("Invalid user model!")
		return err
	}

	err = userModel.CreateUser()
	if err != nil {
		logger.Error("Unable to create user!")
		return err
	}
	return nil
}

type userLoginData struct {
	LoginHash int64 `json:"login_hash"`
}

func UserLogin(rawData []byte) error {
	return nil
}
