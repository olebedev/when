package br

import (
	"regexp"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
)

func Weekday(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)" +
			"(?:\\W|^)" +
			"(?:n[ao]\\s*?)?" +
			"(?:(nest[ae]|ess[ae]|últim[a|o]|próxim[ao])\\s*)?" +
			"(" + WEEKDAY_OFFSET_PATTERN[3:] + // skip '(?:'
			"(?:\\s*(passad[ao]|que\\svem))?" +
			"(?:\\W|$)",
		),

		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			_ = overwrite

			day := strings.ToLower(strings.TrimSpace(m.Captures[1]))
			norm := strings.ToLower(strings.TrimSpace(m.Captures[0] + m.Captures[2]))

			if norm == "" {
				norm = "próxim[ao]"
			}
			dayInt, ok := WEEKDAY_OFFSET[day]
			if !ok {
				return false, nil
			}

			if c.Duration != 0 && !overwrite {
				return false, nil
			}

			// Switch:
			switch {
			case strings.Contains(norm, "passad") || strings.Contains(norm, "últim"):
				diff := int(ref.Weekday()) - dayInt
				if diff > 0 {
					c.Duration = -time.Duration(diff*24) * time.Hour
				} else if diff < 0 {
					c.Duration = -time.Duration(7+diff) * 24 * time.Hour
				} else {
					c.Duration = -(7 * 24 * time.Hour)
				}
			case strings.Contains(norm, "próxim"), strings.Contains(norm, "que vem"):
				diff := dayInt - int(ref.Weekday())
				if diff > 0 {
					c.Duration = time.Duration(diff*24) * time.Hour
				} else if diff < 0 {
					c.Duration = time.Duration(7+diff) * 24 * time.Hour
				} else {
					c.Duration = 7 * 24 * time.Hour
				}
			case strings.Contains(norm, "nest"), strings.Contains(norm, "ess"):
				if int(ref.Weekday()) < dayInt {
					diff := dayInt - int(ref.Weekday())
					if diff > 0 {
						c.Duration = time.Duration(diff*24) * time.Hour
					} else if diff < 0 {
						c.Duration = time.Duration(7+diff) * 24 * time.Hour
					} else {
						c.Duration = 7 * 24 * time.Hour
					}
				} else if int(ref.Weekday()) > dayInt {
					diff := int(ref.Weekday()) - dayInt
					if diff > 0 {
						c.Duration = -time.Duration(diff*24) * time.Hour
					} else if diff < 0 {
						c.Duration = -time.Duration(7+diff) * 24 * time.Hour
					} else {
						c.Duration = -(7 * 24 * time.Hour)
					}
				}
			}

			return true, nil
		},
	}
}
