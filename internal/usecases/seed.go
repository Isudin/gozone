package usecases

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/isudin/gozone/internal/repository/sqlc"
)

func InitDatabase(queries *sqlc.Queries) error {
	err := clearFactions(queries)
	if err != nil {
		return err
	}

	factions, err := initFactions(queries)
	if err != nil {
		return err
	}

	err = initNpc(queries, factions)
	return err
}

func clearFactions(queries *sqlc.Queries) error {
	err := queries.DeleteAllFactions(context.Background())
	if err != nil {
		fmt.Printf("could not clear `factions` table: %s\n", err)
	}

	return err
}

var factionNames = []string{
	"Duty",
	"Freedom",
	"Loners",
	"Bandits",
	"Military",
	"Mercenaries",
	"Ecologists",
	"Monolith",
	"Clear Sky",
	"Renegades",
	"Noontide",
	"Spark",
	"Ward",
	"Corps",
}

func initFactions(queries *sqlc.Queries) (map[string]uuid.UUID, error) {
	factions := make(map[string]uuid.UUID)

	for _, name := range factionNames {
		faction, err := queries.CreateFaction(context.Background(), name)
		if err != nil {
			fmt.Printf("could not initialize `factions` table: %s\n", err)
			return nil, err
		}

		factions[faction.Name] = faction.ID
	}

	return factions, nil
}

func initNpc(queries *sqlc.Queries, factions map[string]uuid.UUID) error {
	npcSlice := []sqlc.CreateNpcParams{
		{Name: "Duty Member", FactionID: factions["Duty"]},
		{Name: "Freedom Member", FactionID: factions["Freedom"]},
		{Name: "The Loner", FactionID: factions["Loners"]},
		{Name: "Another Duty Member", FactionID: factions["Duty"]},
	}

	for _, npc := range npcSlice {
		_, err := queries.CreateNpc(context.Background(), npc)
		if err != nil {
			fmt.Printf("could not initialize `npc` table: %s\n", err)
			return err
		}
	}

	return nil
}
