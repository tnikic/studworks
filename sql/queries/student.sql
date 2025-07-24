-- name: CreateStudent :one
INSERT INTO students (uid, first_name, last_name, email, active, class_name)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetStudentByUID :one
SELECT *
FROM students 
WHERE uid = $1;

-- name: UpdateStudent :one
UPDATE students
SET first_name = $2,
    last_name = $3,
    email = $4,
    active = $5,
    class_name = $6
WHERE uid = $1
RETURNING *;

-- name: DeleteStudent :exec
DELETE FROM students
WHERE uid = $1;
