-- name: ListUserCharacters :many
select * from character where owner = $1 and not deleted;

-- name: GetCharacter :one
select * from character where id = $1 and not deleted limit 1;

-- name: DeleteCharacter :exec
update character set deleted = true where id = $1 and not deleted;

-- name: CreateCharacter :one
insert into character (owner, name, nationality, age, body, breast, butt, eyes_color, hair_style, hair_color, meta_params) values
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) returning id;