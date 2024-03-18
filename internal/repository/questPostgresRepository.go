package repository

import (
	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
	"time"
	"vk/internal/model/quest"
	"vk/internal/model/status"
)

type questPostgresRepository struct {
	db *sqlx.DB
}

func NewQuestPostgresRepository(db *sqlx.DB) *questPostgresRepository {
	return &questPostgresRepository{db: db}
}

func (r *questPostgresRepository) CreateQuest(quest *quest.Quest) error {
	var id uuid.UUID
	query := ` INSERT INTO "quest" ("title", "description", "status", "startDate", "dueDate", "repeatable", "reward") VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	row := r.db.QueryRow(query, quest.Title, quest.Description, quest.Status, quest.StartDate, quest.DueDate, quest.Repeatable, quest.Reward)
	if err := row.Scan(&id); err != nil {
		return err
	}
	quest.ID = id

	return nil
}

func (r *questPostgresRepository) DeleteQuestByID(id uuid.UUID) error {

	query := `DELETE FROM "quest" where "id" = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *questPostgresRepository) FindQuestByID(id uuid.UUID) (*quest.Quest, error) {
	query := `SELECT "id" , "title", "description", "status", "startDate", "dueDate", "repeatable", "reward" FROM quest WHERE "id" = $1`

	var quest quest.Quest
	err := r.db.Get(&quest.ID, quest.Title, quest.Description, quest.Status, &quest.StartDate, &quest.DueDate, &quest.Repeatable, &quest.Reward, query, id)
	if err != nil {
		return nil, err
	}

	return &quest, err

}

func (r *questPostgresRepository) ChangeStatusByID(id uuid.UUID, givenstatus status.Status) (*quest.Quest, float32, error) {
	query := `UPDATE "quest" SET "status"=$1 WHERE "id"=$2 RETURNING *`

	var quest quest.Quest
	err := r.db.Get(&quest, query, givenstatus, id)
	if err != nil {
		return nil, 0, err
	}

	if givenstatus == status.Done {
		return &quest, quest.Reward, nil
	}

	return &quest, 0, nil

}

func (r *questPostgresRepository) ChangeDueDate(id uuid.UUID, newTime time.Time) error {
	query, err := r.db.Prepare(`UPDATE "quest" SET "dueDate"=$1 WHERE "id"=$2`)
	if err != nil {
		return err
	}
	_, err = query.Exec(newTime, id)

	if err != nil {
		return err
	}
	return nil
}

func (r *questPostgresRepository) ChangeTitle(id uuid.UUID, newTitle string) error {
	query, err := r.db.Prepare(`UPDATE "quest" SET "title"=$1 WHERE "id"=$2`)
	if err != nil {
		return err
	}
	_, err = query.Exec(newTitle, id)

	if err != nil {
		return err
	}
	return nil
}

func (r *questPostgresRepository) ChangeDescription(id uuid.UUID, newDescription string) error {
	query, err := r.db.Prepare(`UPDATE "quest" SET "description"=$1 WHERE "id"=$2`)
	if err != nil {
		return err
	}
	_, err = query.Exec(newDescription, id)

	if err != nil {
		return err
	}
	return nil
}
