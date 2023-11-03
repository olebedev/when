package nl

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

func Deadline(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile(
			"(?i)(?:\\W|^)(binnen|in|over|na)\\s*" +
				"(" + INTEGER_WORDS_PATTERN + "|[0-9]+|een(?:\\s*(paar|half|halve))?)\\s*" +
				"(seconden?|minuut|minuten|uur|uren|dag|dagen|week|weken|maand|maanden|jaar|jaren)\\s*" +
				"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			numStr := strings.TrimSpace(m.Captures[1])

			var num int
			var err error

			if n, ok := INTEGER_WORDS[numStr]; ok {
				num = n
			} else if numStr == "een" {
				num = 1
			} else if strings.Contains(numStr, "paar") {
				num = 3
			} else if strings.Contains(numStr, "half") || strings.Contains(numStr, "halve") {
				// pass
			} else {
				num, err = strconv.Atoi(numStr)
				if err != nil {
					return false, errors.Wrapf(err, "convert '%s' to int", numStr)
				}
			}

			exponent := strings.TrimSpace(m.Captures[3])

			if !strings.Contains(numStr, "half") && !strings.Contains(numStr, "halve") {
				switch {
				case strings.Contains(exponent, "second"):
					if c.Duration == 0 || overwrite {
						c.Duration = time.Duration(num) * time.Second
					}
				case strings.Contains(exponent, "min"):
					if c.Duration == 0 || overwrite {
						c.Duration = time.Duration(num) * time.Minute
					}
				case strings.Contains(exponent, "uur"), strings.Contains(exponent, "uren"):
					if c.Duration == 0 || overwrite {
						c.Duration = time.Duration(num) * time.Hour
					}
				case strings.Contains(exponent, "dag"):
					if c.Duration == 0 || overwrite {
						c.Duration = time.Duration(num) * 24 * time.Hour
					}
				case strings.Contains(exponent, "week"), strings.Contains(exponent, "weken"):
					if c.Duration == 0 || overwrite {
						c.Duration = time.Duration(num) * 7 * 24 * time.Hour
					}
				case strings.Contains(exponent, "maand"):
					if c.Month == nil || overwrite {
						c.Month = pointer.ToInt((int(ref.Month()) + num) % 12)
					}
				case strings.Contains(exponent, "jaar"):
					if c.Year == nil || overwrite {
						c.Year = pointer.ToInt(ref.Year() + num)
					}
				}
			} else {
				switch {
				case strings.Contains(exponent, "uur"):
					if c.Duration == 0 || overwrite {
						c.Duration = 30 * time.Minute
					}
				case strings.Contains(exponent, "dag"):
					if c.Duration == 0 || overwrite {
						c.Duration = 12 * time.Hour
					}
				case strings.Contains(exponent, "week"):
					if c.Duration == 0 || overwrite {
						c.Duration = 7 * 12 * time.Hour
					}
				case strings.Contains(exponent, "maand"):
					if c.Duration == 0 || overwrite {
						// 2 weeks
						c.Duration = 14 * 24 * time.Hour
					}
				case strings.Contains(exponent, "jaar"):
					if c.Month == nil || overwrite {
						c.Month = pointer.ToInt((int(ref.Month()) + 6) % 12)
					}
				}
			}

			return true, nil
		},
	}
}
