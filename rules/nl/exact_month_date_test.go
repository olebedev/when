package nl_test

import (
	"github.com/olebedev/when/rules/nl"
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/stretchr/testify/require"
)

func TestExactMonthDate(t *testing.T) {
	w := when.New(nil)
	w.Add(nl.ExactMonthDate(rules.Override))

	fixtok := []Fixture{
		{"derde van maart", 0, "derde van maart", 1368 * time.Hour},
		{"3e van maart", 0, "3e van maart", 1368 * time.Hour},
		{"1 september", 0, "1 september", 5736 * time.Hour},
		{"1 sept", 0, "1 sept", 5736 * time.Hour},
		{"1 sept.", 0, "1 sept.", 5736 * time.Hour},
		{"1e van september", 0, "1e van september", 5736 * time.Hour},
		{"twintigste van december", 0, "twintigste van december", 8376 * time.Hour},
		{"februari", 0, "februari", 744 * time.Hour},
		{"oktober", 0, "oktober", 6576 * time.Hour},
		{"jul.", 0, "jul.", 4368 * time.Hour},
		{"juni", 0, "juni", 3648 * time.Hour},
	}

	ApplyFixtures(t, "nl.ExactMonthDate", w, fixtok)
}

func TestExactMonthDateInvalidDay(t *testing.T) {
	w := when.New(nil)
	w.Add(nl.ExactMonthDate(rules.Override))

	invalid := []string{
		"30 februari",
		"31e van februari",
		"31 april",
		"31 juni",
		"31 september",
		"31 november",
		"0 maart",
	}
	for _, year := range []int{2016, 2017} {
		ref := time.Date(year, time.January, 6, 0, 0, 0, 0, time.UTC)
		for i, text := range invalid {
			res, err := w.Parse(text, ref)
			require.Nil(t, err, "err #%d", i)
			require.Nil(t, res, "%q should not parse in %d", text, year)
		}
	}

	leapRef := time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC)
	res, err := w.Parse("29 februari", leapRef)
	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, time.Date(2016, time.February, 29, 0, 0, 0, 0, time.UTC), res.Time)

	nonLeapRef := time.Date(2017, time.January, 6, 0, 0, 0, 0, time.UTC)
	res, err = w.Parse("29 februari", nonLeapRef)
	require.Nil(t, err)
	require.Nil(t, res, "29 februari should not parse in a non-leap year")
}
