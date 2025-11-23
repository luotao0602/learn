package dto

import "task4/internal/model"

type AuthResponse struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}
