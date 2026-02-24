-- name: CreateImage :one
insert into character_image (character_id, file_hash) values ($1, $2) returning id;

-- name: GetImage :one
select * from character_image where id = $1 limit 1; 

-- name: ListImages :many
select * from character_image where character_id = $1;

-- name: SetMainImage :exec
update character_image set is_main = true where character_id =$1 and id = $2;

-- name: GetMainImage :one
select * from character_image where character_id = $1 and is_main = true limit 1; 