package stalkers

type factionOpts struct {
	militaryTitleChance       float64
	scientificTitleChance     float64
	maleChance                float64
	monikerChance             float64
	monikerSkipLastNameChance float64
	onlyMonikerChance         float64
	isMonolith                bool
}

var dutyOpts = factionOpts{
	militaryTitleChance:       100,
	scientificTitleChance:     0,
	maleChance:                95,
	monikerChance:             35,
	monikerSkipLastNameChance: 5,
	onlyMonikerChance:         0.1,
	isMonolith:                false,
}

var freeStalkers = factionOpts{
	militaryTitleChance:       0,
	scientificTitleChance:     0,
	maleChance:                90,
	monikerChance:             75,
	monikerSkipLastNameChance: 20,
	onlyMonikerChance:         1,
	isMonolith:                false,
}

var ecologists = factionOpts{
	scientificTitleChance: 100,
	maleChance:            70,
	monikerChance:         0.2,
}

var monolith = factionOpts{
	maleChance:        98,
	monikerChance:     100,
	onlyMonikerChance: 100,
	isMonolith:        true,
}
