package ru

import (
	"regexp"
	"strconv"
	"time"

	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

var MINUTES_WORDS = map[string]int{
	"минуту":       1,
	"один":         1,
	"одну":         1,
	"одного":       1,
	"два":          2,
	"две":          2,
	"три":          3,
	"четыре":       4,
	"пять":         5,
	"шесть":        6,
	"семь":         7,
	"восемь":       8,
	"девять":       9,
	"десять":       10,
	"одиннадцать":  11,
	"двенадцать":   12,
	"тринадцать":   13,
	"четырнадцать": 14,
	"пятнадцать":   15,
	"шестнадцать":  16,
	"семнадцать":   17,
	"восемнадцать": 18,
	"девятнадцать": 19,
	"двадцать":     20,
	"тридцать":     30,
	"сорок":        40,
	"пятьдесят":    50,
	"шестьдесят":   60,
	"семдесят":     70,
	"восемдесят":   80,
	"девяносто":    90,
	"сто":          100,
}

var common_WORDS_PATTERN = `один|одну|одного|два|две|три|четыре|пять|шесть|семь|восемь|девять|десять|` +
	`одиннадцать|двенадцать|тринадцать|четырнадцать|пятнадцать|шестнадцать|семнадцать|восемнадцать|девятнадцать|` +
	`двадцать|тридцать|сорок|пятьдесят|шестьдесят|семдесят|восемдесят|девяносто|` +
	`сто`

var MINUTES_WORDS_PATTERN = `(?:минуту|` + common_WORDS_PATTERN + `)`

const (
	vhmPosHour int = iota
	vhmPosDaytime
	vhmPosMinute
	vhmPosMinuteWord
	vhmPosWrongWords
	vhmPosDaytime2
)

func VoiceHourMinute(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|\D|^)в\s+` +
			`(` + INTEGER_WORDS_PATTERN + `|\d{1,2})` + // 0
			`[^:]` +
			`(?:\s*час(?:а|ов|ам)?)?(?:\s*(утра|вечера|дня|ночи))?` + // 1
			`\s*(` + MINUTES_WORDS_PATTERN + `|\d{1,2})?` + // 2
			`(?:\s*(минут(?:у|а|ы)?))?` + // 3
			`(?:\s*(\d|секунд|процент|пункт|раз))?` + // 4
			`(?:\s*(утра|вечера|дня|ночи))?`), // 5
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Hour != nil || c.Minute != nil) && s != rules.Override {
				return false, nil
			}

			var (
				hour, minute int
				err          error
				minutesEmpty bool
			)

			hourCaptures := m.Captures[vhmPosHour]
			if n, ok := INTEGER_WORDS[hourCaptures]; ok {
				hour = n
			} else {
				hour, err = strconv.Atoi(hourCaptures)
				if err != nil {
					return false, errors.Wrap(err, "voice hour minute rule")
				}
			}

			minuteCaptures := m.Captures[vhmPosMinute]
			if minuteCaptures == "" {
				minuteCaptures = "0"
				minutesEmpty = true
			}
			if n, ok := MINUTES_WORDS[minuteCaptures]; ok {
				minute = n
			} else {
				minute, err = strconv.Atoi(minuteCaptures)
				if err != nil {
					return false, errors.Wrap(err, "voice hour minute rule")
				}
			}

			if len(m.Captures[vhmPosWrongWords]) > 0 { // do not match (секунд|процент|пункт|раз)
				if minutesEmpty {
					return false, nil
				}
				minute = 0
			}

			if minutesEmpty && m.Captures[vhmPosMinuteWord] != "" {
				return false, nil
			}

			if hour > 12 {
				c.Hour = &hour
				c.Minute = &minute
				return true, nil
			}

			dayTime := m.Captures[vhmPosDaytime2]
			if dayTime == "" {
				dayTime = m.Captures[vhmPosDaytime]
			}

			switch dayTime {
			case "утра":
				c.Hour = &hour
			case "ночи":
				if hour == 12 {
					hour = 0
				}
				c.Hour = &hour
			case "вечера", "вечером", "дня":
				if hour < 12 {
					hour += 12
				}
				c.Hour = &hour
			default:
				c.Hour = &hour
			}
			c.Minute = &minute

			return true, nil
		},
	}
}
