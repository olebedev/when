package nl_test

import (
	"github.com/olebedev/when/rules/nl"
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
)

func TestWeekday(t *testing.T) {
	// current is Friday
	fixt := []Fixture{
		// past/last
		{"doe het voor afgelopen maandag", 13, "afgelopen maandag", -(2 * 24 * time.Hour)},
		{"afgelopen zaterdag", 0, "afgelopen zaterdag", -(4 * 24 * time.Hour)},
		{"afgelopen vrijdag", 0, "afgelopen vrijdag", -(5 * 24 * time.Hour)},
		{"afgelopen woensdag", 0, "afgelopen woensdag", -(7 * 24 * time.Hour)},
		{"afgelopen dinsdag", 0, "afgelopen dinsdag", -(24 * time.Hour)},
		// next
		{"komende dinsdag", 0, "komende dinsdag", 6 * 24 * time.Hour},
		{"stuur me een bericht komende woensdag", 21, "komende woensdag", 7 * 24 * time.Hour},
		{"komende zaterdag", 0, "komende zaterdag", 3 * 24 * time.Hour},
		{"volgende dinsdag", 0, "volgende dinsdag", 6 * 24 * time.Hour},
		{"stuur me een bericht volgende woensdag", 21, "volgende woensdag", 7 * 24 * time.Hour},
		{"volgende zaterdag", 0, "volgende zaterdag", 3 * 24 * time.Hour},
		// this
		{"deze dinsdag", 0, "deze dinsdag", -(24 * time.Hour)},
		{"stuur me een bericht deze woensdag", 21, "deze woensdag", 0},
		{"deze zaterdag", 0, "deze zaterdag", 3 * 24 * time.Hour},
	}

	w := when.New(nil)

	w.Add(nl.Weekday(rules.Override))

	ApplyFixtures(t, "nl.Weekday", w, fixt)
}
