package br

import (
	"regexp"
	"strings"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
)

func CasualDate(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)(agora|hoje|(?:de\\s|nesta\\s|esta\\s)noite|última(?:s|)\\s*noite|(?:amanhã|ontem)\\s*|amanhã|ontem)(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			lower := strings.ToLower(strings.TrimSpace(m.String()))

			switch {
			case regexContains("(nesta|esta|hoje)(\\s|\\s([aà]|de)\\s)noite", lower):
				if c.Hour == nil && c.Minute == nil || overwrite {
					c.Hour = pointer.ToInt(23)
					c.Minute = pointer.ToInt(0)
				}
			case strings.Contains(lower, "hoje"):
				// c.Hour = pointer.ToInt(18)
			case strings.Contains(lower, "amanhã"):
				if c.Duration == 0 || overwrite {
					c.Duration += time.Hour * 24
				}
			case strings.Contains(lower, "ontem"):
				if c.Duration == 0 || overwrite {
					c.Duration -= time.Hour * 24
				}
			case regexContains("(ontem|última)(\\s|\\s([aà]|de)\\s)noite", lower):
				if (c.Hour == nil && c.Duration == 0) || overwrite {
					c.Hour = pointer.ToInt(23)
					c.Duration -= time.Hour * 24
				}
			}

			return true, nil
		},
	}
}

func regexContains(regex string, text string) bool {
	contains, _ := regexp.MatchString(regex, text)

	return contains
}
