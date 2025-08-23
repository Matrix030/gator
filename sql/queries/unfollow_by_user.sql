-- name: UnfollowFeedByUserAndUrl :one
delete from feed_follows as ff
using feeds as f
where ff.feed_id = f.id
	and ff.user_id = $1
	and f.url = $2
returning ff.id;

