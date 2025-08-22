-- name: ListFeedsWithUsers :many
select
	f.name as feed_name,
	f.url as url,
	u.name as user_name
	from feeds as f
join users as u on u.id = f.user_id
order by f.created_at desc;

