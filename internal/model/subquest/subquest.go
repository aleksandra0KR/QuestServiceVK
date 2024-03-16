package subquest

import (
	"github.com/gofrs/uuid/v5"
	"time"
)

type Subquest struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Status      Status    `json:"status,omitempty"`
	StartDate   time.Time `json:"startDate,omitempty"`
	DueDate     time.Time `json:"dueDate,omitempty"`
}
