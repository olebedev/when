package zh

import (
	"regexp"
	"strconv"
	"time"

	"github.com/olebedev/when/rules"
)

/*
	规则名称：精确到月份的日期
*/

func ExactMonthDate(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("" +
			"(?:\\W|^)" +
			"(1[0-2]|[1-9]|" + MON_WORDS_PATTERN + ")" + "(?:\\s*)" +
			"(月|-|/|\\.|)" + "(?:\\s*)" +
			"(1[0-9]|2[0-9]|3[0-1]|[1-9]|" + DAY_WORDS_PATTERN + ")" + "(?:\\s*)" +
			"(日|号)?",
		),

		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			_ = overwrite
			if m.Captures[1] == "" {
				return false, nil
			}
			monInt, exist := MON_WORDS[compressStr(m.Captures[0])]
			if !exist {
				mon, err := strconv.Atoi(m.Captures[0])
				if err != nil {
					return false, nil
				}
				monInt = mon
			}

			dayInt, exist := DAY_WORDS[compressStr(m.Captures[2])]
			if !exist {
				day, err := strconv.Atoi(m.Captures[2])
				if err != nil {
					return false, nil
				}
				dayInt = day
			}

			c.Month = &monInt

			c.Day = &dayInt

			return true, nil
		},
	}
}
