package en_test

import (
	"testing"
	"time"

	"github.com/omniboost/when"
	"github.com/omniboost/when/rules"
	"github.com/omniboost/when/rules/en"
)

func TestPastTime(t *testing.T) {
	fixt := []Fixture{
		{"half an hour ago", 0, "half an hour ago", -(time.Hour / 2)},
		{"1 hour ago", 0, "1 hour ago", -(time.Hour)},
		{"5 minutes ago", 0, "5 minutes ago", -(time.Minute * 5)},
		{"5 minutes ago I went to the zoo", 0, "5 minutes ago", -(time.Minute * 5)},
		{"we did something 10 days ago.", 17, "10 days ago", -(10 * 24 * time.Hour)},
		{"we did something five days ago.", 17, "five days ago", -(5 * 24 * time.Hour)},
		{"we did something 5 days ago.", 17, "5 days ago", -(5 * 24 * time.Hour)},
		{"5 seconds ago a car was moved", 0, "5 seconds ago", -(5 * time.Second)},
		{"two weeks ago", 0, "two weeks ago", -(14 * 24 * time.Hour)},
		{"a month ago", 0, "a month ago", -(31 * 24 * time.Hour)},
		{"a few months ago", 0, "a few months ago", -(92 * 24 * time.Hour)},
		{"one year ago", 0, "one year ago", -(365 * 24 * time.Hour)},
		{"a week ago", 0, "a week ago", -(7 * 24 * time.Hour)},
	}

	w := when.New(nil)
	w.Add(en.PastTime(rules.Override))

	ApplyFixtures(t, "en.PastTime", w, fixt)
}
