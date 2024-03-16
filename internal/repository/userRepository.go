package repository

import (
	"QuestService/internal/model/quest"
	"QuestService/internal/model/user"
	"github.com/gofrs/uuid/v5"
)

type UserRepository interface {
	CreateUser(user *user.User) error
	DeleteUserByID(id uuid.UUID) error
	FindUserByID(id uuid.UUID) (*user.User, error)
	UpdateName(id uuid.UUID, newName string) (*user.User, error)
	UpdatePassword(id uuid.UUID, newPassword string) (*user.User, error)
	UpdateEmail(id uuid.UUID, newEmail string) (*user.User, error)
	Replenishment(id uuid.UUID, money float32) (*user.User, error)
	GetHistoryOfQuest(id uuid.UUID) ([]quest.Quest, error)
}
