package ru

import (
	"regexp"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
)

// https://play.golang.org/p/aRWlil_64M

func Weekday(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\P{L}|^)" +
			"(?:(на|во?|ко?|до|эт(?:от|ой|у|а)?|прошл(?:ую|ый|ая)|последн(?:юю|ий|ее|ая)|следующ(?:ую|ее|ая|ий))\\s*)?" +
			"(" + WEEKDAY_OFFSET_PATTERN[3:] + // skip '(?:'
			"(?:\\s*на\\s*(этой|прошлой|следующей)\\s*неделе)?" +
			"(?:\\P{L}|$)"),

		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {

			day := strings.ToLower(strings.TrimSpace(m.Captures[1]))
			norm := m.Captures[2]
			if norm == "" {
				norm = m.Captures[0]
			}
			if norm == "" {
				norm = "следующ"
			}
			norm = strings.ToLower(strings.TrimSpace(norm))

			dayInt, ok := WEEKDAY_OFFSET[day]
			if !ok {
				return false, nil
			}

			if c.Duration != 0 && s != rules.Override {
				return false, nil
			}

			// Switch:
			switch {
			case strings.Contains(norm, "прошл") || strings.Contains(norm, "последн"):
				diff := int(ref.Weekday()) - dayInt
				if diff > 0 {
					c.Duration = -time.Duration(diff*24) * time.Hour
				} else if diff < 0 {
					c.Duration = -time.Duration(7+diff) * 24 * time.Hour
				} else {
					c.Duration = -(7 * 24 * time.Hour)
				}
			case strings.Contains(norm, "следующ"),
				norm == "в",
				norm == "к",
				strings.Contains(norm, "во"),
				strings.Contains(norm, "ко"),
				strings.Contains(norm, "до"):
				diff := dayInt - int(ref.Weekday())
				if diff > 0 {
					c.Duration = time.Duration(diff*24) * time.Hour
				} else if diff < 0 {
					c.Duration = time.Duration(7+diff) * 24 * time.Hour
				} else {
					c.Duration = 7 * 24 * time.Hour
				}
			case strings.Contains(norm, "эт"):
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
