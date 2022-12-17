package zh

import (
	"regexp"
	"time"

	"github.com/olebedev/when/rules"
)

/*
	子时 23:00 - 01：00
	丑时 01:00 - 03：00
	寅时 03:00 - 05：00
	卯时 05:00 - 07：00
	辰时 07:00 - 09：00
	巳时 09:00 - 11：00
	午时 11:00 - 13：00
	未时 13:00 - 15：00
	申时 15:00 - 17：00
	酉时 17:00 - 19：00
	戌时 19:00 - 21：00
	亥时 21:00 - 23：00
*/

func TraditionHour(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("" +
			"(?:(子\\s?时|丑\\s?时|寅\\s?时|卯\\s?时|辰\\s?时|巳\\s?时|午\\s?时|未\\s?时|申\\s?时|酉\\s?时|戌\\s?时|亥\\s?时))\\s?" +
			"(?:(一\\s?刻|二\\s?刻|两\\s?刻|三\\s?刻|四\\s?刻|五\\s?刻|六\\s?刻|七\\s?刻|1\\s?刻|2\\s?刻|3\\s?刻|4\\s?刻|5\\s?刻|6\\s?刻|7\\s?刻))?",
		),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Hour != nil && s != rules.Override {
				return false, nil
			}
			hour, exist := TRADITION_HOUR_WORDS[compressStr(m.Captures[0])]
			if !exist {
				return false, nil
			}
			c.Hour = &hour
			zero := 0
			c.Minute = &zero
			if minute, exist := TRADITION_MINUTE_WORDS[compressStr(m.Captures[1])]; exist {
				if minute > 60 {
					hour := *c.Hour + 1
					c.Hour = &hour
					minute = minute - 60
					c.Minute = &minute
				} else {
					c.Minute = &minute
				}
			}
			return true, nil
		},
	}
}

func compressStr(str string) string {
	if str == "" {
		return ""
	}
	reg := regexp.MustCompile(`\s+`)
	return reg.ReplaceAllString(str, "")
}
