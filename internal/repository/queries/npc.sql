-- name: CreateNpc :one
INSERT INTO npc (
  id, created_at, updated_at, name, faction_id
) VALUES ( gen_random_uuid(), NOW(), NOW(), $1, $2)
RETURNING *;

-- name: DeleteNpc :exec
DELETE FROM npc
  WHERE id = $1;

-- name: DeleteAllNpc :exec
DELETE FROM npc;

-- name: GetNpcById :one
SELECT *
  FROM npc
  WHERE id = $1;

-- name: GetNpcByFactionId :many
SELECT *
  FROM npc
  WHERE faction_id = $1;

-- name: CountNpc :one
SELECT COUNT(*) FROM npc;
