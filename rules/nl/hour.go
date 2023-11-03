package nl

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/olebedev/when/rules"
)

/*
	"5u"
	"5 uur"
	"5am"
	"5pm"
	"5A."
	"5P."
	"11 P.M."
	https://play.golang.org/p/2Gh35Sl3KP
*/

func Hour(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(?:\\s*((om)?))" +
			"(\\d{1,2})" +
			"(?:\\s*(U\\.?|UUR|A\\.|P\\.|A\\.M\\.|P\\.M\\.|AM?|PM?))" +
			"(?:\\s*((in de|\\'s) (middags?|avonds?))?)" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Hour != nil && s != rules.Override {
				return false, nil
			}

			lower := strings.ToLower(strings.TrimSpace(m.String()))
			hour, err := strconv.Atoi(m.Captures[2])

			if err != nil {
				return false, errors.Wrap(err, "hour rule")
			}

			zero := 0

			if hour > 23 {
				return false, nil
			}
			c.Hour = &hour

			// pm
			if regexp.MustCompile("p.?(m.?)?").MatchString(strings.ToLower(strings.TrimSpace(m.Captures[3]))) {
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

			c.Minute = &zero
			c.Second = &zero
			return true, nil
		},
	}
}
