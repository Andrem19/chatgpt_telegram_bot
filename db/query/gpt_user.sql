-- name: CreateUser :one
INSERT INTO gpt_user (
  chat_id,
  gpt_token
) VALUES (
  $1, $2
) RETURNING id;

-- name: GetUsers :one
SELECT * FROM gpt_user
WHERE chat_id = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM gpt_user
WHERE chat_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListUsers :many
SELECT * FROM gpt_user
ORDER BY id;

-- name: UpdateUserToken :one
UPDATE gpt_user
SET gpt_token = $2
WHERE chat_id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM gpt_user
WHERE chat_id = $1;