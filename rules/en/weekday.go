package en

import (
	"regexp"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
)

func Weekday(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override
	merge := s == rules.Merge
	skip := s == rules.Skip

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)" +
			"(?:\\W|^)" +
			"(?:on\\s*?)?" +
			"(?:(this|last|past|next)\\s*)?" +
			"(" + WEEKDAY_OFFSET_PATTERN[3:] + // skip '(?:'
			"(?:\\s*(this|last|past|next)\\s*week)?" +
			"(?:\\W|$)",
		),

		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			_ = overwrite

			day := strings.ToLower(strings.TrimSpace(m.Captures[1]))
			norm := strings.ToLower(strings.TrimSpace(m.Captures[0] + m.Captures[2]))
			if norm == "" {
				norm = "next"
			}
			dayInt, ok := WEEKDAY_OFFSET[day]
			if !ok {
				return false, nil
			}

			if c.Duration != 0 && skip {
				return false, nil
			}

			var duration time.Duration
			// Switch:
			switch {
			case strings.Contains(norm, "past") || strings.Contains(norm, "last"):
				diff := int(ref.Weekday()) - dayInt
				if diff > 0 {
					duration = -time.Duration(diff*24) * time.Hour
				} else if diff < 0 {
					duration = -time.Duration(7+diff) * 24 * time.Hour
				} else {
					duration = -(7 * 24 * time.Hour)
				}
			case strings.Contains(norm, "next"):
				diff := dayInt - int(ref.Weekday())
				if diff > 0 {
					duration = time.Duration(diff*24) * time.Hour
				} else if diff < 0 {
					duration = time.Duration(7+diff) * 24 * time.Hour
				} else {
					duration = 7 * 24 * time.Hour
				}
			case strings.Contains(norm, "this"):
				if int(ref.Weekday()) < dayInt {
					diff := dayInt - int(ref.Weekday())
					if diff > 0 {
						duration = time.Duration(diff*24) * time.Hour
					} else if diff < 0 {
						duration = time.Duration(7+diff) * 24 * time.Hour
					} else {
						duration = 7 * 24 * time.Hour
					}
				} else if int(ref.Weekday()) > dayInt {
					diff := int(ref.Weekday()) - dayInt
					if diff > 0 {
						duration = -time.Duration(diff*24) * time.Hour
					} else if diff < 0 {
						duration = -time.Duration(7+diff) * 24 * time.Hour
					} else {
						duration = -(7 * 24 * time.Hour)
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
