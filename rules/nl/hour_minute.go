package nl

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

/*
	{"17:30", 0, "17:30", 0},

	https://play.golang.org/p/hXl7C8MWNr
*/

// 1. - int
// 2. - int
// 3. - ext?

func HourMinute(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(?:\\s*((om)?))" +
			"((?:[0-1]{0,1}[0-9])|(?:2[0-3]))" +
			"(?:\\:|：)" +
			"((?:[0-5][0-9]))" +
			"(?:\\s*(U\\.?|UUR|A\\.|P\\.|A\\.M\\.|P\\.M\\.|AM?|PM?))?" +
			"(?:\\s*((in de|\\'s) (middags?|avonds?))?)" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Hour != nil || c.Minute != nil) && s != rules.Override {
				return false, nil
			}

			lower := strings.ToLower(strings.TrimSpace(m.String()))
			hour, err := strconv.Atoi(m.Captures[2])
			if err != nil {
				return false, errors.Wrap(err, "hour minute rule")
			}

			minutes, err := strconv.Atoi(m.Captures[3])
			if err != nil {
				return false, errors.Wrap(err, "hour minute rule")
			}

			if minutes > 59 {
				return false, nil
			}
			c.Minute = &minutes

			if hour > 23 {
				return false, nil
			}
			c.Hour = &hour

			// pm
			if regexp.MustCompile("p.?(m.?)?").MatchString(strings.ToLower(strings.TrimSpace(m.Captures[4]))) {
				if hour < 12 {
					hour += 12
				}

				c.Hour = &hour
			}

			// afternoon or evening
			if (strings.Contains(lower, "middag") || strings.Contains(lower, "avond")) && hour < 12 {
				hour += 12
				c.Hour = &hour
			}

			seconds := 0 // Truncate seconds
			c.Second = &seconds

			return true, nil
		},
	}
}
