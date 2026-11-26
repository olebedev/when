package br_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/br"
	"github.com/stretchr/testify/require"
)

func TestExactMonthDate(t *testing.T) {
	w := when.New(nil)
	w.Add(br.ExactMonthDate(rules.Override))

	fixtok := []Fixture{
		{"3 de março", 0, "3 de março", 1368 * time.Hour},
		{"1 de setembro", 0, "1 de setembro", 5736 * time.Hour},
		{"1 set", 0, "1 set", 5736 * time.Hour},
		{"1 set.", 0, "1 set.", 5736 * time.Hour},
		{"1º de setembro", 0, "1º de setembro", 5736 * time.Hour},
		{"1º set.", 0, "1º set.", 5736 * time.Hour},
		{"7 de março", 0, "7 de março", 1464 * time.Hour},
		{"21 de outubro", 0, "21 de outubro", 6936 * time.Hour},
		{"vigésimo dia de dezembro", 0, "vigésimo dia de dezembro", 8376 * time.Hour},
		{"10º dia de março", 0, "10º dia de março", 1536 * time.Hour},
		{"4 jan.", 0, "4 jan.", -48 * time.Hour},
		{"fevereiro", 0, "fevereiro", 744 * time.Hour},
		{"outubro", 0, "outubro", 6576 * time.Hour},
		{"jul.", 0, "jul.", 4368 * time.Hour},
		{"junho", 0, "junho", 3648 * time.Hour},
	}

	ApplyFixtures(t, "br.ExactMonthDate", w, fixtok)
}

func TestExactMonthDateInvalidDay(t *testing.T) {
	w := when.New(nil)
	w.Add(br.ExactMonthDate(rules.Override))

	invalid := []string{
		"30 de fevereiro",
		"31 de fevereiro",
		"31 de abril",
		"31 de junho",
		"31 de setembro",
		"31 de novembro",
		"0 de março",
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
	res, err := w.Parse("29 de fevereiro", leapRef)
	require.Nil(t, err)
	require.NotNil(t, res)
	require.Equal(t, time.Date(2016, time.February, 29, 0, 0, 0, 0, time.UTC), res.Time)

	nonLeapRef := time.Date(2017, time.January, 6, 0, 0, 0, 0, time.UTC)
	res, err = w.Parse("29 de fevereiro", nonLeapRef)
	require.Nil(t, err)
	require.Nil(t, res, "29 de fevereiro should not parse in a non-leap year")
}
