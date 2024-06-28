-- +goose StatementBegin
-- +goose Up
SELECT 'up SQL query';
CREATE TABLE "patterns" (
  "user_id" uuid NOT NULL,
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "instructions" text
);

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  "username" VARCHAR(50) NOT NULL,
  "email" VARCHAR(150) UNIQUE NOT NULL CHECK (email ~* '^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,4}$'),
  "hashed_password" VARCHAR(1024) NOT NULL,
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
