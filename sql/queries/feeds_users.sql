-- name: CreateFeedUser :one
INSERT INTO feeds_users (
        id,
        feed_id,
        user_id,
        created_at,
        updated_at
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: DeleteFeedUser :one
DELETE FROM feeds_users
WHERE id = $1
RETURNING *;
-- name: GetFeedsUser :many
SELECT *
FROM feeds_users
WHERE user_id = $1;