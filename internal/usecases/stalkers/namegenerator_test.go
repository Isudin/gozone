package stalkers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandomElement(t *testing.T) {
	cases := []struct {
		name          string
		slice         []string
		emptyExpected bool
	}{
		{"empty slice", []string{}, true},
		{"nil slice", nil, true},
		{"regular", []string{"test1", "test2", "test3"}, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Logf("Checking '%v'", c.name)
			o := getRandomElement(c.slice)
			if c.emptyExpected {
				assert.Empty(t, o, "Returned value should be empty")
			} else {
				assert.Contains(t, c.slice, o, "Slice should contain given value")
			}
		})
	}
}
