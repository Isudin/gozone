-- name: CreateFaction :one
INSERT INTO factions (
   id, created_at, updated_at, name 
) VALUES ( gen_random_uuid(), NOW(), NOW(), $1 )
RETURNING *;
