package ru_test

import (
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/ru"
	"testing"
	"time"
)

func TestDotDateTime(t *testing.T) {
	w := when.New(nil)
	w.Add(ru.DotDateTime(rules.Override))

	fixt := []Fixture{
		// Basic date/time formats
		{"встреча 15.01.2024 09:30", 15, "15.01.2024 09:30", time.Date(2024, 1, 15, 9, 30, 0, 0, time.UTC).Sub(null)},
		{"05.03.2025 15:00 запланирована встреча", 0, "05.03.2025 15:00", time.Date(2025, 3, 5, 15, 0, 0, 0, time.UTC).Sub(null)},
		{"31.12.2023 23:59", 0, "31.12.2023 23:59", time.Date(2023, 12, 31, 23, 59, 0, 0, time.UTC).Sub(null)},
	}

	ApplyFixtures(t, "ru.DateTime", w, fixt)
}

func TestDotDateTimeNil(t *testing.T) {
	w := when.New(nil)
	w.Add(ru.DotDateTime(rules.Override))

	fixt := []Fixture{
		{"это текст без даты и времени", 0, "", 0},
		{"15.01", 0, "", 0},
		{"32.01.2024 15:00", 0, "", 0}, // некорректный день
		{"15.13.2024 15:00", 0, "", 0}, // некорректный месяц
	}

	ApplyFixturesNil(t, "ru.DateTime nil", w, fixt)
}
