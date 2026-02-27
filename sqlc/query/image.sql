-- name: InsertImage :one
insert into character_image (owner_id, character_id, file_hash) values ($1, $2, $3) returning id;

-- name: GetImage :one
select * from character_image where owner_id = $1 and id = $2 limit 1; 

-- name: ListCharacterImages :many
select * from character_image where owner_id = $1 and character_id = $2;

-- name: GetMainImage :one
select * from character_image where character_id = $1 and is_main = true limit 1; 
