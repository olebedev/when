package nl

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
		RegExp: regexp.MustCompile(`(?i)(?:\W|^)((deze|tussen de |maandag|dinsdag|woensdag|donderdag|vrijdag|zaterdag|zondag| )\s*(ochtend|morgen|middag|avond))`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {

			lower := strings.ToLower(strings.TrimSpace(m.String()))

			if (c.Weekday != nil || c.Hour != nil || c.Minute != nil) && !overwrite {
				return false, nil
			}

			if regexp.MustCompile("(maandag|dinsdag|woensdag|donderdag|vrijdag|zaterdag|zondag)").MatchString(lower) {
				weekday := -1

				switch {
				case strings.Contains(lower, "maandag"):
					weekday = 1
				case strings.Contains(lower, "dinsdag"):
					weekday = 2
				case strings.Contains(lower, "woensdag"):
					weekday = 3
				case strings.Contains(lower, "donderdag"):
					weekday = 4
				case strings.Contains(lower, "vrijdag"):
					weekday = 5
				case strings.Contains(lower, "zaterdag"):
					weekday = 6
				case strings.Contains(lower, "zondag"):
					weekday = 7
				}

				if weekday != -1 {
					c.Duration += time.Hour * 24 * time.Duration((weekday+7-(int(ref.Weekday())))%7)
				}
			}

			switch {
			case strings.Contains(lower, "middag") && !strings.Contains(lower, "tussen de middag"):
				if o.Afternoon != 0 {
					c.Hour = &o.Afternoon
				} else {
					c.Hour = pointer.ToInt(15)
				}
				c.Minute = pointer.ToInt(0)
			case strings.Contains(lower, "avond"):
				if o.Evening != 0 {
					c.Hour = &o.Evening
				} else {
					c.Hour = pointer.ToInt(18)
				}
				c.Minute = pointer.ToInt(0)
			case strings.Contains(lower, "ochtend"), strings.Contains(lower, "morgen"):
				if o.Morning != 0 {
					c.Hour = &o.Morning
				} else {
					c.Hour = pointer.ToInt(8)
				}
				c.Minute = pointer.ToInt(0)
			case strings.Contains(lower, "tussen de middag"):
				c.Hour = pointer.ToInt(12)
				c.Minute = pointer.ToInt(0)
			}

			return true, nil
		},
	}
}
