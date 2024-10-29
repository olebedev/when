package ru

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

// https://go.dev/play/p/YsVdaraCwIP

func Date(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\b|^)(\d{1,2})\s*(` + MONTHS_PATTERN + `)\s*(\d{4})(?:\s*в\s*(\d{1,2}):(\d{2}))?(?:\b|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Day != nil || c.Month != nil || c.Year != nil) || s != rules.Override {
				return false, nil
			}

			day, err := strconv.Atoi(m.Captures[0])
			if err != nil {
				return false, errors.Wrap(err, "date rule: day")
			}

			month, ok := MONTHS[strings.ToLower(m.Captures[1])]
			if !ok {
				return false, errors.New("date rule: invalid month")
			}

			year, err := strconv.Atoi(m.Captures[2])
			if err != nil {
				return false, errors.Wrap(err, "date rule: year")
			}

			hour, minute := 0, 0
			if m.Captures[3] != "" && m.Captures[4] != "" {
				hour, err = strconv.Atoi(m.Captures[3])
				if err != nil {
					return false, errors.Wrap(err, "date rule: hour")
				}
				minute, err = strconv.Atoi(m.Captures[4])
				if err != nil {
					return false, errors.Wrap(err, "date rule: minute")
				}
			}

			c.Day = &day
			c.Month = pointerToInt(int(month))
			c.Year = &year
			c.Hour = &hour
			c.Minute = &minute

			return true, nil
		},
	}
}

func pointerToInt(v int) *int {
	return &v
}
