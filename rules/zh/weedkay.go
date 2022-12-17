package zh

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
			"(?:(本|这|下|上|这个|下个|上个|下下)\\s*)?" +
			"(?:(周|礼拜|星期)\\s*)" +
			"(1|2|3|4|5|6|天|一|二|三|四|五|六|日)" +
			"(?:\\W|$)",
		),

		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			_ = overwrite

			if strings.TrimSpace(m.Captures[1]) == "" {
				return false, nil
			}

			day := strings.ToLower(strings.TrimSpace(m.Captures[2]))
			norm := strings.ToLower(strings.TrimSpace(m.Captures[0]))
			if norm == "" {
				norm = "本"
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
			case strings.Contains(norm, "上"):
				diff := int(ref.Weekday()) - dayInt
				c.Duration = -time.Duration(7+diff) * 24 * time.Hour
			case strings.Contains(norm, "下下"):
				diff := dayInt - int(ref.Weekday())
				c.Duration = time.Duration(7+7+diff) * 24 * time.Hour
			case strings.Contains(norm, "下"):
				diff := dayInt - int(ref.Weekday())
				c.Duration = time.Duration(7+diff) * 24 * time.Hour
			case strings.Contains(norm, "本") || strings.Contains(norm, "这"):
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
