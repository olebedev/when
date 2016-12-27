package en

import (
	"strings"
	"time"

	"github.com/AlekSi/pointer"
	"github.com/dlclark/regexp2"
	"github.com/olebedev/when/rules"
)

type CasualDate struct {
	*regexp2.Regexp
}

func (_ *CasualDate) Apply(m *rules.Match, c *rules.Context) error {
	lower := strings.ToLower(strings.TrimSpace(m.String()))

	switch {
	case strings.Contains(lower, "tonight"):
		c.Hour = pointer.ToInt(23)
	case strings.Contains(lower, "today"):
		c.Hour = pointer.ToInt(18)
	case strings.Contains(lower, "tomorrow"), strings.Contains(lower, "tmr"):
		c.Duration += time.Hour * 24
	case strings.Contains(lower, "yesterday"):
		c.Duration -= time.Hour * 24
	case strings.Contains(lower, "last night"):
		c.Hour = pointer.ToInt(23)
		c.Duration -= time.Hour * 24
	}

	return nil
}

func (c *CasualDate) Find(text string) ([]*rules.Match, error) {
	return rules.AllMatch(c.Regexp, text, 1, c)
}

func NewCasualDate() rules.Rule {
	return &CasualDate{
		regexp2.MustCompile(
			`(\W|^)(now|today|tonight|last\s*night|(?:tomorrow|tmr|yesterday)\s*|tomorrow|tmr|yesterday)(?=\W|$)`,
			regexp2.IgnoreCase|regexp2.ECMAScript),
	}
}
