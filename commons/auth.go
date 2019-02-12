package commons

import (
	"log"
	"time"

	"../models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT - generates the token for the customer
func GenerateJWT(data models.Token) string {
	claims := models.Claim{
		UserID: data.UserID,
		Email:  data.Email,
		Token:  data.Token,
		Admin:  data.RolID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "Api - cloud",
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		log.Fatal("Error: Couldn't signed the token")
	}

	return t
}
