-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  username, email, hashed_password, salt
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;
