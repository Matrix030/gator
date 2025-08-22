-- name: CreateFeed :one
insert into feeds (id, created_at, updated_at, name, url, user_id)
values(gen_random_uuid(), now(), now(), $1, $2, $3)
returning *;
