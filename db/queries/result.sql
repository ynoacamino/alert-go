-- name: ListAllResults :many
SELECT
    *
FROM
    result
ORDER BY
    updatedAt DESC;

-- name: ListResults :many
SELECT
    *
FROM
    result
ORDER BY
    updatedAt DESC
LIMIT
    ?
OFFSET
    ?;

-- name: CreateResult :one
INSERT INTO
    result (id, status)
VALUES
    (?, ?) RETURNING *;

-- name: GetResult :one
SELECT
    *
FROM
    result
WHERE
    id = ?;

-- name: UpdateResult :exec
UPDATE result
SET
    status = ?,
    updatedAt = CURRENT_TIMESTAMP
WHERE
    id = ?;

-- name: DeleteResult :exec
DELETE FROM result
WHERE
    id = ?;
