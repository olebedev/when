package ru_test

import (
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/ru"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	w := when.New(nil)
	w.Add(ru.Date(rules.Override))

	fixt := []Fixture{
		// Simple dates
		{"встреча 15 января 2024", 15, "15 января 2024", time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC).Sub(null)},
		{"5 марта 2025 запланирована встреча", 0, "5 марта 2025", time.Date(2025, 3, 5, 0, 0, 0, 0, time.UTC).Sub(null)},
		{"31 декабря 2023", 0, "31 декабря 2023", time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC).Sub(null)},

		// Dates with time
		{"15 января 2024 в 9:30", 0, "15 января 2024 в 9:30", time.Date(2024, 1, 15, 9, 30, 0, 0, time.UTC).Sub(null)},
		{"5 марта 2025 в 15:00 запланирована встреча", 0, "5 марта 2025 в 15:00", time.Date(2025, 3, 5, 15, 0, 0, 0, time.UTC).Sub(null)},
		{"31 декабря 2023 в 23:59", 0, "31 декабря 2023 в 23:59", time.Date(2023, 12, 31, 23, 59, 0, 0, time.UTC).Sub(null)},
	}

	ApplyFixtures(t, "ru.Date", w, fixt)
}

func TestDateNil(t *testing.T) {
	w := when.New(nil)
	w.Add(ru.Date(rules.Override))

	fixt := []Fixture{
		{"это текст без даты", 0, "", 0},
		{"15", 0, "", 0},
		{"15 чего-то", 0, "", 0},
	}

	ApplyFixturesNil(t, "ru.Date nil", w, fixt)
}
