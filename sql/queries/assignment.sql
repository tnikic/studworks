-- name: CreateAssignment :one
INSERT INTO assignments (assignment_template_uuid, student_uid, course_uuid, possible_points, points_earned)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: CreateAssignmentTemplate :one
INSERT INTO assignment_templates (title)
VALUES ($1)
RETURNING *;

-- name: GetAssignmentByUUID :one
SELECT *
FROM assignments
WHERE uuid = $1;

-- name: GetAssignmentTemplateByUUID :one
SELECT *
FROM assignment_templates
WHERE uuid = $1;

-- name: UpdateAssignment :one
UPDATE assignments
SET assignment_template_uuid = $2,
    student_uid = $3,
    course_uuid = $4,
    possible_points = $5,
    points_earned = $6
WHERE uuid = $1
RETURNING *;

-- name: UpdateAssignmentTemplate :one
UPDATE assignment_templates
SET title = $2
WHERE uuid = $1
RETURNING *;

-- name: DeleteAssignment :exec
DELETE FROM assignments
WHERE uuid = $1;

-- name: DeleteAssignmentTemplate :exec
DELETE FROM assignment_templates
WHERE uuid = $1;
