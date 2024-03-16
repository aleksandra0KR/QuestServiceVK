package repository

// TODO test

import (
	"VK/internal/model/quest"
	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
	"time"
)

type questPostgresRepository struct {
	db *sqlx.DB
}

func NewQuestPostgresRepository(db *sqlx.DB) *questPostgresRepository {
	return &questPostgresRepository{db: db}
}

// TODO test
func (r *questPostgresRepository) CreateQuest(quest *quest.Quest) error {
	var id uuid.UUID
	query := "INSERT INTO quest (title, description, status, startDate, dueDate, repeatable, Reward) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	row := r.db.QueryRow(query, quest.Title, quest.Description, quest.Status, quest.StartDate, quest.DueDate, quest.Repeatable, quest.Reward)
	if err := row.Scan(&id); err != nil {
		return err
	}
	quest.ID = id

	return nil
}

// TODO test
func (r *questPostgresRepository) DeleteQuestByID(id uuid.UUID) error {

	query := "DELETE FROM quest where id = $1"

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// TODO test
func (r *questPostgresRepository) FindQuestByID(id uuid.UUID) (*quest.Quest, error) {
	query := "SELECT title, description, status, startDate, dueDate, repeatable, Reward FROM quest WHERE id = $1"

	var quest quest.Quest
	err := r.db.Get(&quest, query, id)

	return &quest, err

}

// TODO test
func (r *questPostgresRepository) ChangeStatusByID(id uuid.UUID, status quest.Status) (*quest.Quest, float32, error) {
	query := "UPDATE quest SET status=$1 WHERE id=$2 RETURNING *"

	var quest quest.Quest
	err := r.db.Get(&quest, query, status, id)
	if err != nil {
		return nil, 0, err
	}

	if status == "Done" {
		return &quest, quest.Reward, nil
	}

	return &quest, 0, nil

}

func (r *questPostgresRepository) ChangeDueDate(id uuid.UUID, newTime time.Time) (*quest.Quest, error) {
	query := "UPDATE quest SET duedate=$1 WHERE id=$2 RETURNING *"

	var quest quest.Quest
	err := r.db.Get(&quest, query, newTime, id)
	if err != nil {
		return nil, err
	}

	return &quest, nil
}
func (r *questPostgresRepository) ChangeTitle(id uuid.UUID, newTitle string) (*quest.Quest, error) {
	query := "UPDATE quest SET title=$1 WHERE id=$2 RETURNING *"

	var quest quest.Quest
	err := r.db.Get(&quest, query, newTitle, id)
	if err != nil {
		return nil, err
	}

	return &quest, nil
}
func (r *questPostgresRepository) ChangeDescription(id uuid.UUID, newDescription string) (*quest.Quest, error) {
	query := "UPDATE quest SET description=$1 WHERE id=$2 RETURNING *"

	var quest quest.Quest
	err := r.db.Get(&quest, query, newDescription, id)
	if err != nil {
		return nil, err
	}

	return &quest, nil
}
