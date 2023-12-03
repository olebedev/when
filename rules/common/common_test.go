package common_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/stretchr/testify/require"
)

var null = time.Date(2016, time.July, 15, 0, 0, 0, 0, time.UTC)

// July 15 days offset from the begining of the year
const OFFSET = 197

type Fixture struct {
	Text   string
	Index  int
	Phrase string
	Want   time.Time
}

func ApplyFixtures(t *testing.T, name string, w *when.Parser, fixt []Fixture) {
	for i, f := range fixt {
		res, err := w.Parse(f.Text, null)
		require.Nil(t, err, "[%s] err #%d", name, i)
		require.NotNil(t, res, "[%s] res #%d", name, i)
		require.Equal(t, f.Index, res.Index, "[%s] index #%d", name, i)
		require.Equal(t, f.Phrase, res.Text, "[%s] text #%d", name, i)
		require.Equal(t, f.Want, res.Time, "[%s] %s diff #%d", name, f.Phrase, i)
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
	w.Add(common.All...)

	// complex cases
	fixt := []Fixture{}
	ApplyFixtures(t, "common.All...", w, fixt)
}

func TestLeapYear(t *testing.T) {
	require.Equal(t, common.GetDays(1999, 2), 28, "Normal year")
	require.Equal(t, common.GetDays(2004, 2), 29, "Leap year")
	require.Equal(t, common.GetDays(3000, 2), 28, "Century")
	require.Equal(t, common.GetDays(2000, 2), 29, "Century divisible by 400")
}
