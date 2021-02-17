package ru

import (
	"regexp"
	"strconv"
	"time"

	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

func VoiceHourNaked(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile(`(?i)(?:\W|\D|^)в\s+` +
			`(` + INTEGER_WORDS_PATTERN + `|\d{1,2})` +
			`(?:[^:]|$)` +
			`\s*(\d|минут|секунд|час|процент|пункт|раз)?`),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if (c.Hour != nil || c.Minute != nil) && s != rules.Override {
				return false, nil
			}

			var hour int
			var err error

			if n, ok := INTEGER_WORDS[m.Captures[0]]; ok {
				hour = n
			} else {
				hour, err = strconv.Atoi(m.Captures[0])
				if err != nil {
					return false, errors.Wrap(err, "voice hour naked rule")
				}
			}

			if len(m.Captures[1]) > 0 { // do not match (\d|минут|секунд|час|процент|пункт)
				return false, nil
			}

			zero := 0
			c.Hour = &hour
			c.Minute = &zero
			return true, nil
		},
	}
}
