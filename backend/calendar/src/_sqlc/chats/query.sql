-- name: GetCalendar :many
SELECT * FROM calendar
WHERE id = ? LIMIT 1;

-- name: ListCalendar :many
SELECT * FROM calendar
ORDER BY name;

-- name: CreateCalendar :execresult
INSERT INTO calendar (
  name, bio
) VALUES (
  ?, ?
);

-- name: UpdateCalendar :exec
UPDATE calendar
SET name = ?, bio = ?
WHERE id = ?;

-- name: DeleteCalendar :exec
DELETE FROM calendar
WHERE id = ?;