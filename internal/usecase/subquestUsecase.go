package usecase

import (
	"VK/internal/model/subquest"
	"github.com/gofrs/uuid/v5"
	"time"
)

type SubquestUseCase interface {
	Create(subquest *subquest.Subquest) (*subquest.Subquest, error)
	DeleteSubquestByID(id uuid.UUID) error
	FindSubquestByID(id uuid.UUID) (*subquest.Subquest, error)
	ChangeDueDate(id uuid.UUID, newTime time.Time) (*subquest.Subquest, error)
	ChangeTitle(id uuid.UUID, newTitle string) (*subquest.Subquest, error)
	ChangeDescription(id uuid.UUID, newDescription string) (*subquest.Subquest, error)
}
