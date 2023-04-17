package common_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/common"
)

func TestSlashMDY(t *testing.T) {
	fixt := []Fixture{
		{"The Deadline is 10/10/2016", 16, "10/10/2016", time.Date(2016, 10, 10, 0, 0, 0, 0, time.UTC)},
		{"The Deadline is 2/1/2016", 16, "2/1/2016", time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC)},
		{"The Deadline is 2/29/2016", 16, "2/29/2016", time.Date(2016, 2, 29, 0, 0, 0, 0, time.UTC)},

		// next year
		{"The Deadline is 2/28", 16, "2/28", time.Date(2017, 2, 28, 0, 0, 0, 0, time.UTC)},
		{"The Deadline is 02/28/2017", 16, "02/28/2017", time.Date(2017, 2, 28, 0, 0, 0, 0, time.UTC)},

		// right after w/o a year
		{"The Deadline is 07/28", 16, "07/28", time.Date(2016, 7, 28, 0, 0, 0, 0, time.UTC)},

		// before w/o a year
		{"The Deadline is 06/30", 16, "06/30", time.Date(2017, 6, 30, 0, 0, 0, 0, time.UTC)},

		// prev day will be added to the future
		{"The Deadline is 07/14", 16, "07/14", time.Date(2017, time.July, 14, 0, 0, 0, 0, time.UTC)},

		// Current day or future months
		{"The Deadline is 8/14", 16, "8/14", time.Date(2016, 8, 14, 0, 0, 0, 0, time.UTC)},
		{"The Deadline is 7/15", 16, "7/15", time.Date(2016, 7, 15, 0, 0, 0, 0, time.UTC)},
	}

	w := when.New(nil)
	w.Add(common.SlashMDY(rules.Override))

	ApplyFixtures(t, "common.SlashMDY", w, fixt)

}

func TestSlashMDYPast(t *testing.T) {
	fixt := []Fixture{
		{"The Deadline is 10/10/2016", 16, "10/10/2016", time.Date(2016, 10, 10, 0, 0, 0, 0, time.UTC)},
		{"The Deadline is 2/1/2016", 16, "2/1/2016", time.Date(2016, 2, 1, 0, 0, 0, 0, time.UTC)},
		{"The Deadline is 2/29/2016", 16, "2/29/2016", time.Date(2016, 2, 29, 0, 0, 0, 0, time.UTC)},

		// before w/o a year says same year
		{"The Deadline is 06/30", 16, "06/30", time.Date(2016, 6, 30, 0, 0, 0, 0, time.UTC)},

		// prev day will still be this year
		{"The Deadline is 07/14", 16, "07/14", time.Date(2016, 7, 14, 0, 0, 0, 0, time.UTC)},

		// after w/o a year is prior year
		{"The Deadline is 07/28", 16, "07/28", time.Date(2015, 7, 28, 0, 0, 0, 0, time.UTC)},

		// Regression tests: current date and furture month
		{"The Deadline is 07/15", 16, "07/15", time.Date(2016, 7, 15, 0, 0, 0, 0, time.UTC)},
	}

	w := when.New(&rules.Options{WantPast: true})
	w.Add(common.SlashMDY(rules.Skip))

	ApplyFixtures(t, "common.SlashMDY", w, fixt)
}
