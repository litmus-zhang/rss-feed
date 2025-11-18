-- name: CreateFeed :one
INSERT INTO feeds (feed_name, url, description) VALUES ($1, $2, $3)
RETURNING *;




-- name: UpdateFeed :one
UPDATE feeds SET  feed_name = $2, url=$3, description=$4 WHERE feed_id = $1
RETURNING *;


-- name: GetOneFeedById :one
SELECT * FROM feeds WHERE feed_id = $1;



-- name: DeleteFeed :exec
DELETE FROM feeds WHERE feed_id = $1;


-- name: GetAllFeeds :many
SELECT * FROM feeds
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;