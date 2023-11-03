package nl_test

import (
	"github.com/olebedev/when/rules/nl"
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
)

func TestPastTime(t *testing.T) {
	fixt := []Fixture{
		{"een half uur geleden", 0, "een half uur geleden", -(time.Hour / 2)},
		{"1 uur geleden", 0, "1 uur geleden", -(time.Hour)},
		{"5 minuten geleden", 0, "5 minuten geleden", -(time.Minute * 5)},
		{"5 minuten geleden ging ik naar de dierentuin", 0, "5 minuten geleden", -(time.Minute * 5)},
		{"we deden iets 10 dagen geleden", 14, "10 dagen geleden", -(10 * 24 * time.Hour)},
		{"we deden iets vijf dagen geleden", 14, "vijf dagen geleden", -(5 * 24 * time.Hour)},
		{"we deden iets 5 dagen geleden", 14, "5 dagen geleden", -(5 * 24 * time.Hour)},
		{"5 seconden geleden werd een auto weggesleept", 0, "5 seconden geleden", -(5 * time.Second)},
		{"twee weken geleden", 0, "twee weken geleden", -(14 * 24 * time.Hour)},
		{"een maand geleden", 0, "een maand geleden", -(31 * 24 * time.Hour)},
		{"een paar maanden geleden", 0, "een paar maanden geleden", -(92 * 24 * time.Hour)},
		{"een jaar geleden", 0, "een jaar geleden", -(365 * 24 * time.Hour)},
		{"een week geleden", 0, "een week geleden", -(7 * 24 * time.Hour)},
	}

	w := when.New(nil)
	w.Add(nl.PastTime(rules.Skip))

	ApplyFixtures(t, "nl.PastTime", w, fixt)
}
