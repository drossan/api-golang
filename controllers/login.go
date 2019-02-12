package controllers

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"net/http"

	"../commons"
	"../configuration"
	"../models"
	"github.com/labstack/echo"
)

// Login -
func Login(c echo.Context) (err error) {
	// Bind
	user := new(models.User)
	if err = c.Bind(user); err != nil {
		return
	}

	db := configuration.GetConnection()
	defer db.Close()

	ps := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", ps)

	db.Where("email = ? and password = ?", user.Email, pwd).First(&user)
	if user.ID > 0 {
		user.Password = ""

		// Create token
		t := models.Token{
			UserID: user.ID,
			Email:  user.Email,
			RolID:  user.Level,
			Token:  updateUserToken(50),
		}

		return c.JSON(http.StatusOK, echo.Map{
			"token": commons.GenerateJWT(t),
		})
	} else {
		return echo.ErrUnauthorized
	}
}

// UserCreate - Register new user
func UserCreate(c echo.Context) (err error) {
	user := new(models.User)
	if err = c.Bind(user); err != nil {
		return err
	}

	if user.Password != user.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, "Error: password do not match")
	}

	pw := sha256.Sum256([]byte(user.Password))
	pwd := fmt.Sprintf("%x", pw)
	user.Password = pwd

	db := configuration.GetConnection()
	defer db.Close()

	err = db.Create(&user).Error
	if err != nil {
		m := fmt.Sprintf("Error: fail to created register, %s", err)
		return c.JSON(http.StatusBadRequest, m)
	}

	return c.JSON(http.StatusCreated, "User register")
}

func updateUserToken(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, n)

	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}

	user := models.User{}
	db := configuration.GetConnection()
	defer db.Close()
	db.Model(&user).Update("token", string(b))

	return string(b)
}
