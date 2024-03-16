-- +goose Up
-- +goose StatementBegin
CREATE VIEW IF NOT EXISTS user_quest_history AS
SELECT u.id as user_id, q.id AS quest_id, sq.id AS subquest_id, q.title AS quest_title, sq.title AS subquest_title,
       q.status AS quest_status, sq.status AS subquest_status, q.description as quest_description,
       q.startdate AS quest_start, q.duedate AS quest_duedate,
       q.repeatable AS quest_repeatable, q.reward AS quest_reward,
       sq.description AS subquest_description,
       sq.startdate AS subquest_start, sq.duedate AS subquest_duedate
FROM usersquests AS uq
         JOIN users u ON u.id = uq.user_id
         JOIN quest q ON q.id = uq.quest_id
         LEFT JOIN subquestOfQuest soc ON q.id = soc.quest_id
         LEFT JOIN subquest sq ON soc.subquest_id = sq.id
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW IF EXISTS user_quest_history;
-- +goose StatementEnd
