-- +goose Up
ALTER TABLE npc
ADD PRIMARY KEY(id);

-- +goose Down
ALTER TABLE npc
DROP CONSTRAINT npc_pkey;
