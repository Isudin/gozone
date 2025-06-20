-- name: CreateFaction :one
INSERT INTO factions (
   id, created_at, updated_at, name 
) VALUES ( gen_random_uuid(), NOW(), NOW(), $1 )
RETURNING *;

-- name: DeleteFaction :exec
DELETE FROM factions
  WHERE id = $1 CASCADE;

-- name: DeleteAllFactions :exec
DELETE FROM factions CASCADE;

-- name: GetFactionById :one
SELECT *
  FROM factions
  WHERE id = $1;

-- name: GetAllFactions :many
SELECT *
  FROM factions;
