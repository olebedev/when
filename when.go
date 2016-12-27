package when

import (
	"sort"
	"time"

	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

type Options struct {
	Distance int
}

type Parser struct {
	options    *Options
	rules      []rules.Rule
	middleware []func(string) (string, error)
}

func (p *Parser) When(text string, ts ...time.Time) (time.Time, int, string, error) {
	start := -1
	seq := ""
	t := time.Time{}
	for _, _t := range ts {
		t = _t
		break
	}

	if t.IsZero() {
		t = time.Now()
	}

	var err error
	// apply middlewares
	for i, b := range p.middleware {
		text, err = b(text)
		if err != nil {
			return t, start, seq, errors.Wrapf(err, "apply middleware func #%d", i)
		}
	}

	// find all matches
	matches := make([]*rules.Match, 0)
	for i, rule := range p.rules {
		r, err := rule.Find(text)
		if err != nil {
			return t, start, seq, errors.Wrapf(err, "find match #%d", i)
		}
		if r != nil {
			matches = append(matches, r...)
		}
	}

	// not found
	if len(matches) == 0 {
		return t, start, seq, nil
	}

	// find a cluster
	sort.Sort(rules.MatchByIndex(matches))

	end := matches[0].Index
	start = matches[0].Index

	for i, m := range matches {
		if m.Index <= end+1 {
			end = m.Index + len(m.String())
		} else {
			matches = matches[:i]
			break
		}
	}

	seq = text[start:end]

	// apply rules
	sort.Sort(rules.MatchByOrderAndIndex(matches))

	ctx := &rules.Context{Text: seq}
	for i, applier := range matches {
		err = applier.Apply(ctx)
		if err != nil {
			return t, start, seq, errors.Wrapf(err, "apply modifications to the context #%d", i)
		}
	}

	t, err = ctx.Time(t)
	if err != nil {
		return t, start, seq, errors.Wrap(err, "bind context")
	}

	return t, start, seq, nil
}

func (p *Parser) Add(r rules.Rule) {
	p.rules = append(p.rules, r)
}

func (p *Parser) Use(f func(string) (string, error)) {
	p.middleware = append(p.middleware, f)
}

func New(o *Options) *Parser {
	return &Parser{options: o}
}

/*
	ask := when.New(nil)

	// to be able to get rid of some corner cases
	ask.Use(func(text string) (string, error) { return text, nil })

	ask.Add(Rule)
	ask.Add(Rule)
	ask.Add(Rule)

	t, index, text, err := ask.When("in 3 hours")

*/
