package en

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

func PastTime(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile(
			"(?i)(?:\\W|^)\\s*" +
				"(" + INTEGER_WORDS_PATTERN + "|[0-9]+|an?(?:\\s*few)?|half(?:\\s*an?)?)\\s*" +
				"(seconds?|min(?:ute)?s?|hours?|days?|weeks?|months?|years?) (ago)\\s*" +
				"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {

			numStr := strings.TrimSpace(m.Captures[0])

			var num int
			var err error

			if n, ok := INTEGER_WORDS[numStr]; ok {
				num = n
			} else if numStr == "a" || numStr == "an" {
				num = 1
			} else if strings.Contains(numStr, "few") {
				num = 3
			} else if strings.Contains(numStr, "half") {
				// pass
			} else {
				num, err = strconv.Atoi(numStr)
				if err != nil {
					return false, errors.Wrapf(err, "convert '%s' to int", numStr)
				}
			}

			exponent := strings.TrimSpace(m.Captures[1])

			if !strings.Contains(numStr, "half") {
				switch {
				case strings.Contains(exponent, "second"):
					if c.Duration == 0 || overwrite {
						c.Duration = -(time.Duration(num) * time.Second)
					}
				case strings.Contains(exponent, "min"):
					if c.Duration == 0 || overwrite {
						c.Duration = -(time.Duration(num) * time.Minute)
					}
				case strings.Contains(exponent, "hour"):
					if c.Duration == 0 || overwrite {
						c.Duration = -(time.Duration(num) * time.Hour)
					}
				case strings.Contains(exponent, "day"):
					if c.Duration == 0 || overwrite {
						c.Duration = -(time.Duration(num) * 24 * time.Hour)
					}
				case strings.Contains(exponent, "week"):
					if c.Duration == 0 || overwrite {
						c.Duration = -(time.Duration(num) * 7 * 24 * time.Hour)
					}
				case strings.Contains(exponent, "month"):
					if c.Month == nil || overwrite {
						c.Month = pointer.ToInt((int(ref.Month()) - num) % 12)
					}
				case strings.Contains(exponent, "year"):
					if c.Year == nil || overwrite {
						c.Year = pointer.ToInt(ref.Year() - num)
					}
				}
			} else {
				switch {
				case strings.Contains(exponent, "hour"):
					if c.Duration == 0 || overwrite {
						c.Duration = -(30 * time.Minute)
					}
				case strings.Contains(exponent, "day"):
					if c.Duration == 0 || overwrite {
						c.Duration = -(12 * time.Hour)
					}
				case strings.Contains(exponent, "week"):
					if c.Duration == 0 || overwrite {
						c.Duration = -(7 * 12 * time.Hour)
					}
				case strings.Contains(exponent, "month"):
					if c.Duration == 0 || overwrite {
						// 2 weeks
						c.Duration = -(14 * 24 * time.Hour)
					}
				case strings.Contains(exponent, "year"):
					if c.Month == nil || overwrite {
						c.Month = pointer.ToInt((int(ref.Month()) - 6) % 12)
					}
				}
			}

			return true, nil
		},
	}
}
