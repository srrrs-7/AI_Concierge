-- name: Getusers :many
SELECT * FROM users
WHERE id = ? LIMIT 1;

-- name: Listusers :many
SELECT * FROM users
ORDER BY name;

-- name: Createusers :execresult
INSERT INTO users (
  name, bio
) VALUES (
  ?, ?
);

-- name: Updateusers :exec
UPDATE users
SET name = ?, bio = ?
WHERE id = ?;

-- name: Deleteusers :exec
DELETE FROM users
WHERE id = ?;