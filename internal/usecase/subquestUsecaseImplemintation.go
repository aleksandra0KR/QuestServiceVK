package usecase

import (
	"fmt"
	"github.com/gofrs/uuid/v5"
	"time"
	"vk/internal/model/subquest"
	"vk/internal/repository"
)

type SubquestUseCaseImplementation struct {
	repository repository.SubquestRepository
}

func NewSubquestUseCaseImplementation(repository repository.SubquestRepository) *SubquestUseCaseImplementation {
	return &SubquestUseCaseImplementation{repository: repository}
}

func (uc *SubquestUseCaseImplementation) Create(subquest *subquest.Subquest) (*subquest.Subquest, error) {
	if subquest.Title == "" {
		fmt.Errorf("title can't be empty")
	}
	err := uc.repository.CreateSubquest(subquest)
	if err != nil {
		return nil, err
	}

	return subquest, nil
}

func (uc *SubquestUseCaseImplementation) DeleteSubquestByID(id uuid.UUID) error {
	err := uc.repository.DeleteSubquestByID(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *SubquestUseCaseImplementation) FindSubquestByID(id uuid.UUID) (*subquest.Subquest, error) {
	foundedSubquest, err := uc.repository.FindSubquestByID(id)
	if err != nil {
		return nil, err
	}

	return foundedSubquest, nil
}

func (uc *SubquestUseCaseImplementation) ChangeDueDate(id uuid.UUID, newTime time.Time) error {
	err := uc.repository.ChangeDueDate(id, newTime)
	if err != nil {
		return err
	}

	return nil
}

func (uc *SubquestUseCaseImplementation) ChangeTitle(id uuid.UUID, newTitle string) error {
	err := uc.repository.ChangeTitle(id, newTitle)
	if err != nil {
		return err
	}

	return nil
}

func (uc *SubquestUseCaseImplementation) ChangeDescription(id uuid.UUID, newDescription string) error {
	err := uc.repository.ChangeDescription(id, newDescription)
	if err != nil {
		return err
	}

	return nil
}

func (uc *SubquestUseCaseImplementation) AttachToQuest(questId uuid.UUID, subquestId uuid.UUID) error {
	err := uc.repository.AttachToQuest(questId, subquestId)
	if err != nil {
		return err
	}

	return nil
}
