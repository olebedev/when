package br_test

import (
	"testing"
	"time"

	"github.com/omniboost/when"
	"github.com/omniboost/when/rules"
	"github.com/omniboost/when/rules/br"
)

func TestExactMonthDate(t *testing.T) {
	w := when.New(nil)
	w.Add(br.ExactMonthDate(rules.Override))

	fixtok := []Fixture{
		{"3 de março", 0, "3 de março", 1368 * time.Hour},
		{"1 de setembro", 0, "1 de setembro", 5736 * time.Hour},
		{"1 set", 0, "1 set", 5736 * time.Hour},
		{"1 set.", 0, "1 set.", 5736 * time.Hour},
		{"1º de setembro", 0, "1º de setembro", 5736 * time.Hour},
		{"1º set.", 0, "1º set.", 5736 * time.Hour},
		{"7 de março", 0, "7 de março", 1464 * time.Hour},
		{"21 de outubro", 0, "21 de outubro", 6936 * time.Hour},
		{"vigésimo dia de dezembro", 0, "vigésimo dia de dezembro", 8376 * time.Hour},
		{"10º dia de março", 0, "10º dia de março", 1536 * time.Hour},
		{"4 jan.", 0, "4 jan.", -48 * time.Hour},
		{"fevereiro", 0, "fevereiro", 744 * time.Hour},
		{"outubro", 0, "outubro", 6576 * time.Hour},
		{"jul.", 0, "jul.", 4368 * time.Hour},
		{"junho", 0, "junho", 3648 * time.Hour},
	}

	ApplyFixtures(t, "br.ExactMonthDate", w, fixtok)
}
