-- name: CreateaProject :one
INSERT INTO projects (title, status, possible_points, course_uuid)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CreateProjectParticipant :one
INSERT INTO project_participants (student_uid, project_uuid, points_earned)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetProjectByUUID :one
SELECT *
FROM projects
WHERE uuid = $1;

-- name: GetProjectParticipantByUUID :one
SELECT *
FROM project_participants
WHERE uuid = $1;

-- name: UpdateProject :one
UPDATE projects
SET title = $2,
    status = $3,
    possible_points = $4,
    course_uuid = $5
WHERE uuid = $1
RETURNING *;

-- name: UpdateProjectParticipant :one
UPDATE project_participants
SET student_uid = $2,
    project_uuid = $3,
    points_earned = $4
WHERE uuid = $1
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects
WHERE uuid = $1;

-- name: DeleteProjectParticipant :exec
DELETE FROM project_participants
WHERE uuid = $1;
