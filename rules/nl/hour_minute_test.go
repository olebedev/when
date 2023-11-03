package nl_test

import (
	"github.com/olebedev/when/rules/nl"
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
)

func TestHourMinute(t *testing.T) {
	w := when.New(nil)
	w.Add(nl.HourMinute(rules.Override))

	fixtok := []Fixture{
		{"17:30u", 0, "17:30u", (17 * time.Hour) + (30 * time.Minute)},
		{"om 17:30 uur", 3, "17:30 uur", (17 * time.Hour) + (30 * time.Minute)},
		{"om 5:59 pm", 3, "5:59 pm", (17 * time.Hour) + (59 * time.Minute)},
		{"om 5:59 am", 3, "5:59 am", (5 * time.Hour) + (59 * time.Minute)},
	}

	fixtnil := []Fixture{
		{"28:30pm", 0, "", 0},
		{"12:61u", 0, "", 0},
		{"24:10", 0, "", 0},
	}

	ApplyFixtures(t, "nl.HourMinute", w, fixtok)
	ApplyFixturesNil(t, "on.HourMinute nil", w, fixtnil)

	w.Add(nl.Hour(rules.Skip))
	ApplyFixtures(t, "nl.HourMinute|nl.Hour", w, fixtok)
	ApplyFixturesNil(t, "on.HourMinute|nl.Hour nil", w, fixtnil)

	w = when.New(nil)
	w.Add(
		nl.Hour(rules.Override),
		nl.HourMinute(rules.Override),
	)

	ApplyFixtures(t, "nl.Hour|nl.HourMinute", w, fixtok)
	ApplyFixturesNil(t, "on.Hour|nl.HourMinute nil", w, fixtnil)
}
