package when

import (
	"fmt"
	"sort"
	"time"

	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
	"github.com/pkg/errors"
)

type Parser struct {
	options    *rules.Options
	rules      []rules.Rule
	middleware []func(string) (string, error)
}

func (p *Parser) Parse(text string, ts ...time.Time) (time.Time, int, string, error) {
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

	if p.options == nil {
		p.options = defaultOptions
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
	for _, rule := range p.rules {
		r := rule.Find(text, p.options)
		if r != nil {
			matches = append(matches, r)
		}
	}

	// not found
	if len(matches) == 0 {
		return t, start, seq, nil
	}

	// find a cluster
	sort.Sort(rules.MatchByIndex(matches))

	end := matches[0].Right
	start = matches[0].Left

	for i, m := range matches {
		if m.Left <= end+1 {
			end = m.Right
		} else {
			matches = matches[:i]
			break
		}
	}

	seq = text[start:end]
	fmt.Println("seq", start, end)

	// apply rules
	sort.Sort(rules.MatchByOrderAndIndex(matches))

	ctx := &rules.Context{Text: seq}
	for i, applier := range matches {
		err = applier.Apply(ctx, p.options, t)
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

func (p *Parser) Add(r ...rules.Rule) {
	p.rules = append(p.rules, r...)
}

func (p *Parser) Use(f ...func(string) (string, error)) {
	p.middleware = append(p.middleware, f...)
}

func (p *Parser) SetOptions(o *rules.Options) {
	p.options = o
}

func New(o *rules.Options) *Parser {
	if o == nil {
		return &Parser{options: defaultOptions}
	}
	return &Parser{options: o}
}

var EN *Parser

var defaultOptions = &rules.Options{
	Morning: 8,
}

// var RU *Parser

func init() {
	EN = New(defaultOptions)
	EN.Add(en.All...)
}
