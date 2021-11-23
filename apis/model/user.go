package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

//user struct (Model)
type User struct {
	// Id         string `json:"_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	University string `json:"university"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateUserInput struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	University string `json:"university" binding:"required"`
}

type UpdateInput struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	University string `json:"university" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user *User) BeforeCreate() (err error) {
	// user.ID = uuid.New().String()
	user.Password, err = HashPassword(user.Password)
	if err != nil {
		return err
	}
	// user.Id = user.Email
	user.Created_at = time.Now().Local().String()
	user.Updated_at = time.Now().Local().String()
	return nil
}

func HashPassword(password string) (string, error) {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}
