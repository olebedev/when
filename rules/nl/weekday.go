package nl

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
			"(?:op\\s*?)?" +
			"(?:(deze|vorige|vorige week|afgelopen|volgende|volgende week|komende|komende week)\\s*)?" +
			"(" + WEEKDAY_OFFSET_PATTERN[3:] + // skip '(?:'
			"(?:\\s*(deze|vorige|afgelopen|volgende|komende)\\s*week)?" +
			"(?:\\W|$)",
		),

		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			_ = overwrite

			day := strings.ToLower(strings.TrimSpace(m.Captures[1]))
			norm := strings.ToLower(strings.TrimSpace(m.Captures[0] + m.Captures[2]))
			if norm == "" {
				norm = "volgende"
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
			case strings.Contains(norm, "vorige week"):
				if dayInt == 6 {
					dayInt = -1
				}
				diff := int(ref.Weekday()) - dayInt
				if diff != 0 && dayInt <= 0 {
					c.Duration = -time.Duration(diff) * 24 * time.Hour
				} else if diff != 0 {
					c.Duration = -time.Duration(7+diff) * 24 * time.Hour
				} else {
					c.Duration = -(7 * 24 * time.Hour)
				}
			case strings.Contains(norm, "afgelopen") || strings.Contains(norm, "vorige"):
				diff := int(ref.Weekday()) - dayInt
				if diff > 0 {
					c.Duration = -time.Duration(diff*24) * time.Hour
				} else if diff < 0 {
					c.Duration = -time.Duration(7+diff) * 24 * time.Hour
				} else {
					c.Duration = -(7 * 24 * time.Hour)
				}
			case strings.Contains(norm, "volgende week"):
				if dayInt == 0 {
					dayInt = 7
				}
				diff := dayInt - int(ref.Weekday())
				c.Duration = time.Duration(7+diff) * 24 * time.Hour
			case strings.Contains(norm, "volgende"), strings.Contains(norm, "komende"):
				diff := dayInt - int(ref.Weekday())
				if diff > 0 {
					c.Duration = time.Duration(diff) * 24 * time.Hour
				} else if diff < 0 {
					c.Duration = time.Duration(7+diff) * 24 * time.Hour
				} else {
					c.Duration = 7 * 24 * time.Hour
				}
			case strings.Contains(norm, "deze"):
				if int(ref.Weekday()) < dayInt {
					diff := dayInt - int(ref.Weekday())
					if diff > 0 {
						c.Duration = time.Duration(diff) * 24 * time.Hour
					} else if diff < 0 {
						c.Duration = time.Duration(7+diff) * 24 * time.Hour
					} else {
						c.Duration = 7 * 24 * time.Hour
					}
				} else if int(ref.Weekday()) > dayInt {
					diff := int(ref.Weekday()) - dayInt
					if diff > 0 {
						c.Duration = -time.Duration(diff) * 24 * time.Hour
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
