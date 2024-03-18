package usecase

import (
	"fmt"
	"github.com/gofrs/uuid/v5"
	"vk/internal/model/status"
	"vk/internal/model/user"
	"vk/internal/model/usershistory"
	"vk/internal/repository"
)

type UserUseCaseImplementation struct {
	repository repository.UserRepository
}

func NewUserUseCaseImplementation(repository repository.UserRepository) *UserUseCaseImplementation {
	return &UserUseCaseImplementation{repository: repository}
}

func (uc *UserUseCaseImplementation) Create(user *user.User) (*user.User, error) {
	if user.Email == "" || user.Password == "" || user.Username == "" {
		fmt.Errorf("all fileds can't be empty")
	}
	err := uc.repository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCaseImplementation) DeleteUserByID(id uuid.UUID) error {
	err := uc.repository.DeleteUserByID(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCaseImplementation) FindUserByID(id uuid.UUID) (*user.User, error) {
	foundedUser, err := uc.repository.FindUserByID(id)
	if err != nil {
		return nil, err
	}

	return foundedUser, nil
}

func (uc *UserUseCaseImplementation) UpdateName(id uuid.UUID, newName string) error {
	err := uc.repository.UpdateName(id, newName)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCaseImplementation) UpdatePassword(id uuid.UUID, newPassword string) error {

	err := uc.repository.UpdatePassword(id, newPassword)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCaseImplementation) UpdateEmail(id uuid.UUID, newEmail string) error {
	err := uc.repository.UpdateEmail(id, newEmail)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCaseImplementation) ShowBalance(id uuid.UUID) (float32, error) {
	user, err := uc.FindUserByID(id)
	if err != nil {
		return 0, err
	}
	return user.Balance, nil
}

func (uc *UserUseCaseImplementation) GetHistoryOfQuest(id uuid.UUID) (usershistory.UsersHistory, error) {

	var history usershistory.UsersHistory
	history, err := uc.repository.GetHistoryOfQuest(id)
	if err != nil {
		return history, err
	}
	return history, nil
}

func (uc *UserUseCaseImplementation) AttachToQuest(questId uuid.UUID, userId uuid.UUID) error {
	err := uc.repository.AttachToQuest(questId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCaseImplementation) AttachToSubquest(subquestId uuid.UUID, userId uuid.UUID) error {
	err := uc.repository.AttachToSubquest(subquestId, userId)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCaseImplementation) ChangeSubquestsStatus(subquestId uuid.UUID, userId uuid.UUID, status status.Status) error {
	err := uc.repository.ChangeSubquestsStatus(subquestId, userId, status)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUseCaseImplementation) ChangeQuestsStatus(questId uuid.UUID, userId uuid.UUID, status status.Status) error {
	err := uc.repository.ChangeQuestsStatus(questId, userId, status)
	if err != nil {
		return err
	}

	return nil
}
