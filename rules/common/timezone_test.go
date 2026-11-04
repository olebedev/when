package common_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"github.com/stretchr/testify/require"
)

func TestTimezone(t *testing.T) {
	fixt := []Fixture{
		{"call me at 10 pm EST", 11, "10 pm EST", 27 * time.Hour},
		{"call me at 10 pm PDT", 11, "10 pm PDT", 29 * time.Hour},
		{"call me at 10 pm CEST", 11, "10 pm CEST", 20 * time.Hour},
		{"at 22:00 UTC", 3, "22:00 UTC", 22 * time.Hour},
		{"at 22:00 utc+2", 3, "22:00 utc+2", 20 * time.Hour},
		{"at 22:00 GMT-05:30", 3, "22:00 GMT-05:30", 27*time.Hour + 30*time.Minute},
		{"at 22:00 GMT+0530", 3, "22:00 GMT+0530", 16*time.Hour + 30*time.Minute},

		// abbreviations only in upper case, "est" is left alone
		{"call me at 10 pm est", 11, "10 pm", 22 * time.Hour},
	}

	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.Timezone(rules.Override))

	ApplyFixtures(t, "common.Timezone", w, fixt)

	nilFixt := []Fixture{
		{"the EST office", 0, "a timezone alone is not a time", 0},
	}
	ApplyFixturesNil(t, "common.Timezone nil", w, nilFixt)

	res, err := w.Parse("call me at 10 pm EST", null)
	require.Nil(t, err)
	name, offset := res.Time.Zone()
	require.Equal(t, "EST", name)
	require.Equal(t, -5*60*60, offset)
}
