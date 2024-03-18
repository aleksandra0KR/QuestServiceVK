package repository

import (
	"github.com/gofrs/uuid/v5"
	"time"
	"vk/internal/model/quest"
	"vk/internal/model/status"
)

type QuestRepository interface {
	CreateQuest(quest *quest.Quest) error
	DeleteQuestByID(id uuid.UUID) error
	FindQuestByID(id uuid.UUID) (*quest.Quest, error)
	ChangeStatusByID(id uuid.UUID, status status.Status) (*quest.Quest, float32, error)
	ChangeDueDate(id uuid.UUID, newTime time.Time) error
	ChangeTitle(id uuid.UUID, newTitle string) error
	ChangeDescription(id uuid.UUID, newDescription string) error
}
