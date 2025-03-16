-- name: CreateEvent :one
INSERT INTO events (name, description, location, "userId")
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetEventByID :one
SELECT * FROM events WHERE id = $1;

-- name: ListEvents :many
SELECT * FROM events ORDER BY datatimecreating DESC;

-- name: UpdateEvent :one
UPDATE events 
SET name = $2, description = $3, location = $4, datetimeupdating = now()
WHERE id = $1
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM events WHERE id = $1;
