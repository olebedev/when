package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
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

		// TODO: allow specifying the xth of the month
		// {"the 1st", 4, "1st", date(2016, 1, 1)},
		// {"the 10th", 4, "10th", date(2015, 1, 10)},
	}

	ApplyFixtures(t, "en.ExactMonthDate", w, fixtok)
}

func date(year int, month time.Month, day int) time.Duration {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Sub(null)
}

func TestExactMonthDatePast(t *testing.T) {
	w := when.New(&rules.Options{Distance: 5, MatchByOrder: true, WantPast: true})
	w.Add(en.ExactMonthDate(rules.Override))

	fixtok := []Fixture{
		{"third of march", 0, "third of march", date(2015, 3, 3)},
		{"march third", 0, "march third", date(2015, 3, 3)},
		{"march 3rd", 0, "march 3rd", date(2015, 3, 3)},
		{"3rd march", 0, "3rd march", date(2015, 3, 3)},
		{"march 3", 0, "march 3", date(2015, 3, 3)},
		{"1 september", 0, "1 september", date(2015, 9, 1)},
		{"1 sept", 0, "1 sept", date(2015, 9, 1)},
		{"1 sept.", 0, "1 sept.", date(2015, 9, 1)},
		{"1st of september", 0, "1st of september", date(2015, 9, 1)},
		{"sept. 1st", 0, "sept. 1st", date(2015, 9, 1)},
		{"march 7th", 0, "march 7th", date(2015, 3, 7)},
		{"october 21st", 0, "october 21st", date(2015, 10, 21)},
		{"twentieth of december", 0, "twentieth of december", date(2015, 12, 20)},
		{"march 10th", 0, "march 10th", date(2015, 3, 10)},
		{"jan 1st", 0, "jan 1st", date(2016, 1, 1)},
		{"jan. 4", 0, "jan. 4", date(2015, 1, 4)},
		{"fourth of jan", 0, "fourth of jan", date(2015, 1, 4)},
		{"january", 0, "january", date(2016, 1, 6)},
		{"october", 0, "october", date(2015, 10, 6)},
		{"jul.", 0, "jul.", date(2015, 7, 6)},
		{"june", 0, "june", date(2015, 6, 6)},

		// TODO: allow specifying the xth of the month
		// {"1st", 0, "1st", date(2016, 1, 1)},
		// {"10th", 0, "10th", date(2015, 1, 10)},
	}

	ApplyFixtures(t, "en.ExactMonthDate WantPast", w, fixtok)
}
