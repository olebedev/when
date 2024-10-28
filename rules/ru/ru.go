package ru

import (
	"github.com/olebedev/when/rules"
	"time"
)

var All = []rules.Rule{
	Weekday(rules.Override),
	CasualDate(rules.Override),
	CasualTime(rules.Override),
	Hour(rules.Override),
	HourMinute(rules.Override),
	Deadline(rules.Override),
	Date(rules.Override),
	DotDateTime(rules.Override),
}

var WEEKDAY_OFFSET = map[string]int{
	"воскресенье":  0,
	"воскресенья":  0,
	"воск":         0,
	"понедельник":  1,
	"понедельнику": 1,
	"понедельника": 1,
	"пн":           1,
	"вторник":      2,
	"вторника":     2,
	"вторнику":     2,
	"вт":           2,
	"среда":        3,
	"среду":        3,
	"среде":        3,
	"ср":           3,
	"четверг":      4,
	"четверга":     4,
	"четвергу":     4,
	"чт":           4,
	"пятница":      5,
	"пятнице":      5,
	"пятницы":      5,
	"пятницу":      5,
	"пт":           5,
	"суббота":      6,
	"субботы":      6,
	"субботе":      6,
	"субботу":      6,
	"сб":           6,
}

var WEEKDAY_OFFSET_PATTERN = "(?:воскресенье|воскресенья|воск|понедельник|понедельнику|понедельника|пн|вторник|вторника|вторнику|вт|среда|среду|среде|ср|четверг|четверга|четвергу|чт|пятница|пятнице|пятницы|пятницу|пт|суббота|субботы|субботе|субботу|сб)"

var INTEGER_WORDS = map[string]int{
	"час":         1,
	"один":        1,
	"одну":        1,
	"одного":      1,
	"два":         2,
	"две":         2,
	"три":         3,
	"четыре":      4,
	"пять":        5,
	"шесть":       6,
	"семь":        7,
	"восемь":      8,
	"девять":      9,
	"десять":      10,
	"одиннадцать": 11,
	"двенадцать":  12,
}

var INTEGER_WORDS_PATTERN = `(?:час|один|одну|одного|два|две|три|четыре|пять|шесть|семь|восемь|девять|десять|одиннадцать|двенадцать)`

var MONTHS = map[string]time.Month{
	"января":   time.January,
	"февраля":  time.February,
	"марта":    time.March,
	"апреля":   time.April,
	"мая":      time.May,
	"июня":     time.June,
	"июля":     time.July,
	"августа":  time.August,
	"сентября": time.September,
	"октября":  time.October,
	"ноября":   time.November,
	"декабря":  time.December,
}

var MONTHS_PATTERN = `(?:января|февраля|марта|апреля|мая|июня|июля|августа|сентября|октября|ноября|декабря)`
