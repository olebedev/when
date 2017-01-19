package ru

import (
	"regexp"
	"strconv"
	"time"

	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

/*
	"5pm"
	"5 pm"
	"5am"
	"5pm"
	"5A."
	"5P."
	"11 P.M."
	https://play.golang.org/p/w2PeQ3l_rp
*/

func Hour(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"(" + INTEGER_WORDS_PATTERN + "|\\d{1,2})" +
			"(?:\\s*час(?:а|ов|ам)?)?(?:\\s*(утра|вечера|дня))" +
			"(?:\\P{L}|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Hour != nil && s != rules.Override {
				return false, nil
			}

			var hour int
			var err error

			if n, ok := INTEGER_WORDS[m.Captures[0]]; ok {
				hour = n
			} else {
				hour, err = strconv.Atoi(m.Captures[0])
				if err != nil {
					return false, errors.Wrap(err, "hour rule")
				}
			}

			if hour > 12 {
				return false, nil
			}
			zero := 0

			switch m.Captures[1] {
			case "утра":
				c.Hour = &hour
			case "вечера", "дня":
				if hour < 12 {
					hour += 12
				}
				c.Hour = &hour
			}
			c.Minute = &zero

			return true, nil
		},
	}
}
