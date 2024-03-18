package repository

import (
	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
	"time"
	"vk/internal/model/status"
	"vk/internal/model/subquest"
)

type subquestPostgresRepository struct {
	db *sqlx.DB
}

func NewSubquestPostgresRepository(db *sqlx.DB) *subquestPostgresRepository {
	return &subquestPostgresRepository{db: db}
}

func (r *subquestPostgresRepository) CreateSubquest(subquest *subquest.Subquest) error {
	var id uuid.UUID
	query := `INSERT INTO "subquest" ("title", "description", "status", "startDate", "dueDate") VALUES ($1, $2, $3, $4, $5) RETURNING id`

	row := r.db.QueryRow(query, subquest.Title, subquest.Description, subquest.Status, subquest.StartDate, subquest.DueDate)
	if err := row.Scan(&id); err != nil {
		return err
	}
	subquest.ID = id

	return nil
}

func (r *subquestPostgresRepository) DeleteSubquestByID(id uuid.UUID) error {

	query := `DELETE FROM "subquest" where "id" = $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *subquestPostgresRepository) FindSubquestByID(id uuid.UUID) (*subquest.Subquest, error) {
	query := `SELECT "title", "description", "status", "startDate", "dueDate" FROM "subquest" WHERE "id" = $1`

	var subquest subquest.Subquest
	err := r.db.Get(&subquest, query, id)

	return &subquest, err
}

func (r *subquestPostgresRepository) ChangeStatusByID(id uuid.UUID, status status.Status) (*subquest.Subquest, error) {
	query := `UPDATE "subquest" SET "status"=$1 WHERE "id"=$2 RETURNING *`

	var subquest subquest.Subquest
	err := r.db.Get(&subquest, query, status, id)
	if err != nil {
		return nil, err
	}
	return &subquest, nil
}

func (r *subquestPostgresRepository) ChangeDueDate(id uuid.UUID, newTime time.Time) error {
	query, err := r.db.Prepare(`UPDATE "subquest" SET "dueDate"=$1 WHERE "id"=$2`)
	if err != nil {
		return err
	}
	_, err = query.Exec(newTime, id)

	if err != nil {
		return err
	}
	return nil
}

func (r *subquestPostgresRepository) ChangeTitle(id uuid.UUID, newTitle string) error {
	query, err := r.db.Prepare(`UPDATE "subquest" SET "title"=$1 WHERE "id"=$2`)
	if err != nil {
		return err
	}
	_, err = query.Exec(newTitle, id)

	if err != nil {
		return err
	}
	return nil
}

func (r *subquestPostgresRepository) ChangeDescription(id uuid.UUID, newDescription string) error {
	query, err := r.db.Prepare(`UPDATE "subquest" SET "description"=$1 WHERE "id"=$2`)
	if err != nil {
		return err
	}
	_, err = query.Exec(newDescription, id)

	if err != nil {
		return err
	}
	return nil
}

func (r *subquestPostgresRepository) AttachToQuest(questId uuid.UUID, subquestId uuid.UUID) error {
	stmt, err := r.db.Prepare(`INSERT INTO "subquestOfQuest" ("subquest_id", "quest_id") VALUES ($1, $2)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(questId, subquestId)
	if err != nil {
		return err
	}

	return nil
}
