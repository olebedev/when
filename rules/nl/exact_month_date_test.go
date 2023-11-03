package nl_test

import (
	"github.com/olebedev/when/rules/nl"
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
)

func TestExactMonthDate(t *testing.T) {
	w := when.New(nil)
	w.Add(nl.ExactMonthDate(rules.Override))

	fixtok := []Fixture{
		{"derde van maart", 0, "derde van maart", 1368 * time.Hour},
		{"3e van maart", 0, "3e van maart", 1368 * time.Hour},
		{"1 september", 0, "1 september", 5736 * time.Hour},
		{"1 sept", 0, "1 sept", 5736 * time.Hour},
		{"1 sept.", 0, "1 sept.", 5736 * time.Hour},
		{"1e van september", 0, "1e van september", 5736 * time.Hour},
		{"twintigste van december", 0, "twintigste van december", 8376 * time.Hour},
		{"februari", 0, "februari", 744 * time.Hour},
		{"oktober", 0, "oktober", 6576 * time.Hour},
		{"jul.", 0, "jul.", 4368 * time.Hour},
		{"juni", 0, "juni", 3648 * time.Hour},
	}

	ApplyFixtures(t, "nl.ExactMonthDate", w, fixtok)
}
