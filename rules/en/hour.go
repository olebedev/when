package en

import (
	"regexp"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/olebedev/when/rules"
)

/*
	"5pm"
	"5 pm"
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
			"(\\d{1,2})" +
			"(?:\\s*(A\\.|P\\.|A\\.M\\.|P\\.M\\.|AM?|PM?))" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Hour != nil && s != rules.Override {
				return false, nil
			}

			hour, err := strconv.Atoi(m.Captures[0])
			if err != nil {
				return false, errors.Wrap(err, "hour rule")
			}

			if hour > 12 {
				return false, nil
			}

			zero := 0
			switch m.Captures[1][0] {
			case 65, 97: // am
				c.Hour = &hour
			case 80, 112: // pm
				if hour < 12 {
					hour += 12
				}
				c.Hour = &hour
			}

			c.Minute = &zero
			return true, nil
		},
	}
}
