package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	UserRepository
	QuestRepository
	SubquestRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository:     NewUserPostgresRepository(db),
		QuestRepository:    NewQuestPostgresRepository(db),
		SubquestRepository: NewSubquestPostgresRepository(db),
	}
}
