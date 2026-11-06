package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
)

// July 15 days offset from the begining of the year
const OFFSET = 197

func TestSlashMDY(t *testing.T) {
	fixt := []Fixture{
		{"The Deadline is 10/10/2016", 16, "10/10/2016", (284 - OFFSET) * 24 * time.Hour},
		{"The Deadline is 2/1/2016", 16, "2/1/2016", (32 - OFFSET) * 24 * time.Hour},
		{"The Deadline is 2/29/2016", 16, "2/29/2016", (60 - OFFSET) * 24 * time.Hour},

		// next year
		{"The Deadline is 2/28", 16, "2/28", (59 + 366 - OFFSET) * 24 * time.Hour},
		{"The Deadline is 02/28/2017", 16, "02/28/2017", (59 + 366 - OFFSET) * 24 * time.Hour},

		// right after w/o a year
		{"The Deadline is 07/28", 16, "07/28", (210 - OFFSET) * 24 * time.Hour},

		// before w/o a year
		{"The Deadline is 06/30", 16, "06/30", (181 + 366 - OFFSET) * 24 * time.Hour},

		// prev day will be added to the future
		{"The Deadline is 07/14", 16, "07/14", (195 + 366 - OFFSET) * 24 * time.Hour},
	}

	w := when.New(nil)
	w.Add(en.SlashMDY(rules.Override))

	null = time.Date(2016, time.July, 15, 0, 0, 0, 0, time.UTC)

	ApplyFixtures(t, "en.SlashMDY", w, fixt)

	null = time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC)

}
