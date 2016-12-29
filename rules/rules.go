package rules

import (
	"regexp"
	"time"
)

type Rule interface {
	Find(string) *Match
}

type Options struct {
	Afternoon, Everning, Morning, Noon int

	Distance int
	// WeekStartsOn time.Weekday
}

type Match struct {
	Left, Right int
	Text        string
	Captures    []string
	Order       float64
	Applier     func(*Match, *Context, *Options, time.Time) error
}

func (m Match) String() string { return m.Text }

func (m *Match) Apply(c *Context, o *Options, t time.Time) error {
	return m.Applier(m, c, o, t)
}

type F struct {
	RegExp  *regexp.Regexp
	Applier func(*Match, *Context, *Options, time.Time) error
}

func (f *F) Find(text string) *Match {
	m := &Match{
		Applier: f.Applier,
		Left:    -1,
	}

	indexes := f.RegExp.FindStringSubmatchIndex(text)

	length := len(indexes)
	if length <= 2 {

		return nil
	}

	for i := 2; i < length; i += 2 {
		if m.Left == -1 && indexes[i] >= 0 {
			m.Left = indexes[i]
		}
		// check if capture was found
		if indexes[i] >= 0 && indexes[i+1] >= 0 {
			m.Captures = append(m.Captures, text[indexes[i]:indexes[i+1]])
			m.Right = indexes[i+1]
		} else {
			m.Captures = append(m.Captures, "")
		}
	}

	if len(m.Captures) == 0 {
		return nil
	}

	m.Text = text[m.Left:m.Right]
	return m
}
