package en

import (
	"regexp"
	"strconv"
	"time"

	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

/*
	{"05:30:10", 0, "05:30:10", 0},
	{"05-30-10", 0, "05-30-10", 0},
	{"05.30.10", 0, "05.30.10", 0},
*/

func HourMinuteSecond(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"((?:[0-1]{0,1}[0-9])|(?:2[0-3]))" +
			"(?:\\:|：|\\-)" +
			"((?:[0-5][0-9]))" +
			"(?:\\:|：|\\-)" +
			"((?:[0-5][0-9]))" +
			"(?:\\s*(A\\.|P\\.|A\\.M\\.|P\\.M\\.|AM?|PM?))?" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Hour != nil || c.Minute != nil || c.Second != nil) && s != rules.Override {
				return false, nil
			}

			hour, err := strconv.Atoi(m.Captures[0])
			if err != nil {
				return false, errors.Wrap(err, "hour minute second rule")
			}

			minutes, err := strconv.Atoi(m.Captures[1])
			if err != nil {
				return false, errors.Wrap(err, "hour minute second rule")
			}

			if minutes > 59 {
				return false, nil
			}
			c.Minute = &minutes

			seconds, err := strconv.Atoi(m.Captures[2])
			if err != nil {
				return false, errors.Wrap(err, "hour minute second rule")
			}

			if seconds > 59 {
				return false, nil
			}
			c.Second = &seconds

			if m.Captures[3] != "" {
				if hour > 12 {
					return false, nil
				}
				switch m.Captures[3][0] {
				case 65, 97: // am
					c.Hour = &hour
				case 80, 112: // pm
					if hour < 12 {
						hour += 12
					}
					c.Hour = &hour
				}
			} else {
				if hour > 23 {
					return false, nil
				}
				c.Hour = &hour
			}

			return true, nil
		},
	}
}
