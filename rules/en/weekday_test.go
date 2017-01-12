package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
)

func TestWeekday(t *testing.T) {
	// current is Friday
	fixt := []Fixture{
		// past/last
		{"do it for the past Monday", 14, "past Monday", -(2 * 24 * time.Hour)},
		{"past saturday", 0, "past saturday", -(4 * 24 * time.Hour)},
		{"past friday", 0, "past friday", -(5 * 24 * time.Hour)},
		{"past wednesday", 0, "past wednesday", -(7 * 24 * time.Hour)},
		{"past tuesday", 0, "past tuesday", -(24 * time.Hour)},
		// next
		{"next tuesday", 0, "next tuesday", 6 * 24 * time.Hour},
		{"drop me a line at next wednesday", 18, "next wednesday", 7 * 24 * time.Hour},
		{"next saturday", 0, "next saturday", 3 * 24 * time.Hour},
		// this
		{"this tuesday", 0, "this tuesday", -(24 * time.Hour)},
		{"drop me a line at this wednesday", 18, "this wednesday", 0},
		{"this saturday", 0, "this saturday", 3 * 24 * time.Hour},
	}

	w := when.New(nil)

	w.Add(en.Weekday(rules.Override))

	ApplyFixtures(t, "en.Weekday", w, fixt)
}
