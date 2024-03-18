package usecase

import (
	"fmt"
	"github.com/gofrs/uuid/v5"
	"time"
	"vk/internal/model/quest"
	"vk/internal/model/subquest"
	"vk/internal/repository"
)

type QuestUseCaseImplementation struct {
	repository repository.QuestRepository
}

func NewQuestUseCaseImplementation(repository repository.QuestRepository) *QuestUseCaseImplementation {
	return &QuestUseCaseImplementation{repository: repository}
}

func (uc *QuestUseCaseImplementation) Create(quest *quest.Quest) (*quest.Quest, error) {
	if quest.Title == "" {
		fmt.Errorf("title can't be empty")
	}
	err := uc.repository.CreateQuest(quest)
	if err != nil {
		return nil, err
	}

	return quest, nil
}

func (uc *QuestUseCaseImplementation) DeleteQuestByID(id uuid.UUID) error {
	err := uc.repository.DeleteQuestByID(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *QuestUseCaseImplementation) FindQuestByID(id uuid.UUID) (*quest.Quest, error) {
	foundedQuest, err := uc.repository.FindQuestByID(id)
	if err != nil {
		return nil, err
	}

	return foundedQuest, nil
}

func (uc *QuestUseCaseImplementation) ChangeDueDate(id uuid.UUID, newTime time.Time) error {
	err := uc.repository.ChangeDueDate(id, newTime)
	if err != nil {
		return err
	}

	return nil
}

func (uc *QuestUseCaseImplementation) ChangeTitle(id uuid.UUID, newTitle string) error {
	err := uc.repository.ChangeTitle(id, newTitle)
	if err != nil {
		return err
	}

	return nil
}

func (uc *QuestUseCaseImplementation) ChangeDescription(id uuid.UUID, newDescription string) error {
	err := uc.repository.ChangeDescription(id, newDescription)
	if err != nil {
		return err
	}

	return nil
}

func (uc *QuestUseCaseImplementation) ShowSubquests(id uuid.UUID) ([]subquest.Subquest, error) {
	foundedQuest, err := uc.FindQuestByID(id)
	if err != nil {
		return nil, err
	}
	return foundedQuest.Subquests, nil
}
