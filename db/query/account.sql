-- name: CreateAccount :one
INSERT INTO accounts (owner, balance, currency)
    VALUES (sqlc.arg (owner), sqlc.arg (balance), sqlc.arg (currency))
RETURNING
    *;

-- name: GetAccount :one
SELECT
    *
FROM
    accounts
WHERE
    id = sqlc.arg (id)
LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT
    *
FROM
    accounts
WHERE
    id = sqlc.arg (id)
LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateAccount :one
UPDATE
    accounts
SET
    balance = sqlc.arg (balance)
WHERE
    id = sqlc.arg (id)
RETURNING
    *;

-- name: AddAccountBalance :one
UPDATE
    accounts
SET
    balance = balance + sqlc.arg (amount)
WHERE
    id = sqlc.arg (id)
RETURNING
    *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = sqlc.arg (id);

