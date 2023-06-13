package zh

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
)

func CasualDate(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)" +
			"(大前|前|昨|今天|今|明|大后|后|下下|下|上|上上)" + "(天|月|个月|年|儿)" +
			"(1[0-9]|2[0-9]|3[0-1]|[1-9]|" + DAY_WORDS_PATTERN + ")?" + "(?:\\s*)?" +
			"(日|号)?" +
			"",
		// "(?:\\W|$)"，
		),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			lower := compressStr(strings.TrimSpace(m.String()))

			switch {
			case strings.Contains(lower, "号"), strings.Contains(lower, "日"):
				day, _ := strconv.Atoi(m.Captures[2])
				c.Day = pointer.ToInt(day)
			}

			switch {

			case strings.Contains(lower, "后年"):
				c.Year = pointer.ToInt(ref.Year() + 2)
			case strings.Contains(lower, "明年"):
				c.Year = pointer.ToInt(ref.Year() + 1)
			case strings.Contains(lower, "下下"):
				monthInt := int(ref.Month()) + 2
				c.Month = pointer.ToInt(monthInt)
			case strings.Contains(lower, "下月"), strings.Contains(lower, "下个月"):
				monthInt := int(ref.Month()) + 1
				c.Month = pointer.ToInt(monthInt)
			case strings.Contains(lower, "上上"):
				monthInt := int(ref.Month()) - 2
				c.Month = pointer.ToInt(monthInt)
			case strings.Contains(lower, "上月"), strings.Contains(lower, "上个月"):
				monthInt := int(ref.Month()) - 1
				c.Month = pointer.ToInt(monthInt)
			case strings.Contains(lower, "今晚"), strings.Contains(lower, "晚上"):
				if c.Hour == nil && c.Minute == nil || overwrite {
					c.Hour = pointer.ToInt(22)
					c.Minute = pointer.ToInt(0)
				}
			case strings.Contains(lower, "今天"), strings.Contains(lower, "今儿"):
				// c.Hour = pointer.ToInt(18)
			case strings.Contains(lower, "明天"), strings.Contains(lower, "明儿"):
				if c.Duration == 0 || overwrite {
					c.Duration += time.Hour * 24
				}
			case strings.Contains(lower, "昨天"):
				if c.Duration == 0 || overwrite {
					c.Duration -= time.Hour * 24
				}
			case strings.Contains(lower, "大前天"):
				if c.Duration == 0 || overwrite {
					c.Duration -= time.Hour * 24 * 3
				}
			case strings.Contains(lower, "前天"):
				if c.Duration == 0 || overwrite {
					c.Duration -= time.Hour * 24 * 2
				}
			case strings.Contains(lower, "昨晚"):
				if (c.Hour == nil && c.Duration == 0) || overwrite {
					c.Hour = pointer.ToInt(23)
					c.Duration -= time.Hour * 24
				}
			case strings.Contains(lower, "大后天"):
				if c.Duration == 0 || overwrite {
					c.Duration += time.Hour * 24 * 3
				}
			case strings.Contains(lower, "后天"):
				if c.Duration == 0 || overwrite {
					c.Duration += time.Hour * 24 * 2
				}
			}

			return true, nil
		},
	}
}
