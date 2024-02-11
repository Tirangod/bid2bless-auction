package controllers

import (
	"bid2bless/src/models"
	"encoding/json"
	"errors"
	"os"

	"github.com/charmbracelet/log"
)

var logger *log.Logger = log.New(os.Stdout)

// User registration request only created when user tries to register
func UserSignup(rawData []byte) error {
	var userModel models.User
	err := json.Unmarshal(rawData, &userModel)
	if err != nil {
		return err
	}

	err = userModel.Create()
	if err != nil {
		logger.Error("Unable to create user!", err)
		return err
	}
	return nil
}

func UserLogin(rawData []byte) error {
	var userModel models.User
	err := json.Unmarshal(rawData, &userModel)
	if err != nil {
		return err
	}
	if !userModel.Exists() {
		return errors.New("user hasnt found")
	}
	return nil
}
