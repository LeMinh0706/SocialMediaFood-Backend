-- name: AddToMenu :one
INSERT INTO menu (
    account_id,
    dish_name,
    quantity,
    price,
    img
)VALUES(
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetDish :one
SELECT id, dish_name, quantity, price, img 
FROM menu
WHERE account_id = $1
LIMIT 1;

-- name: GetMenu :many
SELECT id, dish_name, quantity, price, img 
FROM menu
WHERE account_id = $1
LIMIT $2
OFFSET $3;

-- name: UpdateQuanity :one
UPDATE menu
SET quantity = quantity + sqlc.arg(quantity)
WHERE id = $1
RETURNING *;

-- name: OwnerUpdateQuanity :one
UPDATE menu
SET quantity = $2
WHERE id = $1
RETURNING *;