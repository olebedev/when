package ru_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
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
	}

	ApplyFixtures(t, "ru.All...", w, fixt)
}

func TestVoice(t *testing.T) {
	w := when.New(nil)
	w.Add(
		ru.Hour(rules.Override),
		ru.HourMinute(rules.Override),
		ru.Weekday(rules.Override),
		ru.CasualDate(rules.Override),
		ru.CasualTime(rules.Override),
		ru.Deadline(rules.Override),
		ru.VoiceDeadline(rules.Override),
		ru.VoiceHourNaked(rules.Override),
		ru.VoiceHourMinute(rules.Override),
	)

	// complex cases
	fixt := []Fixture{
		{"В 6:30", 3, "6:30", 6*time.Hour + 30*time.Minute},
		{"через полчаса", 0, "через полчаса", 30 * time.Minute},
		// {"через полтора часа", 0, "через полтора часа", 90 * time.Minute},
		// {"в полпервого", 3, "полпервого", 30 * time.Minute},
		{"через 30 минут", 0, "через 30 минут", 30 * time.Minute},
		// {"в 30 минут", 3, "30 минут", 30 * time.Minute},
		{"в 15 часов 10 минут", 3, "15 часов 10 минут", 15*time.Hour + 10*time.Minute},
		{"в 15 часов", 3, "15", 15 * time.Hour},
		{"в 3 часа", 3, "3", 3 * time.Hour},
		{"в 3 часа 30", 3, "3 часа 30", 3*time.Hour + 30*time.Minute},
		{"В 6 30", 3, "6 30", 6*time.Hour + 30*time.Minute},
		{"включи свет в 7", 25, "7", 7 * time.Hour},

		{"в 7 включи свет", 3, "7", 7 * time.Hour},
		{"в 15 часов 10 минут включи свет", 3, "15 часов 10 минут", 15*time.Hour + 10*time.Minute},
		{"в 15 часов включи свет", 3, "15", 15 * time.Hour},
		{"в 3 часа включи свет", 3, "3", 3 * time.Hour},
		{"в 3 часа 30 включи свет", 3, "3 часа 30", 3*time.Hour + 30*time.Minute},
		{"В 6 30 включи свет", 3, "6 30", 6*time.Hour + 30*time.Minute},
		{"через два часа десять минут", 0, "через два часа десять минут", 2*time.Hour + 10*time.Minute},
	}

	ApplyFixtures(t, "Voice", w, fixt)

	// complex cases
	fixtNone := []Fixture{
		{"яркость на 50", 0, "", 0},
		{"на лампе 1 увеличь яркость на 20", 0, "", 0},
	}

	ApplyFixturesNil(t, "Voice", w, fixtNone)
}
