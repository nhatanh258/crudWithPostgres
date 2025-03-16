-- name: RegisterUserForEvent :one
INSERT INTO register (userId, eventid)
VALUES ($1, $2)
RETURNING *;

-- name: GetRegistrationsByUserID :many
SELECT * FROM register WHERE userid = $1;

-- name: GetRegistrationsByEventID :many
SELECT * FROM register WHERE eventid = $1;
