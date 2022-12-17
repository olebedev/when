package zh

import (
	"regexp"
	"strconv"
	"time"

	"github.com/olebedev/when/rules"
)

/*
	"上午 5点"
	"上午 5 点"
	"下午 3点"
	"下午 3 点"
	"下午 3点半"
	"下午 3点30"
	"下午 3:30"
	"下午 3：30"
    "下午 三点半"
*/

func HourMinute(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)" +
			"(?:(凌\\s*晨|早\\s*晨|早\\s*上|上\\s*午|下\\s*午|晚\\s*上|今晚)?\\s*)" +
			"((?:[0-1]{0,1}[0-9])|(?:2[0-3]))?" + "(?:\\s*)" +
			"(" + INTEGER_WORDS_PATTERN[3:] + "?" +
			"(\\:|：|\\-|点)" +
			"((?:[0-5][0-9]))?" +
			"(" + INTEGER_WORDS_PATTERN + "+)?" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Hour != nil || c.Minute != nil) && s != rules.Override {
				return false, nil
			}

			hour, exist := INTEGER_WORDS[m.Captures[2]] // 中文
			if !exist {
				hour, _ = strconv.Atoi(m.Captures[1])
			}

			if hour > 24 {
				return false, nil
			}

			minutes, exist := INTEGER_WORDS[m.Captures[5]]
			if !exist {
				minutes, _ = strconv.Atoi(m.Captures[4])
			}

			if minutes > 59 {
				return false, nil
			}
			c.Minute = &minutes

			lower := compressStr(m.Captures[0])
			switch lower {
			case "上午", "凌晨", "早晨", "早上":
				c.Hour = &hour
			case "下午", "晚上", "今晚":
				if hour < 12 {
					hour += 12
				}
				c.Hour = &hour
			case "":
				if hour > 23 {
					return false, nil
				}
				c.Hour = &hour

			}
			return true, nil
		},
	}
}
