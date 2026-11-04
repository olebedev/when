package common

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
)

/*
- 10 pm EST
- 22:00 UTC
- 9am GMT+2
- 18:30 utc-05:00
- 18:30 GMT+0530

Sets the location of the time found by the other rules. A timezone on its
own is not a time, so the rule never yields a result by itself.

Abbreviations are matched in upper case only, so "est" in a sentence is
not taken for a timezone, and each one is a fixed offset from UTC.
*/

var timezoneOffsets = map[string]time.Duration{
	"HST":  -10 * time.Hour,
	"AKST": -9 * time.Hour,
	"AKDT": -8 * time.Hour,
	"PST":  -8 * time.Hour,
	"PDT":  -7 * time.Hour,
	"MST":  -7 * time.Hour,
	"MDT":  -6 * time.Hour,
	"CST":  -6 * time.Hour,
	"CDT":  -5 * time.Hour,
	"EST":  -5 * time.Hour,
	"EDT":  -4 * time.Hour,
	"BST":  1 * time.Hour,
	"CET":  1 * time.Hour,
	"CEST": 2 * time.Hour,
	"EET":  2 * time.Hour,
	"EEST": 3 * time.Hour,
	"JST":  9 * time.Hour,
	"AEST": 10 * time.Hour,
	"AEDT": 11 * time.Hour,
}

func Timezone(s rules.Strategy) rules.Rule {
	abbreviations := make([]string, 0, len(timezoneOffsets))
	for abbreviation := range timezoneOffsets {
		abbreviations = append(abbreviations, abbreviation)
	}
	sort.Strings(abbreviations)

	return &rules.F{
		RegExp: regexp.MustCompile("(?:\\W|^)" +
			"((?i:utc|gmt)(?:([+-])(\\d{1,2})(?::?(\\d{2}))?)?|" +
			strings.Join(abbreviations, "|") + ")" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Location != nil && s != rules.Override {
				return false, nil
			}

			name := strings.ToUpper(m.Captures[0])
			offset, ok := timezoneOffsets[name]
			if !ok {
				hours, _ := strconv.Atoi(m.Captures[2])
				minutes, _ := strconv.Atoi(m.Captures[3])
				if hours > 14 || minutes > 59 {
					return false, nil
				}
				offset = time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute
				if m.Captures[1] == "-" {
					offset = -offset
				}
			}

			c.Location = time.FixedZone(name, int(offset.Seconds()))
			return false, nil
		},
	}
}
