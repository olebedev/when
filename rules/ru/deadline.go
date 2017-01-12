package ru

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

// https://play.golang.org/p/A-cF_q9U34

func Deadline(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\P{L}|^)" +
			"(в\\sтечении|за|через)\\s*" +
			"(" + INTEGER_WORDS_PATTERN + "|[0-9]+|полу?|несколько|нескольких)?\\s*" +
			"(секунд(?:у|ы)?|минут(?:у|ы)?|час(?:а|ов)?|день|дня|дней|недел(?:я|ь|и|ю)|месяц(?:а|ев)?|год(?:а)?|лет)\\s*" +
			"(?:\\P{L}|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Duration != 0 && s != rules.Override {
				return false, nil
			}

			numStr := strings.TrimSpace(m.Captures[1])

			var num int
			var err error

			if n, ok := INTEGER_WORDS[numStr]; ok {
				num = n
			} else if numStr == "" {
				num = 1
			} else if strings.Contains(numStr, "неск") {
				num = 3
			} else if strings.Contains(numStr, "пол") {
				// pass
			} else {
				num, err = strconv.Atoi(numStr)
				if err != nil {
					return false, errors.Wrapf(err, "convert '%s' to int", numStr)
				}
			}

			exponent := strings.TrimSpace(m.Captures[2])

			if !strings.Contains(numStr, "пол") {
				switch {
				case strings.Contains(exponent, "секунд"):
					c.Duration = time.Duration(num) * time.Second
				case strings.Contains(exponent, "мин"):
					c.Duration = time.Duration(num) * time.Minute
				case strings.Contains(exponent, "час"):
					c.Duration = time.Duration(num) * time.Hour
				case strings.Contains(exponent, "дн") || strings.Contains(exponent, "день"):
					c.Duration = time.Duration(num) * 24 * time.Hour
				case strings.Contains(exponent, "недел"):
					c.Duration = time.Duration(num) * 7 * 24 * time.Hour
				case strings.Contains(exponent, "месяц"):
					c.Month = pointer.ToInt((int(ref.Month()) + num) % 12)
				case strings.Contains(exponent, "год") || strings.Contains(exponent, "лет"):
					c.Year = pointer.ToInt(ref.Year() + num)
				}
			} else {
				switch {
				case strings.Contains(exponent, "час"):
					c.Duration = 30 * time.Minute
				case strings.Contains(exponent, "дн") || strings.Contains(exponent, "день"):
					c.Duration = 12 * time.Hour
				case strings.Contains(exponent, "недел"):
					c.Duration = 7 * 12 * time.Hour
				case strings.Contains(exponent, "месяц"):
					// 2 weeks
					c.Duration = 14 * 24 * time.Hour
				case strings.Contains(exponent, "год") || strings.Contains(exponent, "лет"):
					c.Month = pointer.ToInt((int(ref.Month()) + 6) % 12)
				}
			}

			return true, nil
		},
	}
}
