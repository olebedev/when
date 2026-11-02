package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
	"github.com/stretchr/testify/require"
)

func TestPastTime(t *testing.T) {
	fixt := []Fixture{
		{"half an hour ago", 0, "half an hour ago", -(time.Hour / 2)},
		{"1 hour ago", 0, "1 hour ago", -(time.Hour)},
		{"5 minutes ago", 0, "5 minutes ago", -(time.Minute * 5)},
		{"5 minutes ago I went to the zoo", 0, "5 minutes ago", -(time.Minute * 5)},
		{"we did something 10 days ago.", 17, "10 days ago", -(10 * 24 * time.Hour)},
		{"we did something five days ago.", 17, "five days ago", -(5 * 24 * time.Hour)},
		{"we did something 5 days ago.", 17, "5 days ago", -(5 * 24 * time.Hour)},
		{"5 seconds ago a car was moved", 0, "5 seconds ago", -(5 * time.Second)},
		{"two weeks ago", 0, "two weeks ago", -(14 * 24 * time.Hour)},
		{"a month ago", 0, "a month ago", -(31 * 24 * time.Hour)},
		{"a few months ago", 0, "a few months ago", -(92 * 24 * time.Hour)},
		{"one year ago", 0, "one year ago", -(365 * 24 * time.Hour)},
		{"a week ago", 0, "a week ago", -(7 * 24 * time.Hour)},
	}

	w := when.New(nil)
	w.Add(en.PastTime(rules.Skip))

	ApplyFixtures(t, "en.PastTime", w, fixt)
}

func TestPastTimeMonthsCrossYear(t *testing.T) {
	ref := time.Date(2016, time.February, 6, 0, 0, 0, 0, time.UTC)
	fixt := []struct {
		text string
		want time.Time
	}{
		{"2 months ago", time.Date(2015, time.December, 6, 0, 0, 0, 0, time.UTC)},
		{"14 months ago", time.Date(2014, time.December, 6, 0, 0, 0, 0, time.UTC)},
		{"half a year ago", time.Date(2015, time.August, 6, 0, 0, 0, 0, time.UTC)},
	}

	w := when.New(nil)
	w.Add(en.PastTime(rules.Skip))

	for _, f := range fixt {
		res, err := w.Parse(f.text, ref)
		require.Nil(t, err, f.text)
		require.NotNil(t, res, f.text)
		require.True(t, f.want.Equal(res.Time), "%s: got %s", f.text, res.Time)
	}
}

// Parsed against the 29th of February, the target year has no such day and
// used to overflow into March.
func TestPastTimeYearFromLeapDay(t *testing.T) {
	ref := time.Date(2024, time.February, 29, 0, 0, 0, 0, time.UTC)
	fixt := []struct {
		text string
		want time.Time
	}{
		{"a year ago", time.Date(2023, time.February, 28, 0, 0, 0, 0, time.UTC)},
		{"4 years ago", time.Date(2020, time.February, 29, 0, 0, 0, 0, time.UTC)},
	}

	w := when.New(nil)
	w.Add(en.PastTime(rules.Skip))

	for _, f := range fixt {
		res, err := w.Parse(f.text, ref)
		require.Nil(t, err, f.text)
		require.NotNil(t, res, f.text)
		require.True(t, f.want.Equal(res.Time), "%s: got %s", f.text, res.Time)
	}
}
