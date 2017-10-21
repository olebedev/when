package en

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/olebedev/when/rules"
)

// <[]string{"third of march", "third", "", "march", "", ""}>
// <[]string{"march third", "", "", "march", "third", ""}>
// <[]string{"march 3rd", "", "", "march", "3rd", ""}>
// <[]string{"3rd march", "3rd", "", "march", "", ""}>
// <[]string{"march 3", "", "", "march", "", "3"}>
// <[]string{"1st of september", "1st", "", "september", "", ""}>
// <[]string{"sept. 1st", "", "", "sept.", "1st", ""}>
// <[]string{"march 7th", "", "", "march", "7th", ""}>
// <[]string{"october 21st", "", "", "october", "21st", ""}>
// <[]string{"twentieth of december", "twentieth", "", "december", "", ""}>
// <[]string{"march 10th", "", "", "march", "10th", ""}>
// <[]string{"jan. 6", "", "", "jan.", "", "6"}>
// <[]string{"february", "", "", "february", "", ""}>
// <[]string{"october", "", "", "october", "", ""}>
// <[]string{"jul.", "", "", "jul.", "", ""}>
// <[]string{"june", "", "", "june", "", ""}>

// https://play.golang.org/p/Zfjl6ERNkq

// 1. - ordinal day?
// 2. - numeric day?
// 3. - month
// 4. - ordinal day?
// 5. - ordinal day?

func ExactMonthDate(s rules.Strategy) rules.Rule {
	overwrite := s == rules.Override

	return &rules.F{
		RegExp: regexp.MustCompile("(?i)" +
			"(?:\\W|^)" +
			"(?:(?:(" + ORDINAL_WORDS_PATTERN[3:] + "(?:\\s+of)?|([0-9]+))\\s*)?" +
			"(" + MONTH_OFFSET_PATTERN[3:] + // skip '(?:'
			"(?:\\s*(?:(" + ORDINAL_WORDS_PATTERN[3:] + "|([0-9]+)))?" +
			"(?:\\W|$)",
		),

		Applier: func(m *rules.Match, c *rules.Context, o *rules.Options, ref time.Time) (bool, error) {
			_ = overwrite

			ord1 := strings.ToLower(strings.TrimSpace(m.Captures[0]))
			num1 := strings.ToLower(strings.TrimSpace(m.Captures[1]))
			mon := strings.ToLower(strings.TrimSpace(m.Captures[2]))
			ord2 := strings.ToLower(strings.TrimSpace(m.Captures[3]))
			num2 := strings.ToLower(strings.TrimSpace(m.Captures[4]))

			monInt, ok := MONTH_OFFSET[mon]
			if !ok {
				return false, nil
			}

			c.Month = &monInt

			if ord1 != "" {
				ordInt, ok := ORDINAL_WORDS[ord1]
				if !ok {
					return false, nil
				}

				c.Day = &ordInt
			}

			if num1 != "" {
				n, err := strconv.ParseInt(num1, 10, 8)
				if err != nil {
					return false, nil
				}

				num := int(n)

				c.Day = &num
			}

			if ord2 != "" {
				ordInt, ok := ORDINAL_WORDS[ord2]
				if !ok {
					return false, nil
				}

				c.Day = &ordInt
			}

			if num2 != "" {
				n, err := strconv.ParseInt(num2, 10, 8)
				if err != nil {
					return false, nil
				}

				num := int(n)

				c.Day = &num
			}

			return true, nil
		},
	}
}
