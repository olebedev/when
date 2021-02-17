package ru_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/ru"
)

func Test_VoiceDeadline(t *testing.T) {
	w := when.New(nil)
	w.Add(ru.VoiceDeadline(rules.Override))

	fixt := []Fixture{
		{"через 7 часов", 0, "через 7 часов", 7 * time.Hour},
		{"через 15 минут", 0, "через 15 минут", 15 * time.Minute},
		{"через 7 часов 30 минут", 0, "через 7 часов 30 минут", 7*time.Hour + 30*time.Minute},
		{"через 15 часов", 0, "через 15 часов", 15 * time.Hour},
		{"включи через 7 часов 50 процентов яркости", 13, "через 7 часов", 7 * time.Hour},
		{"включи через 12 минут", 13, "через 12 минут", 12 * time.Minute},
		{"включи через 10 часов", 13, "через 10 часов", 10 * time.Hour},
		{"через 7 часов 30", 0, "через 7 часов", 7 * time.Hour},

		{"через 7 часов 30 минут десять секунд", 0,
			"через 7 часов 30 минут десять секунд", 7*time.Hour + 30*time.Minute + 10*time.Second},
		{"через 59 минут 59 секунд", 0, "через 59 минут 59 секунд", 59*time.Minute + 59*time.Second},
		{"через 118 секунд", 0, "через 118 секунд", 118 * time.Second},
	}

	ApplyFixtures(t, "ru.VoiceDeadline", w, fixt)

	fixtNone := []Fixture{
		{"через 7 30", 3, "7 30", 0},
		{"через 15 40", 3, "15 40", 0},

		{"В 6:", 0, "", 0},
		{"В 6:30", 0, "", 0},
		{"сделай яркость в 20 процентов", 0, "", 0},
		{"увеличь яркость в 2 раза", 0, "", 0},
		{"В 20 минут", 0, "", 0},
		{"через пень колоду", 0, "", 0},
		{"через полчаса", 0, "", 0},
	}

	ApplyFixturesNil(t, "ru.VoiceDeadline (none)", w, fixtNone)
}
