package user

import "github.com/gofrs/uuid/v5"

type User struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	Balance  float64   `json:"balance,omitempty"`
}
