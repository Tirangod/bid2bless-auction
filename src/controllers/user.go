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
func UserSignup(rawData []byte) error {
	var userModel models.User
	err := json.Unmarshal(rawData, &userModel)
	if err != nil {
		return err
	}

	err = validate.Struct(&userModel)
	if err != nil {
		logger.Error("Invalid user model!")
		return err
	}

	err = userModel.Create()
	if err != nil {
		logger.Error("Unable to create user!")
		return err
	}
	return nil
}

func UserLogin(rawData []byte) error {
	return nil
}
