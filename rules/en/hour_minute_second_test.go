package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
)

func TestHourMinuteSecond(t *testing.T) {
	w := when.New(nil)
	w.Add(en.HourMinuteSecond(rules.Override))

	fixtok := []Fixture{
		{"5:30:10pm", 0, "5:30:10pm", (17 * time.Hour) + (30 * time.Minute) + (10 * time.Second)},
	}

	fixtnil := []Fixture{
		{"28:30:30pm", 0, "", 0},
		{"12:61:30pm", 0, "", 0},
		{"24:10:61", 0, "", 0},
		{"24:10:61", 0, "", 0},
	}

	ApplyFixtures(t, "en.HourMinuteSecond", w, fixtok)
	ApplyFixturesNil(t, "on.HourMinuteSecond nil", w, fixtnil)

	w.Add(en.Hour(rules.Skip))
	ApplyFixtures(t, "en.HourMinuteSecond|en.Hour", w, fixtok)
	ApplyFixturesNil(t, "on.HourMinuteSecond|en.Hour nil", w, fixtnil)

	w = when.New(nil)
	w.Add(
		en.HourMinuteSecond(rules.Override),
		en.Hour(rules.Override),
		en.HourMinute(rules.Override),
	)

	ApplyFixtures(t, "en.Hour|en.HourMinute|en.HourMinuteSecond", w, fixtok)
	ApplyFixturesNil(t, "on.Hour|en.HourMinute|en.HourMinuteSecond nil", w, fixtnil)
}
