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
		{"The Deadline is 10/10/2016", 16, "10/10/2016", (284 - OFFSET) * 24 * time.Hour},
		{"The Deadline is 1/2/2016", 16, "1/2/2016", (32 - OFFSET) * 24 * time.Hour},
		{"The Deadline is 29/2/2016", 16, "29/2/2016", (60 - OFFSET) * 24 * time.Hour},

		// next year
		{"The Deadline is 28/2", 16, "28/2", (59 + 366 - OFFSET) * 24 * time.Hour},
		{"The Deadline is 28/02/2017", 16, "28/02/2017", (59 + 366 - OFFSET) * 24 * time.Hour},

		// right after w/o a year
		{"The Deadline is 28/07", 16, "28/07", (210 - OFFSET) * 24 * time.Hour},

		// before w/o a year
		{"The Deadline is 30/06", 16, "30/06", (181 + 366 - OFFSET) * 24 * time.Hour},

		// prev day will be added to the future
		{"The Deadline is 14/07", 16, "14/07", (195 + 366 - OFFSET) * 24 * time.Hour},
	}

	w := when.New(nil)
	w.Add(common.SlashDMY(rules.Skip))

	ApplyFixtures(t, "common.SlashDMY", w, fixt)

	nilFixt := []Fixture{
		{"The Deadline is 1/20/2016", 16, "no match for mm/dd/yyyy", 0},
	}
	ApplyFixturesNil(t, "common.SlashDMY nil", w, nilFixt)
}
