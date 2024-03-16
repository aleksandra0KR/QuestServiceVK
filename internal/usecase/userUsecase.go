package usecase

import (
	"VK/internal/model/quest"
	"VK/internal/model/user"
	"github.com/gofrs/uuid/v5"
)

type UserUseCase interface {
	Create(user *user.User) (*user.User, error)
	DeleteUserByID(id uuid.UUID) error
	FindUserByID(id uuid.UUID) (*user.User, error)
	UpdateName(id uuid.UUID, newName string) (*user.User, error)
	UpdatePassword(id uuid.UUID, newPassword string) (*user.User, error)
	UpdateEmail(id uuid.UUID, newEmail string) (*user.User, error)
	ShowBalance(id uuid.UUID) (float32, error)
	GetHistoryOfQuest(id uuid.UUID) ([]quest.Quest, error)
}
