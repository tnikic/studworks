-- name: CreateTeacher :one
INSERT INTO teachers (uid, first_name, last_name, email, active)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTeacherByUID :one
SELECT *
FROM teachers 
WHERE uid = $1;

-- name: UpdateTeacher :one
UPDATE teachers
SET first_name = $2,
    last_name = $3,
    email = $4,
    active = $5
WHERE uid = $1
RETURNING *;

-- name: DeleteTeacher :exec
DELETE FROM teachers
WHERE uid = $1;
