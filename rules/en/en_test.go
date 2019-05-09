package en_test

import (
	"testing"
	"time"

	"github.com/omniboost/when"
	"github.com/omniboost/when/rules"
	"github.com/omniboost/when/rules/en"
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
		require.Nil(t, err, "[%s] err #%d", name, i)
		require.NotNil(t, res, "[%s] res #%d", name, i)
		require.Equal(t, f.Index, res.Index, "[%s] index #%d", name, i)
		require.Equal(t, f.Phrase, res.Text, "[%s] text #%d", name, i)
		require.Equal(t, f.Diff, res.Time.Sub(null), "[%s] diff #%d", name, i)
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
	w.Add(en.All...)

	// complex cases
	fixt := []Fixture{
		{"tonight at 11:10 pm", 0, "tonight at 11:10 pm", (23 * time.Hour) + (10 * time.Minute)},
		{"at Friday afternoon", 3, "Friday afternoon", ((2 * 24) + 15) * time.Hour},
		{"in next tuesday at 14:00", 3, "next tuesday at 14:00", ((6 * 24) + 14) * time.Hour},
		{"in next tuesday at 2p", 3, "next tuesday at 2p", ((6 * 24) + 14) * time.Hour},
		{"in next wednesday at 2:25 p.m.", 3, "next wednesday at 2:25 p.m.", (((7 * 24) + 14) * time.Hour) + (25 * time.Minute)},
		{"at 11 am past tuesday", 3, "11 am past tuesday", -13 * time.Hour},
	}

	ApplyFixtures(t, "en.All...", w, fixt)
}

func TestAllMerge(t *testing.T) {
	w := when.New(nil)
	w.Add(
		en.Weekday(rules.Merge),
		en.CasualDate(rules.Merge),
		en.CasualTime(rules.Merge),
		en.Hour(rules.Override),
		en.HourMinute(rules.Override),
		en.HourMinuteSecond(rules.Override),
		en.Deadline(rules.Override),
		en.PastTime(rules.Merge),
		en.ExactMonthDate(rules.Override),
	)

	// current is Friday
	fixt := []Fixture{
		{"this monday one week ago", 0, "this monday one week ago", -9 * 24 * time.Hour},
		{"next monday one week ago", 0, "next monday one week ago", -2 * 24 * time.Hour},
	}

	ApplyFixtures(t, "en.All...", w, fixt)
}
