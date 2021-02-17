package ru_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/ru"
)

func Test_VoiceHourMinute(t *testing.T) {
	w := when.New(nil)
	w.Add(ru.VoiceHourMinute(rules.Override))

	fixt := []Fixture{
		{"в 7 часов", 3, "7", 7 * time.Hour},
		{"в 7 30", 3, "7 30", 7*time.Hour + 30*time.Minute},
		{"в 15 40", 3, "15 40", 15*time.Hour + 40*time.Minute},
		{"в 7 30 вечера", 3, "7 30 вечера", 19*time.Hour + 30*time.Minute},
		{"в 7 часов вечера 30 минут", 3, "7 часов вечера 30 минут", 19*time.Hour + 30*time.Minute},
		{"в 7 часов 30", 3, "7 часов 30", 7*time.Hour + 30*time.Minute},
		{"в 7 часов 30 минут", 3, "7 часов 30 минут", 7*time.Hour + 30*time.Minute},
		{"в 7 часов 30 минут вечера", 3, "7 часов 30 минут вечера", 19*time.Hour + 30*time.Minute},
		// {"вечером в 7 30", 0, "вечером в 7 30", 19*time.Hour + 30*time.Minute},
		{"в 15 часов", 3, "15", 15 * time.Hour},
		{"включи в 7 50 процентов яркости", 16, "7 50 процент", 7 * time.Hour},
		{"включи в 12 ночи", 16, "12 ночи", 0 * time.Hour},
	}

	ApplyFixtures(t, "ru.All...", w, fixt)

	fixtNone := []Fixture{
		{"В 6:", 0, "", 0},
		{"В 6:30", 0, "", 0},
		{"через 15 минут", 11, "", 0},
		{"сделай яркость в 20 процентов", 0, "", 0},
		{"увеличь яркость в 2 раза", 0, "", 0},
		{"В 20 минут", 0, "", 0},
	}

	ApplyFixturesNil(t, "ru.None", w, fixtNone)
}
