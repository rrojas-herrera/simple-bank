-- name: CreateEntry :one
INSERT INTO entries (account_id, amount)
    VALUES (sqlc.arg (account_id), sqlc.arg (amount))
RETURNING
    *;

-- name: GetEntry :one
SELECT
    *
FROM
    entries
WHERE
    id = sqlc.arg (id)
LIMIT 1;

-- name: ListEntries :many
SELECT
    *
FROM
    entries
ORDER BY
    id
LIMIT sqlc.arg ( LIMIT)
    OFFSET sqlc.arg (OFFSET);

-- name: UpdateEntry :one
UPDATE
    entries
SET
    amount = sqlc.arg (amount)
WHERE
    id = sqlc.arg (id)
RETURNING
    *;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = sqlc.arg (id);

