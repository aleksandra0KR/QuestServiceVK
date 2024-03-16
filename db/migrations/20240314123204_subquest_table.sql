-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS subquest (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    status VARCHAR(255) NOT NULL,
    startDate timestamp with time zone NOT NULL,
    dueDate timestamp with time zone NOT NULL
    );

CREATE TABLE IF NOT EXISTS subquestOfQuest (
    subquest_id UUID REFERENCES subquest(id) ON DELETE CASCADE,
    quest_id UUID REFERENCES quest(id) ON DELETE CASCADE
 );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS subquest;
DROP TABLE IF EXISTS subquestOfQuest
-- +goose StatementEnd

// export GOOSE_DBSTRING=postgresql://postgres:FuFa2020@127.0.0.1:5432/Quest?sslmode=disable
