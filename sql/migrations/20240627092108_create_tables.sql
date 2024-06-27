-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE "patterns" (
  "user_id" uuid,
  "id" uuid PRIMARY KEY,
  "instructions" text
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY,
  "username" varchar,
  "email" varchar,
  "hashed_password" varchar,
  "salt" varchar
);

ALTER TABLE "patterns" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE "patterns";
DROP TABLE "users";
-- +goose StatementEnd
