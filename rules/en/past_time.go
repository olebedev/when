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
	merge := s == rules.Merge
	skip := s == rules.Skip

	return &rules.F{
		RegExp: regexp.MustCompile(
			"(?i)(?:\\W|^)\\s*" +
				"(" + INTEGER_WORDS_PATTERN + "|[0-9]+|an?(?:\\s*few)?|half(?:\\s*an?)?)\\s*" +
				"(seconds?|min(?:ute)?s?|hours?|days?|weeks?|months?|years?) (ago)\\s*" +
				"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Duration != 0 && skip {
				return false, nil
			}

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

			var duration time.Duration
			if !strings.Contains(numStr, "half") {
				switch {
				case strings.Contains(exponent, "second"):
					duration = -(time.Duration(num) * time.Second)
				case strings.Contains(exponent, "min"):
					duration = -(time.Duration(num) * time.Minute)
				case strings.Contains(exponent, "hour"):
					duration = -(time.Duration(num) * time.Hour)
				case strings.Contains(exponent, "day"):
					duration = -(time.Duration(num) * 24 * time.Hour)
				case strings.Contains(exponent, "week"):
					duration = -(time.Duration(num) * 7 * 24 * time.Hour)
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
					duration = -(30 * time.Minute)
				case strings.Contains(exponent, "day"):
					duration = -(12 * time.Hour)
				case strings.Contains(exponent, "week"):
					duration = -(7 * 12 * time.Hour)
				case strings.Contains(exponent, "month"):
					// 2 weeks
					duration = -(14 * 24 * time.Hour)
				case strings.Contains(exponent, "year"):
					if c.Month == nil || overwrite {
						c.Month = pointer.ToInt((int(ref.Month()) - 6) % 12)
					}
				}
			}

			if overwrite {
				c.Duration = duration
			} else if merge {
				c.Duration = c.Duration + duration
			}

			return true, nil
		},
	}
}
