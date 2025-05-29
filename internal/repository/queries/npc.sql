-- name: CreateNpc :one
INSERT INTO npc (
  id, created_at, updated_at, name, faction_id
) VALUES ( gen_random_uuid(), NOW(), NOW(), $1, $2)
RETURNING *;
