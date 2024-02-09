package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// User registration request only created when user tries to register
type signupRequestData struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	LoginHash int64  `json:"login_hash"`
}

type loginRequestData struct {
	LoginHash int64 `json:"login_hash"`
}

