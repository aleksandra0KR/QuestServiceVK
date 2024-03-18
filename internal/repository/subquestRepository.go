package repository

import (
	"github.com/gofrs/uuid/v5"
	"time"
	"vk/internal/model/status"
	"vk/internal/model/subquest"
)

type SubquestRepository interface {
	CreateSubquest(subquest *subquest.Subquest) error
	DeleteSubquestByID(id uuid.UUID) error
	FindSubquestByID(id uuid.UUID) (*subquest.Subquest, error)
	ChangeStatusByID(id uuid.UUID, status status.Status) (*subquest.Subquest, error)
	ChangeDueDate(id uuid.UUID, newTime time.Time) error
	ChangeTitle(id uuid.UUID, newTitle string) error
	ChangeDescription(id uuid.UUID, newDescription string) error
	AttachToQuest(questId uuid.UUID, subquestId uuid.UUID) error
}
