package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
	"github.com/stretchr/testify/require"
)

func TestExactMonthDate(t *testing.T) {
	w := when.New(nil)
	w.Add(en.ExactMonthDate(rules.Override))

	fixtok := []Fixture{
		{"third of march", 0, "third of march", 1368 * time.Hour},
		{"march third", 0, "march third", 1368 * time.Hour},
		{"march 3rd", 0, "march 3rd", 1368 * time.Hour},
		{"3rd march", 0, "3rd march", 1368 * time.Hour},
		{"march 3", 0, "march 3", 1368 * time.Hour},
		{"1 september", 0, "1 september", 5736 * time.Hour},
		{"1 sept", 0, "1 sept", 5736 * time.Hour},
		{"1 sept.", 0, "1 sept.", 5736 * time.Hour},
		{"1st of september", 0, "1st of september", 5736 * time.Hour},
		{"sept. 1st", 0, "sept. 1st", 5736 * time.Hour},
		{"march 7th", 0, "march 7th", 1464 * time.Hour},
		{"october 21st", 0, "october 21st", 6936 * time.Hour},
		{"twentieth of december", 0, "twentieth of december", 8376 * time.Hour},
		{"march 10th", 0, "march 10th", 1536 * time.Hour},
		{"jan. 4", 0, "jan. 4", -48 * time.Hour},
		{"february", 0, "february", 744 * time.Hour},
		{"october", 0, "october", 6576 * time.Hour},
		{"jul.", 0, "jul.", 4368 * time.Hour},
		{"june", 0, "june", 3648 * time.Hour},
	}

	ApplyFixtures(t, "en.ExactMonthDate", w, fixtok)
}

// Parsed against the 31st, a shorter target month used to overflow into the
// next one: "september 5th" resolved to the 5th of October.
func TestExactMonthDateFromLongMonth(t *testing.T) {
	w := when.New(nil)
	w.Add(en.ExactMonthDate(rules.Override))

	ref := time.Date(2016, time.October, 31, 0, 0, 0, 0, time.UTC)

	fixt := []struct {
		Text     string
		Expected time.Time
	}{
		{"september 5th", time.Date(2016, time.September, 5, 0, 0, 0, 0, time.UTC)},
		{"february 11", time.Date(2016, time.February, 11, 0, 0, 0, 0, time.UTC)},
		{"1st of june", time.Date(2016, time.June, 1, 0, 0, 0, 0, time.UTC)},
		// No day in the text: the reference day clamps to the month's last.
		{"september", time.Date(2016, time.September, 30, 0, 0, 0, 0, time.UTC)},
		// Months that do have a 31st keep the reference day untouched.
		{"december", time.Date(2016, time.December, 31, 0, 0, 0, 0, time.UTC)},
	}

	for i, f := range fixt {
		res, err := w.Parse(f.Text, ref)
		require.Nil(t, err, "[en.ExactMonthDate] err #%d", i)
		require.NotNil(t, res, "[en.ExactMonthDate] res #%d", i)
		require.Equal(t, f.Expected, res.Time, "[en.ExactMonthDate] time #%d", i)
	}
}
