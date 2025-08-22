-- name: CreateFeedFollow :one
WITH inserted AS (
  INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
  VALUES (gen_random_uuid(), now(), now(), $1, $2)
  RETURNING id, created_at, updated_at, user_id, feed_id
)
SELECT
  i.id, i.created_at, i.updated_at, i.user_id,
  u.name AS user_name,
  i.feed_id, f.name AS feed_name, f.url AS feed_url
FROM inserted i
JOIN users u ON u.id = i.user_id
JOIN feeds f ON f.id = i.feed_id;
