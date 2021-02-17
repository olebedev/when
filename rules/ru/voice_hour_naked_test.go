package ru_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/ru"
)

func Test_VoiceHourNaked(t *testing.T) {
	w := when.New(nil)
	w.Add(ru.VoiceHourNaked(rules.Override))

	fixt := []Fixture{
		{"включи свет в 7", 25, "7", 7 * time.Hour},
		{"в 7 включи свет", 3, "7", 7 * time.Hour},
		{"в комнате в 7 включи свет", 21, "7", 7 * time.Hour},
		{"увеличь яркость на 50 процентов в 2", 60, "2", 2 * time.Hour},
		{"увеличь яркость в 2 на 50%", 33, "2", 2 * time.Hour},
		{"в три", 3, "три", 3 * time.Hour},
	}

	ApplyFixtures(t, "ru.VoiceHourNaked...", w, fixt)

	fixtNone := []Fixture{
		{"В 6:30", 0, "", 0},
		{"в 15 часов", 3, "15", 15 * time.Hour},
		{"в 15 40", 0, "", 0},
		{"через 15 минут", 11, "", 0},
		{"сделай яркость в 20 процентов", 0, "", 0},
		{"увеличь яркость в 2 раза", 0, "", 0},
	}

	ApplyFixturesNil(t, "ru.VoiceHourNaked (none)", w, fixtNone)
}
