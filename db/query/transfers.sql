-- name: CreateTransfer :one
INSERT INTO transfers (from_account_id, to_account_id, amount)
    VALUES (sqlc.arg (from_account_id), sqlc.arg (to_account_id), sqlc.arg (amount))
RETURNING
    *;

-- name: GetTransfer :one
SELECT
    *
FROM
    transfers
WHERE
    id = sqlc.arg (id)
LIMIT 1;

-- name: ListTransfers :many
SELECT
    *
FROM
    transfers
ORDER BY
    id
LIMIT $1 OFFSET $2;

-- name: UpdateTransfer :one
UPDATE
    transfers
SET
    amount = sqlc.arg (amount)
WHERE
    id = sqlc.arg (id)
RETURNING
    *;

-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = sqlc.arg (id);

