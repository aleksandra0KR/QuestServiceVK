package usecase

import (
	"github.com/gofrs/uuid/v5"
	"time"
	"vk/internal/model/quest"
	"vk/internal/model/subquest"
)

type QuestUseCase interface {
	Create(quest *quest.Quest) (*quest.Quest, error)
	DeleteQuestByID(id uuid.UUID) error
	FindQuestByID(id uuid.UUID) (*quest.Quest, error)
	ChangeDueDate(id uuid.UUID, newTime time.Time) error
	ChangeTitle(id uuid.UUID, newTitle string) error
	ChangeDescription(id uuid.UUID, newDescription string) error
	ShowSubquests(id uuid.UUID) ([]subquest.Subquest, error)
}
