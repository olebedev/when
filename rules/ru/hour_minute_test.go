package ru_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/ru"
)

func TestHourMinute(t *testing.T) {
	w := when.New(nil)
	w.Add(ru.HourMinute(rules.OverWrite))

	fixtok := []Fixture{
		{"5:30вечера", 0, "5:30вечера", (17 * time.Hour) + (30 * time.Minute)},
		{"в 5:30 вечера", 3, "5:30 вечера", (17 * time.Hour) + (30 * time.Minute)},
		{"в 5:59 вечера", 3, "5:59 вечера", (17 * time.Hour) + (59 * time.Minute)},
		{"в 5-59 вечера", 3, "5-59 вечера", (17 * time.Hour) + (59 * time.Minute)},
		{"в 17-59 вечерело", 3, "17-59", (17 * time.Hour) + (59 * time.Minute)},
		{"до 11.10 вечера", 5, "11.10 вечера", (23 * time.Hour) + (10 * time.Minute)},
	}

	fixtnil := []Fixture{
		{"28:30вечера", 0, "", 0},
		{"12:61вечера", 0, "", 0},
		{"24:10", 0, "", 0},
	}

	ApplyFixtures(t, "ru.HourMinute", w, fixtok)
	ApplyFixturesNil(t, "on.HourMinute nil", w, fixtnil)

	w.Add(ru.Hour(rules.Skip))
	ApplyFixtures(t, "ru.HourMinute|ru.Hour", w, fixtok)
	ApplyFixturesNil(t, "on.HourMinute|ru.Hour nil", w, fixtnil)

	w = when.New(nil)
	w.Add(
		ru.Hour(rules.OverWrite),
		ru.HourMinute(rules.OverWrite),
	)

	ApplyFixtures(t, "ru.Hour|ru.HourMinute", w, fixtok)
	ApplyFixturesNil(t, "ru.Hour|ru.HourMinute nil", w, fixtnil)
}
