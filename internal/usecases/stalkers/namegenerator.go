// Package stalkers is provides utility functions for manipulating stalker objects, which includes
// creating and modifying
package stalkers

import (
	"log"
	"math/rand/v2"

	domain "github.com/isudin/gozone/internal/domain/stalkers"
	infra "github.com/isudin/gozone/internal/infrastructure/stalkers"
)

func GenerateName(count int, factionName string) []domain.StalkerName {
	factionOpts, exists := FactionOpts[factionName]
	if !exists {
		log.Fatalf("Faction %s doesn't exist", factionName)
	}

	names := []domain.StalkerName{}
	for range count {
		name := buildName(factionOpts)
		names = append(names, name)
	}

	return names
}

func getRandomElement(slice []string) string {
	count := len(slice)
	if count == 0 {
		return ""
	}
	rnd := rand.IntN(count)
	return slice[rnd]
}

func buildName(opts factionOpts) domain.StalkerName {
	name := domain.StalkerName{}
	var firstNames, lastNames, monikers []string
	if checkChance(opts.maleChance) {
		firstNames = infra.FirstNamesMale
		lastNames = infra.LastNamesMale
	} else {
		firstNames = infra.FirstNamesFemale
		lastNames = infra.LastNamesFemale
	}

	if opts.isMonolith {
		monikers = infra.MonolithMonikers
	} else {
		monikers = infra.GeneralMonikers
	}

	if checkChance(opts.militaryTitleChance) {
		name.Title = getRandomElement(infra.MilitaryTitles)
	} else if checkChance(opts.scientificTitleChance) {
		name.Title = getRandomElement(infra.ScientificTitles)
	}

	if checkChance(opts.monikerChance) {
		name.Moniker = getRandomElement(monikers)
		if !checkChance(opts.onlyMonikerChance) {
			name.FirstName = getRandomElement(firstNames)
			if !checkChance(opts.monikerSkipLastNameChance) {
				name.LastName = getRandomElement(lastNames)
			}
		}
	}

	return name
}

func checkChance(chance float64) bool {
	return chance > 0 && rand.Float64() <= chance
}
