package user

import (
	"crypto/sha256"
	"encoding/base64"
)

type User struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserPage struct {
	PageIndex int    `json:"page_index"`
	PageSize  int    `json:"page_size"`
	Count     int    `json:"count"`
	Data      []User `json:"data"`
}

func (user *User) SetPassword(password string) {
	user.Password = user.getPasswordHash(password)
}

func (user *User) ComparePassword(password string) bool {
	return user.Password == user.getPasswordHash(password)
}

func (user *User) getPasswordHash(password string) string {
	b := sha256.Sum256([]byte(password))
	d := base64.StdEncoding.EncodeToString(b[:])
	return d
}
