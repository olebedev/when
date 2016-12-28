package rules

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Rule interface {
	Find(string, *Options) *Match
}

type Options struct {
	Strict bool

	Afternoon, Everning, Morning, Noon int
}

type Match struct {
	Left, Right int
	Groups      []string
	Order       float64
	Applier     func(*Match, *Context, *Options, time.Time) error
}

func (m Match) String() string {
	if len(m.Groups) == 0 {
		return ""
	}
	return strings.Join(m.Groups[1:], "")
}

func (m *Match) Apply(c *Context, o *Options, t time.Time) error {
	return m.Applier(m, c, o, t)
}

type F struct {
	RegExp       *regexp.Regexp
	RegExpStrict *regexp.Regexp
	Applier      func(*Match, *Context, *Options, time.Time) error
	Order        float64
}

func (f *F) Find(text string, o *Options) *Match {
	m := &Match{
		Order:   f.Order,
		Applier: f.Applier,
	}

	if o.Strict && f.RegExpStrict != nil {
		m.Groups = f.RegExpStrict.FindStringSubmatch(text)
		if len(m.Groups) != 0 {
			i := f.RegExpStrict.FindStringSubmatchIndex(text)
			if len(i) >= 3 {
				m.Left = i[3]
				m.Right = i[3] + len(m.String())
			}
		} else {
			return nil
		}
	} else {
		m.Groups = f.RegExp.FindStringSubmatch(text)
		if len(m.Groups) != 0 {
			fmt.Println(1)
			i := f.RegExp.FindStringSubmatchIndex(text)
			fmt.Println(2, i)
			if len(i) >= 3 {
				m.Left = i[2]
				m.Right = i[len(i)-1]
			}
			fmt.Println(3, m.Left, m.Right)
		} else {
			return nil
		}
	}

	return m
}
