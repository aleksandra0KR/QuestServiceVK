package repository

import (
	"VK/internal/model/quest"
	"github.com/gofrs/uuid/v5"
	"time"
)

type QuestRepository interface {
	CreateQuest(quest *quest.Quest) error
	DeleteQuestByID(id uuid.UUID) error
	FindQuestByID(id uuid.UUID) (*quest.Quest, error)
	ChangeStatusByID(id uuid.UUID, status quest.Status) (*quest.Quest, float32, error)
	ChangeDueDate(id uuid.UUID, newTime time.Time) (*quest.Quest, error)
	ChangeTitle(id uuid.UUID, newTitle string) (*quest.Quest, error)
	ChangeDescription(id uuid.UUID, newDescription string) (*quest.Quest, error)
}
