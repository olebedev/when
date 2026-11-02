package en

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
)

func Timezone(s rules.Strategy) rules.Rule {
	return &rules.F{
		RegExp: regexp.MustCompile("(?i)(?:\\W|^)" +
			"((?:UTC|GMT|EST|EDT|CST|CDT|MST|MDT|PST|PDT|[+-]\\d{2}:?\\d{2}))" +
			"(?:\\W|$)"),
		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			if c.Location != nil && s != rules.Override {
				return false, nil
			}

			tzStr := m.Captures[0]
			loc, err := parseTimezone(tzStr)
			if err != nil {
				return false, nil
			}

			c.Location = loc
			return true, nil
		},
	}
}

func parseTimezone(tz string) (*time.Location, error) {
	upper := strings.ToUpper(tz)
	switch upper {
	case "UTC", "GMT":
		return time.UTC, nil
	case "EST":
		return time.FixedZone("EST", -5*3600), nil
	case "EDT":
		return time.FixedZone("EDT", -4*3600), nil
	case "CST":
		return time.FixedZone("CST", -6*3600), nil
	case "CDT":
		return time.FixedZone("CDT", -5*3600), nil
	case "MST":
		return time.FixedZone("MST", -7*3600), nil
	case "MDT":
		return time.FixedZone("MDT", -6*3600), nil
	case "PST":
		return time.FixedZone("PST", -8*3600), nil
	case "PDT":
		return time.FixedZone("PDT", -7*3600), nil
	}

	loc, err := time.LoadLocation(tz)
	if err == nil {
		return loc, nil
	}

	return parseNumericOffset(tz)
}

func parseNumericOffset(tz string) (*time.Location, error) {
	re := regexp.MustCompile(`^([+-])(\d{2}):?(\d{2})?$`)
	matches := re.FindStringSubmatch(tz)
	if matches == nil {
		return nil, errors.New("invalid timezone offset")
	}

	sign := matches[1]
	hours, _ := time.ParseDuration(matches[2] + "h")
	var minutes time.Duration
	if matches[3] != "" {
		minutes, _ = time.ParseDuration(matches[3] + "m")
	}
	offset := int(hours + minutes)
	if sign == "-" {
		offset = -offset
	}

	return time.FixedZone(tz, offset), nil
}
