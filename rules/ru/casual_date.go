package ru

import (
	"regexp"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
)

// https://play.golang.org/p/QrFtjmjUoJ

func CasualDate(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\P{L}|^)" +
			"((?:до|прямо)\\s+)?" +
			"(сейчас|сегодня|завтра|вчера)" +
			"(?:\\P{L}|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			lower := strings.ToLower(strings.TrimSpace(m.String()))

			switch {
			case strings.Contains(lower, "сегодня"):
				// c.Hour = pointer.ToInt(18)
			case strings.Contains(lower, "завтра"):
				if c.Duration == 0 || s == rules.Override {
					c.Duration += time.Hour * 24
				}
			case strings.Contains(lower, "вчера"):
				if c.Duration == 0 || s == rules.Override {
					c.Duration -= time.Hour * 24
				}
			}

			return true, nil
		},
	}
}
