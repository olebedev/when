package ru

import "github.com/olebedev/when/rules"

var All = []rules.Rule{
	// Weekday(rules.OverWrite),
	CasualDate(rules.OverWrite),
	CasualTime(rules.OverWrite),
	Deadline(rules.OverWrite),
	Hour(rules.OverWrite),
	// HourMinute(rules.OverWrite),
}

var WEEKDAY_OFFSET = map[string]int{
	"воскресенье": 0,
	"воск":        0,
	"понедельник": 1,
	"пн":      1,
	"вторник": 2,
	"вт":      2,
	"среда":   3,
	"среду":   3,
	"ср":      3,
	"четверг": 4,
	"чт":      4,
	"пятница": 5,
	"пятницу": 5,
	"пт":      5,
	"суббота": 6,
	"субботу": 6,
	"сб":      6,
}

var WEEKDAY_OFFSET_PATTERN = "(?:воскресенье|воскр|вс|понедельник|пн|вторник|вт|среда|среду|ср|четверг|чт|пятница|пятницу|пт|суббота|субботу|сб)"

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
