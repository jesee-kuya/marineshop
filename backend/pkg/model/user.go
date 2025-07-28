package model

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) ValidateUserDetails() error {
	username := strings.TrimSpace(user.Username)
	email := strings.TrimSpace(user.Email)
	password := strings.TrimSpace(user.Password)
	role := strings.TrimSpace(user.Role)
	status := strings.TrimSpace(user.Status)

	if username == "" {
		return errors.New("username is required")
	}

	if email == "" || !ValidateEmail(email) {
		return errors.New("email is required")
	}

	if password == "" || len(password) < 6 {
		return errors.New("password is required")
	}

	if role == "" || (role != "admin" && role != "user" && role != "seller") {
		return errors.New("role is required")
	}

	if status == "" || (status != "active" && status != "inactive") {
		return errors.New("status is required")
	}

	return nil
}

func ValidateEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
