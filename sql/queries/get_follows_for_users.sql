-- name: GetFeedFollowsForUser :many
SELECT
  ff.id,
  ff.created_at,
  ff.updated_at,
  ff.user_id,
  u.name  AS user_name,
  ff.feed_id,
  f.name  AS feed_name,
  f.url   AS feed_url
FROM feed_follows ff
JOIN users u ON u.id = ff.user_id
JOIN feeds f ON f.id = ff.feed_id
WHERE ff.user_id = $1
ORDER BY ff.created_at DESC;

