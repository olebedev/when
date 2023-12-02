package nl_test

import (
	"testing"
	"time"

	"github.com/olebedev/when/rules/nl"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
)

func TestWeekday(t *testing.T) {
	// current is Wednesday
	fixt := []Fixture{
		// past week
		{"vorige week maandag", 0, "vorige week maandag", -(9 * 24 * time.Hour)},
		{"vorige week dinsdag", 0, "vorige week dinsdag", -(8 * 24 * time.Hour)},
		{"vorige week woensdag", 0, "vorige week woensdag", -(7 * 24 * time.Hour)},
		{"vorige week donderdag", 0, "vorige week donderdag", -(6 * 24 * time.Hour)},
		{"vorige week vrijdag", 0, "vorige week vrijdag", -(5 * 24 * time.Hour)},
		{"vorige week zaterdag", 0, "vorige week zaterdag", -(4 * 24 * time.Hour)},
		{"vorige week zondag", 0, "vorige week zondag", -(3 * 24 * time.Hour)},
		// past/last
		{"doe het voor afgelopen maandag", 13, "afgelopen maandag", -(2 * 24 * time.Hour)},
		{"afgelopen zaterdag", 0, "afgelopen zaterdag", -(4 * 24 * time.Hour)},
		{"afgelopen vrijdag", 0, "afgelopen vrijdag", -(5 * 24 * time.Hour)},
		{"afgelopen woensdag", 0, "afgelopen woensdag", -(7 * 24 * time.Hour)},
		{"afgelopen dinsdag", 0, "afgelopen dinsdag", -(24 * time.Hour)},
		// next week
		{"volgende week maandag", 0, "volgende week maandag", 5 * 24 * time.Hour},
		{"volgende week dinsdag", 0, "volgende week dinsdag", 6 * 24 * time.Hour},
		{"volgende week woensdag", 0, "volgende week woensdag", 7 * 24 * time.Hour},
		{"volgende week donderdag", 0, "volgende week donderdag", 8 * 24 * time.Hour},
		{"volgende week vrijdag", 0, "volgende week vrijdag", 9 * 24 * time.Hour},
		{"volgende week zaterdag", 0, "volgende week zaterdag", 10 * 24 * time.Hour},
		{"volgende week zondag", 0, "volgende week zondag", 11 * 24 * time.Hour},
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
