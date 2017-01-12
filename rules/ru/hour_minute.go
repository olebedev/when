package ru

import (
	"regexp"
	"strconv"
	"time"

	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

/*
	{"5:30pm", 0, "5:30pm", 0},
	{"5:30 pm", 0, "5:30 pm", 0},
	{"7-10pm", 0, "7-10pm", 0},
	{"5-30", 0, "5-30", 0},
	{"05:30pm", 0, "05:30pm", 0},
	{"05:30 pm", 0, "05:30 pm", 0},
	{"05:30", 0, "05:30", 0},
	{"05-30", 0, "05-30", 0},
	{"7-10 pm", 0, "7-10 pm", 0},
	{"11.1pm", 0, "11.1pm", 0},
	{"11.10 pm", 0, "11.10 pm", 0},

	https://play.golang.org/p/PmPBjHK4PA
*/

// 1. - int
// 2. - int
// 3. - ext?

func HourMinute(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|\\D|^)" +
			"((?:[0-1]{0,1}[0-9])|(?:2[0-3]))" +
			"(?:\\:|：|\\-|\\.)" +
			"((?:[0-5][0-9]))" +
			"(?:\\s*(утра|вечера|дня))?" +
			"(?:\\P{L}|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Hour != nil || c.Minute != nil) && s != rules.Override {
				return false, nil
			}

			hour, err := strconv.Atoi(m.Captures[0])
			if err != nil {
				return false, errors.Wrap(err, "hour minute rule")
			}

			minutes, err := strconv.Atoi(m.Captures[1])
			if err != nil {
				return false, errors.Wrap(err, "hour minute rule")
			}

			c.Minute = &minutes

			if m.Captures[2] != "" {
				if hour > 12 {
					return false, nil
				}
				switch m.Captures[2] {
				case "утра": // am
					c.Hour = &hour
				case "вечера", "дня": // pm
					if hour < 12 {
						hour += 12
					}
					c.Hour = &hour
				}
			} else {
				c.Hour = &hour
			}

			return true, nil
		},
	}
}
