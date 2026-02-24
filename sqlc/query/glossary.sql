-- name: ListGlossaryByType :many
select * from glossary where type = $1;


-- name: ListGlossaryByTypeLocal :many
select glossary.type, glossary.value, gl.name from glossary join public.glossary_localization gl on glossary.id = gl.glossary_id where glossary.type = $1 and gl.local = $2;