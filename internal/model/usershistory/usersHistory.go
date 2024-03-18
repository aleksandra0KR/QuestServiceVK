package usershistory

import (
	"vk/internal/model/quest"
	"vk/internal/model/user"
)

type UsersHistory struct {
	User    user.User     `json:"user,omitempty"`
	Quest   []quest.Quest `json:"quest,omitempty"`
	Balance float32       `json:"balance,omitempty"`
}
