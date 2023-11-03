package nl_test

import (
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/nl"
	"testing"
	"time"
)

func TestHour(t *testing.T) {
	fixt := []Fixture{
		{"5pm", 0, "5pm", 17 * time.Hour},
		{"5 uur in de avond", 0, "5 uur in de avond", 17 * time.Hour},
		{"5 uur 's avonds", 0, "5 uur 's avonds", 17 * time.Hour},
		{"om 17 uur", 3, "17 uur", 17 * time.Hour},
		{"om 5 P.", 3, "5 P.", 17 * time.Hour},
		{"om 12 P.", 3, "12 P.", 12 * time.Hour},
		{"om 1 P.", 3, "1 P.", 13 * time.Hour},
		{"om 5 am", 3, "5 am", 5 * time.Hour},
		{"om 5A", 3, "5A", 5 * time.Hour},
		{"om 5A.", 3, "5A.", 5 * time.Hour},
		{"5A.", 0, "5A.", 5 * time.Hour},
		{"11 P.M.", 0, "11 P.M.", 23 * time.Hour},
	}

	w := when.New(nil)
	w.Add(nl.Hour(rules.Override))

	ApplyFixtures(t, "nl.Hour", w, fixt)
}
