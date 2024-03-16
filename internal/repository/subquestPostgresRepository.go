package repository

import (
	"VK/internal/model/subquest"
	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
	"time"
)

type subquestPostgresRepository struct {
	db *sqlx.DB
}

func NewSubquestPostgresRepository(db *sqlx.DB) *subquestPostgresRepository {
	return &subquestPostgresRepository{db: db}
}

// TODO test
func (r *subquestPostgresRepository) CreateSubquest(subquest *subquest.Subquest) error {
	var id uuid.UUID
	query := "INSERT INTO subquest (title, description, status, startDate, dueDate) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	row := r.db.QueryRow(query, subquest.Title, subquest.Description, subquest.Status, subquest.StartDate, subquest.DueDate)
	if err := row.Scan(&id); err != nil {
		return err
	}
	subquest.ID = id

	return nil
}

// TODO test
func (r *subquestPostgresRepository) DeleteSubquestByID(id uuid.UUID) error {

	query := "DELETE FROM subquest where id = $1"

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

// TODO test
func (r *subquestPostgresRepository) FindSubquestByID(id uuid.UUID) (*subquest.Subquest, error) {
	query := "SELECT title, description, status, startDate, dueDate FROM subquest WHERE id = $1"

	var subquest subquest.Subquest
	err := r.db.Get(&subquest, query, id)

	return &subquest, err

}

// TODO test
func (r *subquestPostgresRepository) ChangeStatusByID(id uuid.UUID, status subquest.Status) (*subquest.Subquest, error) {
	query := "UPDATE subquest SET status=$1 WHERE id=$2 RETURNING *"

	var subquest subquest.Subquest
	err := r.db.Get(&subquest, query, status, id)
	if err != nil {
		return nil, err
	}
	return &subquest, nil

}

func (r *subquestPostgresRepository) ChangeDueDate(id uuid.UUID, newTime time.Time) (*subquest.Subquest, error) {
	query := "UPDATE subquest SET duedate=$1 WHERE id=$2 RETURNING *"

	var subquest subquest.Subquest
	err := r.db.Get(&subquest, query, newTime, id)
	if err != nil {
		return nil, err
	}

	return &subquest, nil
}
func (r *subquestPostgresRepository) ChangeTitle(id uuid.UUID, newTitle string) (*subquest.Subquest, error) {
	query := "UPDATE subquest SET title=$1 WHERE id=$2 RETURNING *"

	var subquest subquest.Subquest
	err := r.db.Get(&subquest, query, newTitle, id)
	if err != nil {
		return nil, err
	}

	return &subquest, nil
}
func (r *subquestPostgresRepository) ChangeDescription(id uuid.UUID, newDescription string) (*subquest.Subquest, error) {
	query := "UPDATE subquest SET description=$1 WHERE id=$2 RETURNING *"

	var subquest subquest.Subquest
	err := r.db.Get(&subquest, query, newDescription, id)
	if err != nil {
		return nil, err
	}

	return &subquest, nil
}
