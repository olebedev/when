package common

import "github.com/olebedev/when/rules"

var MONTHS_DAYS = []int{
	0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31,
}

func GetDays(year, month int) int {
	if month > 12 || month < 0 {
		return 0
	}

	if month != 2 {
		return MONTHS_DAYS[month]
	}

	if year%4 == 0 {
		if year%400 == 0 {
			return 29
		}
		if year%100 == 0 {
			return 28
		}
		return 29
	}

	return 28
}

var All = []rules.Rule{
	SlashDMY(rules.Override),
}
