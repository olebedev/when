package br_test

import (
	"testing"
	"time"

	"github.com/omniboost/when"
	"github.com/omniboost/when/rules"
	"github.com/omniboost/when/rules/br"
)

func TestWeekday(t *testing.T) {
	// current is Friday
	fixt := []Fixture{
		// past/last
		{"faça isto para a Segunda passada", 18, "Segunda passada", -(2 * 24 * time.Hour)},
		{"sábado passado", 0, "sábado passado", -(4 * 24 * time.Hour)},
		{"sexta-feira passada", 0, "sexta-feira passada", -(5 * 24 * time.Hour)},
		{"quarta-feira passada", 0, "quarta-feira passada", -(7 * 24 * time.Hour)},
		{"terça passada", 0, "terça passada", -(24 * time.Hour)},
		// // next
		{"na próxima terça-feira", 3, "próxima terça-feira", 6 * 24 * time.Hour},
		{"me ligue na próxima quarta", 12, "próxima quarta", 7 * 24 * time.Hour},
		{"sábado que vem", 0, "sábado que vem", 3 * 24 * time.Hour},
		// // this
		{"essa terça-feira", 0, "essa terça-feira", -(24 * time.Hour)},
		{"liga pra mim nesta quarta", 13, "nesta quarta", 0},
		{"neste sábado", 0, "neste sábado", 3 * 24 * time.Hour},
	}

	w := when.New(nil)

	w.Add(br.Weekday(rules.Override))

	ApplyFixtures(t, "br.Weekday", w, fixt)
}
