package en

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
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)(now|today|tonight|last\\s*night|(?:tomorrow|tmr|yesterday)\\s*|tomorrow|tmr|yesterday)(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			lower := strings.ToLower(strings.TrimSpace(m.String()))

			switch {
			case strings.Contains(lower, "tonight"):
				if c.Hour == nil && c.Minute == nil || overwrite {
					c.Hour = pointer.ToInt(23)
					c.Minute = pointer.ToInt(0)
				}
			case strings.Contains(lower, "today"):
				// c.Hour = pointer.ToInt(18)
			case strings.Contains(lower, "tomorrow"), strings.Contains(lower, "tmr"):
				if c.Duration == 0 || overwrite {
					c.Duration += time.Hour * 24
				}
			case strings.Contains(lower, "yesterday"):
				if c.Duration == 0 || overwrite {
					c.Duration -= time.Hour * 24
				}
			case strings.Contains(lower, "last night"):
				if (c.Hour == nil && c.Duration == 0) || overwrite {
					c.Hour = pointer.ToInt(23)
					c.Duration -= time.Hour * 24
				}
			}

			return true, nil
		},
	}
}
