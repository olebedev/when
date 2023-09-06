package nl_test

import (
	"github.com/olebedev/when/rules/nl"
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
)

func TestDeadline(t *testing.T) {
	fixt := []Fixture{
		{"binnen een half uur", 0, "binnen een half uur", time.Hour / 2},
		{"binnen 1 uur", 0, "binnen 1 uur", time.Hour},
		{"in 5 minuten", 0, "in 5 minuten", time.Minute * 5},
		{"Binnen 5 minuten ga ik naar huis", 0, "Binnen 5 minuten", time.Minute * 5},
		{"we moeten binnen 10 dagen iets doen", 10, "binnen 10 dagen", 10 * 24 * time.Hour},
		{"we moeten binnen vijf dagen iets doen", 10, "binnen vijf dagen", 5 * 24 * time.Hour},
		{"we moeten over 5 dagen iets doen", 10, "over 5 dagen", 5 * 24 * time.Hour},
		{"In 5 seconde moet een auto verplaatsen", 0, "In 5 seconde", 5 * time.Second},
		{"binnen twee weken", 0, "binnen twee weken", 14 * 24 * time.Hour},
		{"binnen een maand", 0, "binnen een maand", 31 * 24 * time.Hour},
		{"na een maand", 0, "na een maand", 31 * 24 * time.Hour},
		{"binnen een paar maanden", 0, "binnen een paar maanden", 91 * 24 * time.Hour},
		{"binnen een jaar", 0, "binnen een jaar", 366 * 24 * time.Hour},
		{"in een week", 0, "in een week", 7 * 24 * time.Hour},
	}

	w := when.New(nil)
	w.Add(nl.Deadline(rules.Skip))

	ApplyFixtures(t, "nl.Deadline", w, fixt)
}
