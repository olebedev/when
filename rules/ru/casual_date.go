package ru

import (
	"regexp"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
)

// https://play.golang.org/p/D19wPQBraq

func CasualDate(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(сейчас|сегодня|завтра|вчера)" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			lower := strings.ToLower(strings.TrimSpace(m.String()))

			switch {
			// case strings.Contains(lower, "вечер"):
			// 	if c.Hour == nil || s == rules.OverWrite {
			// 		c.Hour = pointer.ToInt(23)
			// 		c.Minute = pointer.ToInt(0)
			// 	}
			case strings.Contains(lower, "сегодня"):
				// c.Hour = pointer.ToInt(18)
			case strings.Contains(lower, "завтра"):
				if c.Duration == 0 || s == rules.OverWrite {
					c.Duration += time.Hour * 24
				}
			case strings.Contains(lower, "вчера"):
				if c.Duration == 0 || s == rules.OverWrite {
					c.Duration -= time.Hour * 24
				}
				// case strings.Contains(lower, "last night"):
				// 	if (c.Hour == nil && c.Duration == 0) || overwrite {
				// 		c.Hour = pointer.ToInt(23)
				// 		c.Duration -= time.Hour * 24
				// 	}
			}

			return true, nil
		},
	}
}
