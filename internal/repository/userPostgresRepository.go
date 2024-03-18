package repository

import (
	"errors"
	"github.com/gofrs/uuid/v5"
	"github.com/jmoiron/sqlx"
	"vk/internal/model/quest"
	"vk/internal/model/status"
	"vk/internal/model/subquest"
	"vk/internal/model/user"
	"vk/internal/model/usershistory"
)

type userPostgresRepository struct {
	db *sqlx.DB
}

func NewUserPostgresRepository(db *sqlx.DB) *userPostgresRepository {
	return &userPostgresRepository{db: db}
}

func (r *userPostgresRepository) CreateUser(user *user.User) error {
	var id uuid.UUID
	query := `INSERT INTO "users" ("username", "email", "password", "balance") VALUES ($1, $2, $3, $4) RETURNING id`

	row := r.db.QueryRow(query, user.Username, user.Email, user.Password, user.Balance)
	if err := row.Scan(&id); err != nil {
		return err
	}
	user.ID = id

	return nil
}

func (r *userPostgresRepository) DeleteUserByID(id uuid.UUID) error {

	query := `DELETE FROM "users" where "id"= $1`

	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userPostgresRepository) FindUserByID(id uuid.UUID) (*user.User, error) {
	query := `SELECT "id", "username", "email", "password", "balance" FROM "users" WHERE "id" = $1`

	var user user.User
	err := r.db.Get(&user, query, id)

	return &user, err
}

func (r *userPostgresRepository) UpdateName(id uuid.UUID, newName string) error {
	query, err := r.db.Prepare(`UPDATE "users" SET "username"=$1 WHERE "id"=$2`)
	if err != nil {
		return err
	}
	_, err = query.Exec(newName, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userPostgresRepository) UpdatePassword(id uuid.UUID, newPassword string) error {
	query, err := r.db.Prepare(`UPDATE "users" SET "password"=$1 WHERE "id"=$2`)
	if err != nil {
		return err
	}
	_, err = query.Exec(newPassword, id)

	if err != nil {
		return err
	}
	return nil
}

func (r *userPostgresRepository) UpdateEmail(id uuid.UUID, newEmail string) error {

	query, err := r.db.Prepare(`UPDATE "users" SET "email"=$1 WHERE "id"=$2`)
	if err != nil {
		return err
	}
	_, err = query.Exec(newEmail, id)

	if err != nil {
		return err
	}
	return nil
}

func (r *userPostgresRepository) Replenishment(id uuid.UUID, money float32) (*user.User, error) {
	query := `UPDATE "users" SET "balance" = "balance" + $1 WHERE "id" = $2 RETURNING *`

	var user user.User
	err := r.db.Get(&user, query, money, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userPostgresRepository) GetHistoryOfQuest(id uuid.UUID) (usershistory.UsersHistory, error) {
	query := `SELECT "quest_id", "subquest_id", "quest_title", "subquest_title", "quest_status", "subquest_status", 
       		  "quest_description", "quest_start", "quest_duedate", "quest_repeatable", "quest_reward", 
       		"subquest_description","subquest_start", "subquest_duedate"
			FROM "user_quest_history" WHERE "user_id" = $1`

	var history usershistory.UsersHistory
	rows, err := r.db.Query(query, id)
	if err != nil {
		return history, err
	}

	for rows.Next() {
		var q quest.Quest
		var s subquest.Subquest
		err = rows.Scan(&q.ID, &s.ID, &q.Title, &s.Title, &q.Status, &s.Status,
			&q.Description, &q.StartDate, &q.DueDate, &q.Repeatable, &q.Reward,
			&s.Description, &s.StartDate, &s.DueDate)

		if err != nil {
			return history, err
		}

		q.Subquests = append(q.Subquests, s)
		history.Quest = append(history.Quest, q)

	}
	return history, nil
}

func (r *userPostgresRepository) AttachToQuest(questId uuid.UUID, userId uuid.UUID) error {
	stmt, err := r.db.Prepare(`INSERT INTO "usersQuests" ("user_id", "quest_id", "status") VALUES ($1, $2, $3)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(questId, userId, status.InProcess)
	if err != nil {
		return err
	}
	return nil
}

func (r *userPostgresRepository) AttachToSubquest(subquestId uuid.UUID, userId uuid.UUID) error {
	stmt, err := r.db.Prepare(`INSERT INTO "usersSubquests" ("user_id", "subquest_id", "status") VALUES ($1, $2, $3)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(subquestId, userId, status.InProcess)
	if err != nil {
		return err
	}

	return nil
}

func (r *userPostgresRepository) ChangeSubquestsStatus(subquestId uuid.UUID, userId uuid.UUID, givenStatus status.Status) error {

	query := `SELECT "status" FROM "usersSubquests" WHERE "user_id"=$1 and "subquest_id"=$2`

	var previousStatus status.Status
	err := r.db.Get(&previousStatus, query, subquestId, userId)

	if err != nil {
		return err
	}
	if previousStatus == status.Done {
		return errors.New("subquest can't be done one more time")
	}

	stmt, err2 := r.db.Prepare(`UPDATE "usersSubquests" SET "status" =$1 WHERE  "user_id"=$2 AND "subquest_id"=$3`)
	if err2 != nil {
		return err
	}
	_, err = stmt.Exec(givenStatus, subquestId, userId)

	if err != nil {
		return err
	}
	return nil
}

func (r *userPostgresRepository) ChangeQuestsStatus(questId uuid.UUID, userId uuid.UUID, givenStatus status.Status) error {

	query := `SELECT "status" FROM "usersQuests" WHERE "user_id"=$1 and "quest_id"=$2`

	var previousStatus status.Status
	err := r.db.Get(&previousStatus, query, questId, userId)

	if err != nil {
		return err
	}
	if previousStatus == status.Done {
		return errors.New("quest can't be done one more time")
	}

	stmt, err2 := r.db.Prepare(`UPDATE "usersQuests" SET "status" =$1 WHERE  "user_id"=$2 AND "quest_id"=$3`)
	if err2 != nil {
		return err
	}
	_, err = stmt.Exec(givenStatus, questId, userId)

	if err != nil {
		return err
	}

	if previousStatus != status.Done && givenStatus == status.Done {
		query = `SELECT "reward" FROM "quest" WHERE "id" = $1`

		var reward float32
		err = r.db.Get(&reward, query, userId)
		_, err = r.Replenishment(questId, reward)
		if err != nil {
			return err
		}
	}
	return nil
}
