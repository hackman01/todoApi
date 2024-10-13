-- name: AddItem :one

INSERT INTO list (id,item,created_at,updated_at)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: GetItems :many

SELECT * FROM list;

-- name: UpdateItem :one

Update list 
SET item=$1, updated_at = NOW()
where id=$2
RETURNING *;

-- name: DeleteItem :one

DELETE FROM list
WHERE id=$1
RETURNING *;