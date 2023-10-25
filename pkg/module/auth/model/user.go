package model

type User struct {
	Id                 string
	Username           string
	PasswordHash       string
	Email              string
	EmailVerified      bool
	Phone              string
	PhoneVerified      bool
	NormalizedUsername string
	NormalizedEmail    string
	LoginFailures      int
	IsLocked           bool
}

type UserFilter struct {
}
