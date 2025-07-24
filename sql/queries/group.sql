-- name: CreateGroup :one
INSERT INTO groups (id, course_uuid)
VALUES ($1, $2)
RETURNING *;

-- name: GetGroupByUUID :one
SELECT *
FROM groups 
WHERE uuid = $1;

-- name: UpdateGroup :one
UPDATE groups
SET id = $2,
    course_uuid = $3
WHERE uuid = $1
RETURNING *;

-- name: DeleteGroup :exec
DELETE FROM groups
WHERE uuid = $1;
