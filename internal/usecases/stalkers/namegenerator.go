package stalkers

import (
	"math/rand/v2"

	domain "github.com/isudin/gozone/internal/domain/stalkers"
	infra "github.com/isudin/gozone/internal/infrastructure/stalkers"
)

func GenerateName(count int, isMonolith bool) []domain.StalkerName {
	names := []domain.StalkerName{}
	for range count {
		name := domain.StalkerName{
			FirstName: getRandomElement(infra.FirstNamesMale),
			LastName:  getRandomElement(infra.LastNamesMale),
			Moniker:   getRandomElement(infra.GeneralMonikers),
			Title:     getRandomElement(infra.MilitaryTitles),
		}

		names = append(names, name)
	}

	return names
}

func getRandomElement(slice []string) string {
	count := len(slice)
	rnd := rand.IntN(count)
	return slice[rnd]
}
