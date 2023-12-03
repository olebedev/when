package common

import (
	"regexp"
	"strconv"
	"time"

	"github.com/olebedev/when/rules"
)

/*

- DD/MM/YYYY
- 11/3/2015
- 11/3/2015
- 11/3

also with "\", gift for windows' users

https://play.golang.org/p/29LkTfe1Xr
*/

var MONTHS_DAYS = []int{
	0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
}

func getDays(year, month int) int {
	// naive leap year check
	if (year-2000)%4 == 0 && month == 2 {
		return 29
	}
	return MONTHS_DAYS[month]
}

func SlashDMY(s rules.Strategy) rules.Rule {

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(0{0,1}[1-9]|1[0-9]|2[0-9]|3[01])" +
			"[\\/\\\\]" +
			"(0{0,1}[1-9]|1[0-2])" +
			"(?:[\\/\\\\]" +
			"((?:1|2)[0-9]{3})\\s*)?" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Day != nil || c.Month != nil || c.Year != nil) && s != rules.Override {
				return false, nil
			}

			day, _ := strconv.Atoi(m.Captures[0])
			month, _ := strconv.Atoi(m.Captures[1])
			year := -1
			if m.Captures[2] != "" {
				year, _ = strconv.Atoi(m.Captures[2])
			}

			if day == 0 {
				return false, nil
			}

		WithYear:
			if year != -1 {
				if getDays(year, month) >= day {
					c.Year = &year
					c.Month = &month
					c.Day = &day
				} else {
					return false, nil
				}
				return true, nil
			}

			if month < int(ref.Month()) {
				year = ref.Year() + 1
			} else if month == int(ref.Month()) {
				if day > getDays(ref.Year(), month) {
					// invalid date: day is after last day of the month
					return false, nil
				}

				if day >= ref.Day() {
					year = ref.Year()
				} else {
					year = ref.Year() + 1
				}
			} else {
				year = ref.Year()
			}

			goto WithYear
		},
	}
}
