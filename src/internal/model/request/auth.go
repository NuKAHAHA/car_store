package request

import "time"

type LoginRequest struct {
	Email    string `json:"email" xml:"email"`
	Password string `json:"password" xml:"password"`
}

type RegisterRequest struct {
	Email           string    `json:"email"`
	Password        string    `json:"password"`
	ConfirmPassword string    `json:"confirmPassword"`
	Name            string    `json:"name"`
	Surname         string    `json:"surname"`
	Birthday        time.Time `json:"birthday"`
}