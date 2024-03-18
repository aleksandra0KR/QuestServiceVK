package usecase

import (
	"github.com/gofrs/uuid/v5"
	"vk/internal/model/status"
	"vk/internal/model/user"
	"vk/internal/model/usershistory"
)

type UserUseCase interface {
	Create(user *user.User) (*user.User, error)
	DeleteUserByID(id uuid.UUID) error
	FindUserByID(id uuid.UUID) (*user.User, error)
	UpdateName(id uuid.UUID, newName string) error
	UpdatePassword(id uuid.UUID, newPassword string) error
	UpdateEmail(id uuid.UUID, newEmail string) error
	ShowBalance(id uuid.UUID) (float32, error)
	GetHistoryOfQuest(id uuid.UUID) (usershistory.UsersHistory, error)
	AttachToQuest(questId uuid.UUID, userId uuid.UUID) error
	AttachToSubquest(subquestId uuid.UUID, userId uuid.UUID) error
	ChangeSubquestsStatus(subquestId uuid.UUID, userId uuid.UUID, status status.Status) error
	ChangeQuestsStatus(questId uuid.UUID, userId uuid.UUID, status status.Status) error
}
