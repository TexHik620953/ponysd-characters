-- name: ListUserCharacters :many
select * from character where owner_id = $1 and not deleted;

-- name: ListPublicCharacters :many
select * from character where public and not deleted limit $1 offset $2;

-- name: GetCharacter :one
select * from character 
where id = $1 
  and not deleted 
  and (
    public 
    or owner_id = $2
  )
limit 1;

-- name: DeleteCharacter :exec
update character set deleted = true where owner_id = $1 and id = $2 and not deleted;

-- name: CreateCharacter :one
insert into character (owner_id, name,biography, nationality, age, body, breast, butt, eyes_color, hair_style, hair_color, meta_params) values
    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) returning id;