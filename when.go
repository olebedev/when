package when

import (
	"sort"
	"time"

	"github.com/olebedev/when/rules"
	"github.com/pkg/errors"
)

// Parser is a struct which contains options
// rules, and middlewares to call
type Parser struct {
	options    *rules.Options
	rules      []rules.Rule
	middleware []func(string) (string, error)
}

// Result is a struct which contains parsing meta-info
type Result struct {
	// Index is a start index
	Index int
	// Text is a text found and processed
	Text string
	// Source is input string
	Source string
	// Time is an output time
	Time time.Time
}

// Parse returns Result and error if any. If have not matches it returns nil, nil.
func (p *Parser) Parse(text string, base time.Time) (*Result, error) {
	res := Result{
		Source: text,
		Time:   base,
		Index:  -1,
	}

	if p.options == nil {
		p.options = defaultOptions
	}

	var err error
	// apply middlewares
	for _, b := range p.middleware {
		text, err = b(text)
		if err != nil {
			return nil, err
		}
	}

	// find all matches
	matches := make([]*rules.Match, 0)
	c := float64(0)
	for _, rule := range p.rules {
		r := rule.Find(text)
		if r != nil {
			r.Order = c
			c++
			matches = append(matches, r)
		}
	}

	// not found
	if len(matches) == 0 {
		return nil, nil
	}

	// find a cluster
	sort.Sort(rules.MatchByIndex(matches))

	// get borders of the matches
	end := matches[0].Right
	res.Index = matches[0].Left

	for i, m := range matches {
		if m.Left <= end+p.options.Distance {
			end = m.Right
		} else {
			matches = matches[:i]
			break
		}
	}

	res.Text = text[res.Index:end]

	// apply rules
	sort.Sort(rules.MatchByOrder(matches))

	ctx := &rules.Context{Text: res.Text}
	for _, applier := range matches {
		err = applier.Apply(ctx, p.options, res.Time)
		if err != nil {
			return nil, err
		}
	}

	res.Time, err = ctx.Time(res.Time)
	if err != nil {
		return nil, errors.Wrap(err, "bind context")
	}

	return &res, nil
}

// Add adds  given rules to the main chain.
func (p *Parser) Add(r ...rules.Rule) {
	p.rules = append(p.rules, r...)
}

// Use adds give functions to middlewares.
func (p *Parser) Use(f ...func(string) (string, error)) {
	p.middleware = append(p.middleware, f...)
}

// SetOptions sets options object to use.
func (p *Parser) SetOptions(o *rules.Options) {
	p.options = o
}

// New returns Parser initialised with given options.
func New(o *rules.Options) *Parser {
	if o == nil {
		return &Parser{options: defaultOptions}
	}
	return &Parser{options: o}
}

// default options for internal usage
var defaultOptions = &rules.Options{
	Distance: 1,
}
