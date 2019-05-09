package br_test

import (
	"testing"
	"time"

	"github.com/omniboost/when"
	"github.com/omniboost/when/rules"
	"github.com/omniboost/when/rules/br"
)

func TestHourMinute(t *testing.T) {
	w := when.New(nil)
	w.Add(br.HourMinute(rules.Override))

	fixtok := []Fixture{
		{"5:30pm", 0, "5:30pm", (17 * time.Hour) + (30 * time.Minute)},
		{"at 5:30 pm", 3, "5:30 pm", (17 * time.Hour) + (30 * time.Minute)},
		{"at 5:59 pm", 3, "5:59 pm", (17 * time.Hour) + (59 * time.Minute)},
		{"at 5-59 pm", 3, "5-59 pm", (17 * time.Hour) + (59 * time.Minute)},
		{"at 17-59 pam", 3, "17-59", (17 * time.Hour) + (59 * time.Minute)},
		{"up to 11:10 pm", 6, "11:10 pm", (23 * time.Hour) + (10 * time.Minute)},
		{"19h35m", 0, "19h35", (19 * time.Hour) + (35 * time.Minute)},
	}

	fixtnil := []Fixture{
		{"28:30pm", 0, "", 0},
		{"12:61pm", 0, "", 0},
		{"24:10", 0, "", 0},
	}

	ApplyFixtures(t, "br.HourMinute", w, fixtok)
	ApplyFixturesNil(t, "on.HourMinute nil", w, fixtnil)

	w.Add(br.Hour(rules.Skip))
	ApplyFixtures(t, "br.HourMinute|br.Hour", w, fixtok)
	ApplyFixturesNil(t, "on.HourMinute|br.Hour nil", w, fixtnil)

	w = when.New(nil)
	w.Add(
		br.Hour(rules.Override),
		br.HourMinute(rules.Override),
	)

	ApplyFixtures(t, "br.Hour|br.HourMinute", w, fixtok)
	ApplyFixturesNil(t, "on.Hour|br.HourMinute nil", w, fixtnil)
}
