package ru

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
)

// <[]string{"3 марта", "3", "марта"}>
// <[]string{"21 фев", "21", "фев"}>
// <[]string{"15 августа", "15", "августа"}>
// <[]string{"31 нояб", "31", "ноя"}>
// <[]string{"31 ноя.", "31", "ноя."}>

// https://play.golang.org/p/-Q3yzGily-1

// 1. - ordinal day?
// 2. - numeric day?
// 3. - month

func ExactMonthDate(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)" +
			"(?:\\W|^)" +
			"(?:([0-9]+)\\s*)?" +
			"(" + MONTH_OFFSET_PATTERN[3:] +
			"(?:\\W|$)",
		),

		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			_ = overwrite

			num1 := strings.ToLower(strings.TrimSpace(m.Captures[0]))
			mon := strings.ToLower(strings.TrimSpace(m.Captures[1]))

			monInt, ok := MONTH_OFFSET[mon]
			if !ok {
				return false, nil
			}

			c.Month = &monInt

			if num1 != "" {
				n, err := strconv.ParseInt(num1, 10, 8)
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
