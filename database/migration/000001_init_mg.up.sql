CREATE TYPE status AS ENUM ('done', 'in process', 'available');

CREATE TABLE IF NOT EXISTS "users" (
    "id" uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    "username" VARCHAR(255) NOT NULL,
    "email" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "balance" integer NOT NULL
);

CREATE TABLE IF NOT EXISTS "quest" (
    "id" uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "status" status NOT NULL,
    "startDate" timestamp with time zone NOT NULL,
    "dueDate" timestamp with time zone NOT NULL,
    "repeatable" boolean NOT NULL,
    "reward" float(32) NOT NULL
);


CREATE TABLE IF NOT EXISTS "usersQuests" (
    "user_id" UUID REFERENCES "users"(id) ON DELETE CASCADE,
    "quest_id" UUID REFERENCES "quest"(id) ON DELETE CASCADE,
    "status" status NOT NULL
);


CREATE TABLE IF NOT EXISTS "subquest" (
    "id" uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    "title" VARCHAR(255) NOT NULL,
    "description" VARCHAR(255) NOT NULL,
    "status" status NOT NULL,
    "startDate" timestamp with time zone NOT NULL,
    "dueDate" timestamp with time zone NOT NULL
);

CREATE TABLE IF NOT EXISTS "subquestOfQuest" (
    "subquest_id" UUID REFERENCES "subquest"(id) ON DELETE CASCADE,
    "quest_id" UUID REFERENCES "quest"(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "usersSubquests" (
    "user_id" UUID REFERENCES "users"(id) ON DELETE CASCADE,
    "subquest_id" UUID REFERENCES "subquest"(id) ON DELETE CASCADE,
    "status" status NOT NULL
);

CREATE VIEW  "user_quest_history" AS
SELECT u."id" as "user_id", q."id" AS "quest_id", sq."id" AS "subquest_id", q."title" AS "quest_title", sq."title" AS "subquest_title",
       q."status" AS "quest_status", sq."status" AS "subquest_status", q."description" as "quest_description",
       q."startDate" AS "quest_start", q."dueDate" AS "quest_duedate",
       q."repeatable" AS "quest_repeatable", q."reward" AS "quest_reward",
       sq."description" AS "subquest_description",
       sq."startDate" AS "subquest_start", sq."dueDate" AS "subquest_duedate"
FROM "usersQuests" AS uq
    JOIN "users" AS u ON u."id" = uq."user_id"
    JOIN "quest" AS q ON q."id" = uq."quest_id"
    LEFT JOIN "subquestOfQuest" AS soc ON q."id" = soc."quest_id"
    LEFT JOIN "subquest" AS sq ON soc."subquest_id" = sq."id"