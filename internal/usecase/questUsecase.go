package usecase

import (
	"VK/internal/model/quest"
	"VK/internal/model/subquest"
	"github.com/gofrs/uuid/v5"
	"time"
)

type QuestUseCase interface {
	Create(quest *quest.Quest) (*quest.Quest, error)
	DeleteQuestByID(id uuid.UUID) error
	FindQuestByID(id uuid.UUID) (*quest.Quest, error)
	ChangeDueDate(id uuid.UUID, newTime time.Time) (*quest.Quest, error)
	ChangeTitle(id uuid.UUID, newTitle string) (*quest.Quest, error)
	ChangeDescription(id uuid.UUID, newDescription string) (*quest.Quest, error)
	ShowSubquests(id uuid.UUID) ([]subquest.Subquest, error)
}
