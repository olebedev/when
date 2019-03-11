package br

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
			"(?i)(?:\\W|^)(dentro\\sde|em)\\s*" +
				"(?:(" + INTEGER_WORDS_PATTERN + "|[0-9]+|(?:\\s*pouc[oa](?:s|)?|algu(?:mas|m|ns)?|mei[oa]?))\\s*" +
				"(segundos?|min(?:uto)?s?|horas?|dias?|semanas?|mês|meses|anos?)\\s*)" +
				"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			numStr := strings.TrimSpace(m.Captures[1])

			var num int
			var err error

			if n, ok := INTEGER_WORDS[numStr]; ok {
				num = n
			} else if strings.Contains(numStr, "pouc") || strings.Contains(numStr, "algu") {
				num = 3
			} else if strings.Contains(numStr, "mei") {
				// pass
			} else {
				num, err = strconv.Atoi(numStr)
				if err != nil {
					return false, errors.Wrapf(err, "convert '%s' to int", numStr)
				}
			}

			exponent := strings.TrimSpace(m.Captures[2])

			if !strings.Contains(numStr, "mei") {
				switch {
				case strings.Contains(exponent, "segundo"):
					if c.Duration == 0 || overwrite {
						c.Duration = time.Duration(num) * time.Second
					}
				case strings.Contains(exponent, "min"):
					if c.Duration == 0 || overwrite {
						c.Duration = time.Duration(num) * time.Minute
					}
				case strings.Contains(exponent, "hora"):
					if c.Duration == 0 || overwrite {
						c.Duration = time.Duration(num) * time.Hour
					}
				case strings.Contains(exponent, "dia"):
					if c.Duration == 0 || overwrite {
						c.Duration = time.Duration(num) * 24 * time.Hour
					}
				case strings.Contains(exponent, "semana"):
					if c.Duration == 0 || overwrite {
						c.Duration = time.Duration(num) * 7 * 24 * time.Hour
					}
				case strings.Contains(exponent, "mês"), strings.Contains(exponent, "meses"):
					if c.Month == nil || overwrite {
						c.Month = pointer.ToInt((int(ref.Month()) + num) % 12)
					}
				case strings.Contains(exponent, "ano"):
					if c.Year == nil || overwrite {
						c.Year = pointer.ToInt(ref.Year() + num)
					}
				}
			} else {
				switch {
				case strings.Contains(exponent, "hora"):
					if c.Duration == 0 || overwrite {
						c.Duration = 30 * time.Minute
					}
				case strings.Contains(exponent, "dia"):
					if c.Duration == 0 || overwrite {
						c.Duration = 12 * time.Hour
					}
				case strings.Contains(exponent, "semana"):
					if c.Duration == 0 || overwrite {
						c.Duration = 7 * 12 * time.Hour
					}
				case strings.Contains(exponent, "mês"), strings.Contains(exponent, "meses"):
					if c.Duration == 0 || overwrite {
						// 2 weeks
						c.Duration = 14 * 24 * time.Hour
					}
				case strings.Contains(exponent, "ano"):
					if c.Month == nil || overwrite {
						c.Month = pointer.ToInt((int(ref.Month()) + 6) % 12)
					}
				}
			}

			return true, nil
		},
	}
}
