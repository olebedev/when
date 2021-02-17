package ru

import (
	"regexp"
	"strconv"
	"time"

	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

var SECONDS_WORDS = map[string]int{
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

var SECONDS_WORDS_PATTERN = `(?:секунду|` + common_WORDS_PATTERN + `)`

const (
	vdPosHour int = iota + 1
	_
	vdPosMinute
	_
	vdPosSecond
)

func VoiceDeadline(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|\D|^)(через\s+)` +
			`(?:(` + INTEGER_WORDS_PATTERN + `|\d{1,2})(\s*час(?:а|ов|ам)?))?` +
			`\s*(?:(` + MINUTES_WORDS_PATTERN + `|\d{1,3})(\s*минут(?:у|а|ы)?))?` +
			`\s*(?:(` + SECONDS_WORDS_PATTERN + `|\d{1,3})(\s*секунд(?:у|а|ы)?))?` +
			`(?:\P{L}|$)`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Duration != 0 && s != rules.Override {
				return false, nil
			}

			var (
				hour, minute, second int
				err                  error
			)

			hourCaptures := m.Captures[vdPosHour]
			minuteCaptures := m.Captures[vdPosMinute]
			secondCaptures := m.Captures[vdPosSecond]

			if hourCaptures == "" && minuteCaptures == "" && secondCaptures == "" {
				return false, nil
			}

			if hourCaptures != "" {
				if n, ok := INTEGER_WORDS[hourCaptures]; ok {
					hour = n
				} else {
					hour, err = strconv.Atoi(hourCaptures)
					if err != nil {
						return false, errors.Wrap(err, "voice deadline rule")
					}
				}
			}

			if minuteCaptures != "" {
				if n, ok := MINUTES_WORDS[minuteCaptures]; ok {
					minute = n
				} else {
					minute, err = strconv.Atoi(minuteCaptures)
					if err != nil {
						return false, errors.Wrap(err, "voice deadline rule")
					}
				}
			}

			if secondCaptures != "" {
				if n, ok := SECONDS_WORDS[secondCaptures]; ok {
					second = n
				} else {
					second, err = strconv.Atoi(secondCaptures)
					if err != nil {
						return false, errors.Wrap(err, "voice deadline rule")
					}
				}
			}

			c.Duration = time.Duration(hour)*time.Hour +
				time.Duration(minute)*time.Minute +
				time.Duration(second)*time.Second

			return true, nil
		},
	}
}
