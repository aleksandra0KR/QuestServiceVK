-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS quest (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    startDate timestamp with time zone NOT NULL,
    dueDate timestamp with time zone NOT NULL,
    repeatable boolean NOT NULL,
    Reward float(32) NOT NULL
    );

CREATE TABLE IF NOT EXISTS usersQuests (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    quest_id UUID REFERENCES quest(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS quest;
DROP TABLE IF EXISTS usersQuests;
-- +goose StatementEnd
