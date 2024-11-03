package ru_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/ru"
	"github.com/stretchr/testify/require"
)

var null = time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC)

type Fixture struct {
	Text   string
	Index  int
	Phrase string
	Diff   time.Duration
}

func ApplyFixtures(t *testing.T, name string, w *when.Parser, fixt []Fixture) {
	for i, f := range fixt {
		res, err := w.Parse(f.Text, null)
		require.Nil(t, err, "[%s] err #%d - %s", name, i, f.Text)
		require.NotNil(t, res, "[%s] res #%d - %s", name, i, f.Text)
		require.Equal(t, f.Index, res.Index, "[%s] index #%d - %s", name, i, f.Text)
		require.Equal(t, f.Phrase, res.Text, "[%s] text #%d - %s", name, i, f.Text)
		require.Equal(t, f.Diff, res.Time.Sub(null), "[%s] diff #%d - %s", name, i, f.Text)
	}
}

func ApplyFixturesNil(t *testing.T, name string, w *when.Parser, fixt []Fixture) {
	for i, f := range fixt {
		res, err := w.Parse(f.Text, null)
		require.Nil(t, err, "[%s] err #%d", name, i)
		require.Nil(t, res, "[%s] res #%d", name, i)
	}
}

func ApplyFixturesErr(t *testing.T, name string, w *when.Parser, fixt []Fixture) {
	for i, f := range fixt {
		_, err := w.Parse(f.Text, null)
		require.NotNil(t, err, "[%s] err #%d", name, i)
		require.Equal(t, f.Phrase, err.Error(), "[%s] err text #%d", name, i)
	}
}

func TestAll(t *testing.T) {
	w := when.New(nil)
	w.Add(ru.All...)

	// complex cases
	fixt := []Fixture{
		{"завтра в 11:10 вечера", 0, "завтра в 11:10 вечера", (47 * time.Hour) + (10 * time.Minute)},
		{"вечером в следующий понедельник", 0, "вечером в следующий понедельник", ((5 * 24) + 18) * time.Hour},
		{"вечером в прошлый понедельник", 0, "вечером в прошлый понедельник", ((-2 * 24) + 18) * time.Hour},
		{"в следующий понедельник вечером", 3, "следующий понедельник вечером", ((5 * 24) + 18) * time.Hour},
		{"в Пятницу после обеда", 0, "в Пятницу после обеда", ((2 * 24) + 15) * time.Hour},
		{"в следующий вторник в 14:00", 3, "следующий вторник в 14:00", ((6 * 24) + 14) * time.Hour},
		{"в следующий вторник в четыре вечера", 3, "следующий вторник в четыре вечера", ((6 * 24) + 16) * time.Hour},
		{"в следующую среду в 2:25 вечера", 3, "следующую среду в 2:25 вечера", (((7 * 24) + 14) * time.Hour) + (25 * time.Minute)},
		{"в 11 утра в прошлый вторник", 3, "11 утра в прошлый вторник", -13 * time.Hour},

		{"написать письмо во вторник после обеда", 30, "во вторник после обеда", ((6 * 24) + 15) * time.Hour},
		{"написать письмо ко вторнику", 30, "ко вторнику", 6 * 24 * time.Hour},
		{"написать письмо до утра субботы ", 30, "до утра субботы", ((3 * 24) + 8) * time.Hour},
		{"написать письмо к субботе после обеда ", 30, "к субботе после обеда", ((3 * 24) + 15) * time.Hour},
		{"В субботу вечером", 0, "В субботу вечером", ((3 * 24) + 18) * time.Hour},

		{"встреча 15 января 2024", 15, "15 января 2024", time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC).Sub(null)},
		{"5 марта 2025 запланирована встреча", 0, "5 марта 2025", time.Date(2025, 3, 5, 0, 0, 0, 0, time.UTC).Sub(null)},
		{"31 декабря 2023", 0, "31 декабря 2023", time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC).Sub(null)},
		{"15 января 2024 в 9:30", 0, "15 января 2024 в 9:30", time.Date(2024, 1, 15, 9, 30, 0, 0, time.UTC).Sub(null)},
		{"5 марта 2025 в 15:00 запланирована встреча", 0, "5 марта 2025 в 15:00", time.Date(2025, 3, 5, 15, 0, 0, 0, time.UTC).Sub(null)},
		{"31 декабря 2023 в 23:59", 0, "31 декабря 2023 в 23:59", time.Date(2023, 12, 31, 23, 59, 0, 0, time.UTC).Sub(null)},
		{"31 декабря", 0, "31 декабря", time.Date(time.Now().Year(), 12, 31, 0, 0, 0, 0, time.UTC).Sub(null)},
		{"встреча 15.01.2024 09:30", 15, "15.01.2024 09:30", time.Date(2024, 1, 15, 9, 30, 0, 0, time.UTC).Sub(null)},
		{"05.03.2025 15:00 запланирована встреча", 0, "05.03.2025 15:00", time.Date(2025, 3, 5, 15, 0, 0, 0, time.UTC).Sub(null)},
		{"31.12.2023 23:59", 0, "31.12.2023 23:59", time.Date(2023, 12, 31, 23, 59, 0, 0, time.UTC).Sub(null)},
		{"31.12.2023", 0, "31.12.2023", time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC).Sub(null)},
	}

	ApplyFixtures(t, "ru.All...", w, fixt)
}
