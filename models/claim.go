package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim Token de user
type Claim struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	RolID  uint   `json:"role_id"`
	Token  string `json:"token"`
	Admin  uint
	jwt.StandardClaims
}
