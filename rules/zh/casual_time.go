package zh

import (
	"regexp"
	"strings"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
)

func CasualTime(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|^)((今天)?\s*(早晨|下午|傍晚|中午|晚上))`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {

			lower := strings.ToLower(strings.TrimSpace(m.String()))

			if (c.Hour != nil || c.Minute != nil) && !overwrite {
				return false, nil
			}

			switch {
			case strings.Contains(lower, "晚上"):
				if o.Evening != 0 {
					c.Hour = &o.Evening
				} else {
					c.Hour = pointer.ToInt(20)
				}
				c.Minute = pointer.ToInt(0)
			case strings.Contains(lower, "下午"):
				if o.Afternoon != 0 {
					c.Hour = &o.Afternoon
				} else {
					c.Hour = pointer.ToInt(15)
				}
				c.Minute = pointer.ToInt(0)
			case strings.Contains(lower, "傍晚"):
				if o.Evening != 0 {
					c.Hour = &o.Evening
				} else {
					c.Hour = pointer.ToInt(18)
				}
				c.Minute = pointer.ToInt(0)
			case strings.Contains(lower, "早晨"):
				if o.Morning != 0 {
					c.Hour = &o.Morning
				} else {
					c.Hour = pointer.ToInt(8)
				}
				c.Minute = pointer.ToInt(0)
			case strings.Contains(lower, "中午"):
				if o.Noon != 0 {
					c.Hour = &o.Noon
				} else {
					c.Hour = pointer.ToInt(12)
				}
				c.Minute = pointer.ToInt(0)
			}

			return true, nil
		},
	}
}
