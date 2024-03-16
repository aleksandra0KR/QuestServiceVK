package repository

import (
	"VK/internal/model/quest"
	"VK/internal/model/subquest"
	"VK/internal/model/user"
	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
)

type userPostgresRepository struct {
	db *sqlx.DB
}

func NewUserPostgresRepository(db *sqlx.DB) *userPostgresRepository {
	return &userPostgresRepository{db: db}
}

func (r *userPostgresRepository) CreateUser(user *user.User) error {
	var id uuid.UUID
	query := "INSERT INTO users (username, email, password, balance) VALUES ($1, $2, $3, $4) RETURNING id"

	row := r.db.QueryRow(query, user.Username, user.Email, user.Password, user.Balance)
	if err := row.Scan(&id); err != nil {
		return err
	}
	user.ID = id

	return nil
}

func (r *userPostgresRepository) DeleteUserByID(id uuid.UUID) error {

	query := "DELETE FROM users where id = $1"

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userPostgresRepository) FindUserByID(id uuid.UUID) (*user.User, error) {
	query := "SELECT id, username, email, password, balance FROM users WHERE id = $1"

	var user user.User
	err := r.db.Get(&user, query, id)

	return &user, err

}

func (r *userPostgresRepository) UpdateName(id uuid.UUID, newName string) (*user.User, error) {
	query := "UPDATE users SET username=$1 WHERE id=$2 RETURNING *"

	var user user.User
	err := r.db.Get(&user, query, newName, id)
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *userPostgresRepository) UpdatePassword(id uuid.UUID, newPassword string) (*user.User, error) {
	query := "UPDATE users SET password=$1 WHERE id=$2 RETURNING *"

	var user user.User
	err := r.db.Get(&user, query, newPassword, id)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (r *userPostgresRepository) UpdateEmail(id uuid.UUID, newEmail string) (*user.User, error) {
	query := "UPDATE users SET email=$1 WHERE id=$2 RETURNING *"

	var user user.User
	err := r.db.Get(&user, query, newEmail, id)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func (r *userPostgresRepository) Replenishment(id uuid.UUID, money float32) (*user.User, error) {
	query := "UPDATE users SET balance = balance + $1 WHERE id = $2 RETURNING *"

	var user user.User
	err := r.db.Get(&user, query, money, id)
	if err != nil {
		return nil, err
	}
	// TODO rowsAffected, err := res.RowsAffected()
	//if err != nil {
	//return err
	//}
	// if rowsAffected == 0 {
	// 	return ErrNoRowsUpdated
	// }
	return &user, nil

}

// TODO test
func (r *userPostgresRepository) GetHistoryOfQuest(id uuid.UUID) ([]quest.Quest, error) {
	query := "SELECT * FROM user_quest_history WHERE user_id = $1"

	var quests []quest.Quest
	rows, err := r.db.Query(query, id)
	if err != nil {
		return []quest.Quest{}, err
	}
	for rows.Next() {
		var q quest.Quest
		var s subquest.Subquest
		err = rows.Scan(&q.ID, &q.Title, &s.Description, &q.Status, &q.StartDate, &q.DueDate,
			&q.Repeatable, &q.Reward, &s.Title, &s.Description, &s.Status, &s.StartDate, &s.DueDate)
		if err != nil {
			return quests, err
		}

		q.Subquests = append(q.Subquests, s)

		quests = append(quests, q)
	}
	return quests, nil

}
