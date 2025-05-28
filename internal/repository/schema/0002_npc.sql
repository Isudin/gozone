-- +goose Up
CREATE TABLE npc (
  id UUID NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL,
  faction_id UUID NOT NULL,
  FOREIGN KEY(faction_id) REFERENCES factions(id)
);

-- +goose Down
DROP TABLE npc;
