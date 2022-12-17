package zh

import (
	"regexp"
	"strconv"
	"time"

	"github.com/olebedev/when/rules"
)

/*
5/五 分钟后
5 小时后
*/

func AfterTime(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)" +
			"((?:[0-9]{0,3}))?" +
			"(" + INTEGER_WORDS_PATTERN[3:] + "?" + "\\s*" +
			"(?:(分|分钟|小时|天|周|月)\\s*)" +
			"(后)" +
			"(?:\\W|$)",
		),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Hour != nil && s != rules.Override {
				return false, nil
			}
			duration, _ := strconv.Atoi(m.Captures[0])

			if d, exist := INTEGER_WORDS[compressStr(m.Captures[1])]; exist {
				duration = d
			}
			if m.Captures[1] == "半" && m.Captures[2] == "小时" {
				c.Duration = time.Minute * time.Duration(30)
				return true, nil
			}

			switch m.Captures[2] {
			case "分钟", "分":
				c.Duration = time.Minute * time.Duration(duration)
			case "小时":
				c.Duration = time.Hour * time.Duration(duration)
			case "天":
				c.Duration = time.Hour * 24 * time.Duration(duration)
			case "周":
				c.Duration = time.Hour * 24 * 7 * time.Duration(duration)
			case "月":
				_, _ = c.Time(time.Now().AddDate(0, duration, 0))
			}

			return true, nil
		},
	}
}
