package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
)

func TestHourMinute(t *testing.T) {
	w := when.New(nil)
	w.Add(en.HourMinute(rules.OverWrite))

	fixtok := []Fixture{
		{"5:30pm", 0, "5:30pm", (17 * time.Hour) + (30 * time.Minute)},
		{"at 5:30 pm", 3, "5:30 pm", (17 * time.Hour) + (30 * time.Minute)},
		{"at 5:59 pm", 3, "5:59 pm", (17 * time.Hour) + (59 * time.Minute)},
		{"at 5-59 pm", 3, "5-59 pm", (17 * time.Hour) + (59 * time.Minute)},
		{"at 17-59 pam", 3, "17-59", (17 * time.Hour) + (59 * time.Minute)},
		{"up to 11:10 pm", 6, "11:10 pm", (23 * time.Hour) + (10 * time.Minute)},
	}

	fixtnil := []Fixture{
		{"28:30pm", 0, "", 0},
		{"12:61pm", 0, "", 0},
		{"24:10", 0, "", 0},
	}

	ApplyFixtures(t, "en.HourMinute", w, fixtok)
	ApplyFixturesNil(t, "on.HourMinute nil", w, fixtnil)

	w.Add(en.Hour(rules.Skip))
	ApplyFixtures(t, "en.HourMinute|en.Hour", w, fixtok)
	ApplyFixturesNil(t, "on.HourMinute|en.Hour nil", w, fixtnil)

	w = when.New(nil)
	w.Add(
		en.Hour(rules.OverWrite),
		en.HourMinute(rules.OverWrite),
	)

	ApplyFixtures(t, "en.Hour|en.HourMinute", w, fixtok)
	ApplyFixturesNil(t, "on.Hour|en.HourMinute nil", w, fixtnil)
}
