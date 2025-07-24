-- name: CreateClass :one
INSERT INTO classes (name, program_code, year, study_type, active)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetClassByName :one
SELECT *
FROM classes 
WHERE name = $1;

-- name: UpdateClass :one
UPDATE classes
SET program_code = $2,
    year = $3,
    study_type = $4,
    active = $5
WHERE name = $1
RETURNING *;

-- name: DeleteClass :exec
DELETE FROM classes
WHERE name = $1;
