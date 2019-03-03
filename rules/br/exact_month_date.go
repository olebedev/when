package br

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
)

func ExactMonthDate(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)" +
			"(?:\\W|^)" +
			"(?:(?:(\\d{1,2})|(" + ORDINAL_WORDS_PATTERN[3:] + // skip '(?:'
			")(?:\\sdia\\sde\\s|\\sde\\s|\\s))*" +
			"(" + MONTH_OFFSET_PATTERN[3:] + // skip '(?:'
			"(?:\\W|$)",
		),

		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			_ = overwrite

			num := strings.ToLower(strings.TrimSpace(m.Captures[0]))
			ord := strings.ToLower(strings.TrimSpace(m.Captures[1]))
			mon := strings.ToLower(strings.TrimSpace(m.Captures[2]))

			monInt, ok := MONTH_OFFSET[mon]
			if !ok {
				return false, nil
			}

			c.Month = &monInt

			if ord != "" {
				ordInt, ok := ORDINAL_WORDS[ord]
				if !ok {
					return false, nil
				}

				c.Day = &ordInt
			}

			if num != "" {
				n, err := strconv.ParseInt(num, 10, 8)
				if err != nil {
					return false, nil
				}

				num := int(n)

				c.Day = &num
			}

			return true, nil
		},
	}
}
