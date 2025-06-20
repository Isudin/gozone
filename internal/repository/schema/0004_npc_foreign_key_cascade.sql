-- +goose Up
ALTER TABLE Npc
  DROP CONSTRAINT npc_faction_id_fkey;

ALTER TABLE Npc
  ADD CONSTRAINT npc_faction_id_fkey
  FOREIGN KEY(faction_id) REFERENCES factions(id) ON DELETE CASCADE;
  
-- +goose Down
ALTER TABLE Npc
  DROP CONSTRAINT npc_faction_id_fkey;

ALTER TABLE Npc
  ADD CONSTRAINT npc_faction_id_fkey
  FOREIGN KEY(faction_id) REFERENCES factions(id);
