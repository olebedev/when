package common_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/common"
)

func TestSlashDMY(t *testing.T) {
	fixt := []Fixture{
		{"The Deadline is 10/10/2016", 16, "10/10/2016", time.Date(2016, 10, 10, 0, 0, 0, 0, time.UTC)},
		{"The Deadline is 1/2/2016", 16, "1/2/2016", time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC)},
		{"The Deadline is 29/2/2016", 16, "29/2/2016", time.Date(2016, 2, 29, 0, 0, 0, 0, time.UTC)},

		// next year
		{"The Deadline is 28/2", 16, "28/2", time.Date(2017, 2, 28, 0, 0, 0, 0, time.UTC)},
		{"The Deadline is 28/02/2017", 16, "28/02/2017", time.Date(2017, 2, 28, 0, 0, 0, 0, time.UTC)},

		// right after w/o a year
		{"The Deadline is 28/07", 16, "28/07", time.Date(2016, 7, 28, 0, 0, 0, 0, time.UTC)},

		// before w/o a year
		{"The Deadline is 30/06", 16, "30/06", time.Date(2017, 6, 30, 0, 0, 0, 0, time.UTC)},

		// prev day will be added to the future
		{"The Deadline is 14/07", 16, "14/07", time.Date(2017, 7, 14, 0, 0, 0, 0, time.UTC)},

		// Existing doesn't work for a month in the future
		{"The Deadline is 14/08", 16, "14/08", time.Date(2016, 8, 14, 0, 0, 0, 0, time.UTC)},
		{"The Deadline is 15/07", 16, "15/07", time.Date(2016, 7, 15, 0, 0, 0, 0, time.UTC)},
	}

	w := when.New(nil)
	w.Add(common.SlashDMY(rules.Skip))

	ApplyFixtures(t, "common.SlashDMY", w, fixt)

}

func TestSlashDMYPast(t *testing.T) {
	fixt := []Fixture{
		{"The Deadline is 10/10/2016", 16, "10/10/2016", time.Date(2016, 10, 10, 0, 0, 0, 0, time.UTC)},
		{"The Deadline is 1/2/2016", 16, "1/2/2016", time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC)},
		{"The Deadline is 29/2/2016", 16, "29/2/2016", time.Date(2016, 2, 29, 0, 0, 0, 0, time.UTC)},

		// before w/o a year says same year
		{"The Deadline is 30/06", 16, "30/06", time.Date(2016, 6, 30, 0, 0, 0, 0, time.UTC)},

		// prev day will still be this year
		{"The Deadline is 14/07", 16, "14/07", time.Date(2016, 7, 14, 0, 0, 0, 0, time.UTC)},

		// after w/o a year is prior year
		{"The Deadline is 28/07", 16, "28/07", time.Date(2015, 7, 28, 0, 0, 0, 0, time.UTC)},

		// Regression tests: current date and furture month
		{"The Deadline is 15/07", 16, "15/07", time.Date(2016, 7, 15, 0, 0, 0, 0, time.UTC)},
	}

	w := when.New(&rules.Options{WantPast: true})
	w.Add(common.SlashDMY(rules.Skip))

	ApplyFixtures(t, "common.SlashDMY", w, fixt)
}
