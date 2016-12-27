package rules

import (
	"time"

	"github.com/dlclark/regexp2"
	"github.com/pkg/errors"
)

type Rule interface {
	Find(string) ([]*Match, error)
}

type Applier interface {
	Apply(*Match, *Context) error
}

type Context struct {
	Text string

	// accumulator of relative values
	Duration time.Duration

	// Aboslute values
	Year, Month, Weekday, Day, Hour, Minute, Second *int
	Location                                        *time.Location
}

func (c *Context) Time(t time.Time) (time.Time, error) {
	if t.IsZero() {
		t = time.Now()
	}

	if c.Duration != 0 {
		t = t.Add(c.Duration)
	}

	if c.Year != nil {
		t = time.Date(*c.Year, t.Month(), t.Day(), t.Hour(),
			t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	if c.Month != nil {
		t = time.Date(t.Year(), time.Month(*c.Month), t.Day(),
			t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	if c.Weekday != nil {
		diff := int(time.Weekday(*c.Weekday) - t.Weekday())
		t = time.Date(t.Year(), t.Month(), t.Day()+diff, t.Hour(),
			t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	if c.Day != nil {
		t = time.Date(t.Year(), t.Month(), *c.Day, t.Hour(),
			t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	if c.Hour != nil {
		t = time.Date(t.Year(), t.Month(), t.Day(), *c.Hour,
			t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	if c.Minute != nil {
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(),
			*c.Minute, t.Second(), t.Nanosecond(), t.Location())
	}

	if c.Second != nil {
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(),
			t.Minute(), *c.Second, t.Nanosecond(), t.Location())
	}

	if c.Location != nil {
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(),
			t.Minute(), t.Second(), t.Nanosecond(), c.Location)
	}

	return t, nil
}

type Match struct {
	*regexp2.Match
	Order   float64
	Applier Applier
}

func (m *Match) Apply(c *Context) error {
	return m.Applier.Apply(m, c)
}

type MatchByIndex []*Match

func (m MatchByIndex) Len() int {
	return len(m)
}

func (m MatchByIndex) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MatchByIndex) Less(i, j int) bool {
	return m[i].Index < m[j].Index
}

type MatchByOrderAndIndex []*Match

func (m MatchByOrderAndIndex) Len() int {
	return len(m)
}

func (m MatchByOrderAndIndex) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MatchByOrderAndIndex) Less(i, j int) bool {
	if m[i].Order < m[j].Order {
		return true
	} else if m[i].Order > m[j].Order {
		return false
	}
	return m[i].Index < m[j].Index
}

func AllMatch(re *regexp2.Regexp, text string, order float64, applier Applier) ([]*Match, error) {
	m, err := re.FindStringMatch(text)
	if err != nil {
		return nil, errors.Wrap(err, "find string match")
	}

	matches := []*Match{&Match{
		Match:   m,
		Order:   order,
		Applier: applier,
	}}

	for {
		m, err = re.FindNextMatch(m)
		if err != nil {
			return nil, errors.Wrap(err, "find next match")
		}
		if m == nil {
			break
		} else {
			matches = append(matches, &Match{
				Match:   m,
				Order:   order,
				Applier: applier,
			})
		}
	}
	return matches, nil
}
