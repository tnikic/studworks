-- name: CreateCourse :one
INSERT INTO courses (subject_uuid, teacher_uid, class_name)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateSubject :one
INSERT INTO subjects (name, program_code, semester)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetCourseByUUID :one
SELECT *
FROM courses
WHERE uuid = $1;

-- name: GetSubjectByUUID :one
SELECT *
FROM subjects
WHERE uuid = $1;

-- name: UpdateCourse :one
UPDATE courses
SET subject_uuid = $2,
    teacher_uid = $3,
    class_name = $4
WHERE uuid = $1
RETURNING *;

-- name: UpdateSubject :one
UPDATE subjects
SET name = $2,
    program_code = $3,
    semester = $4
WHERE uuid = $1
RETURNING *;

-- name: DeleteCourse :exec
DELETE FROM courses
WHERE uuid = $1;

-- name: DeleteSubject :exec
DELETE FROM subjects
WHERE uuid = $1;
