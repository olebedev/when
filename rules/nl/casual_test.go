package nl_test

import (
	"github.com/olebedev/when/rules/nl"
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
)

func TestCasualDate(t *testing.T) {
	fixt := []Fixture{
		{"De deadline is nu, ok", 15, "nu", 0},
		{"De deadline is vandaag", 15, "vandaag", 0},
		{"De deadline is vannacht", 15, "vannacht", 23 * time.Hour},
		{"De deadline is morgenavond", 15, "morgenavond", (18 + 24) * time.Hour},
		{"De deadline is gisteravond", 15, "gisteravond", -((24 - 18) * time.Hour)},
		{"De deadline is gisteren", 15, "gisteren", -(time.Hour * 24)},
	}

	w := when.New(nil)
	w.Add(nl.CasualDate(rules.Skip))

	ApplyFixtures(t, "nl.CasualDate", w, fixt)
}

func TestCasualTime(t *testing.T) {
	fixt := []Fixture{
		{"De deadline was deze morgen", 16, "deze morgen", 8 * time.Hour},
		{"De deadline was tussen de middag", 16, "tussen de middag", 12 * time.Hour},
		{"De deadline was deze middag", 16, "deze middag", 15 * time.Hour},
		{"De deadline was deze avond", 16, "deze avond", 18 * time.Hour},
		{"De deadline is donderdagavond", 15, "donderdagavond", (18 + 24) * time.Hour},
		{"De deadline is vrijdagavond", 15, "vrijdagavond", (18 + 24*2) * time.Hour},
	}

	w := when.New(nil)
	w.Add(nl.CasualTime(rules.Skip))

	ApplyFixtures(t, "nl.CasualTime", w, fixt)
}

func TestCasualDateCasualTime(t *testing.T) {
	fixt := []Fixture{
		{"De deadline is morgenmiddag", 15, "morgenmiddag", (15 + 24) * time.Hour},
		{"De deadline is morgenavond", 15, "morgenavond", (18 + 24) * time.Hour},
	}

	w := when.New(nil)
	w.Add(
		nl.CasualDate(rules.Skip),
		nl.CasualTime(rules.Override),
	)

	ApplyFixtures(t, "nl.CasualDate|nl.CasualTime", w, fixt)
}
