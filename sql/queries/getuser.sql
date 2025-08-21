-- name: GetUser :one
select *
from users
where name = $1
limit 1;

-- name: GetUsers :many
select *
from users;
