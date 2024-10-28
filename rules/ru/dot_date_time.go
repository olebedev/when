package ru

import (
	"regexp"
	"strconv"
	"time"

	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

// https://go.dev/play/p/vRzLhHHupUJ

func DotDateTime(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:^|\b)(\d{2})\.(\d{2})\.(\d{4})\s+(\d{2}):(\d{2})(?:\b|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Day != nil || c.Month != nil || c.Year != nil {
				return false, nil
			}

			day, err := strconv.Atoi(m.Captures[0])
			if err != nil {
				return false, errors.Wrap(err, "dot date time rule: day")
			}

			month, err := strconv.Atoi(m.Captures[1])
			if err != nil {
				return false, errors.Wrap(err, "dot date time rule: month")
			}

			year, err := strconv.Atoi(m.Captures[2])
			if err != nil {
				return false, errors.Wrap(err, "dot date time rule: year")
			}

			hour, err := strconv.Atoi(m.Captures[3])
			if err != nil {
				return false, errors.Wrap(err, "dot date time rule: hour")
			}

			minute, err := strconv.Atoi(m.Captures[4])
			if err != nil {
				return false, errors.Wrap(err, "dot date time rule: minute")
			}

			if day > 0 && day <= 31 && month > 0 && month <= 12 {
				c.Day = &day
				c.Month = &month
				c.Year = &year
				c.Hour = &hour
				c.Minute = &minute
				return true, nil
			}

			return false, nil
		},
	}
}
