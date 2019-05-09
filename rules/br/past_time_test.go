package br_test

import (
	"testing"
	"time"

	"github.com/omniboost/when"
	"github.com/omniboost/when/rules"
	"github.com/omniboost/when/rules/br"
)

func TestPastTime(t *testing.T) {
	fixt := []Fixture{
		{"meia hora atrás", 0, "meia hora atrás", -(time.Hour / 2)},
		{"1 hora atrás", 0, "1 hora atrás", -(time.Hour)},
		{"5 minutos atrás", 0, "5 minutos atrás", -(time.Minute * 5)},
		{"5 minutos atrás eu fui ao zoológico", 0, "5 minutos atrás", -(time.Minute * 5)},
		{"nós fizemos algo 10 dias atrás.", 18, "10 dias atrás", -(10 * 24 * time.Hour)},
		{"nós fizemos algo cinco dias atrás.", 18, "cinco dias atrás", -(5 * 24 * time.Hour)},
		{"fizemos algo 5 dias atrás.", 13, "5 dias atrás", -(5 * 24 * time.Hour)},
		{"5 segundos atrás, um carro foi movido", 0, "5 segundos atrás", -(5 * time.Second)},
		{"duas semanas atrás", 0, "duas semanas atrás", -(14 * 24 * time.Hour)},
		{"um mês atrás", 0, "um mês atrás", -(31 * 24 * time.Hour)},
		{"uns meses atrás", 0, "uns meses atrás", -(92 * 24 * time.Hour)},
		{"há um ano", 4, "um ano", -(365 * 24 * time.Hour)},
		{"há duas semanas", 4, "duas semanas", -(2 * 7 * 24 * time.Hour)},
		{"poucas semanas atrás", 0, "poucas semanas atrás", -(3 * 7 * 24 * time.Hour)},
		{"há poucas semanas", 4, "poucas semanas", -(3 * 7 * 24 * time.Hour)},
		{"alguns dias atrás", 0, "alguns dias atrás", -(3 * 24 * time.Hour)},
		{"há alguns dias", 4, "alguns dias", -(3 * 24 * time.Hour)},
	}

	w := when.New(nil)
	w.Add(br.PastTime(rules.Skip))

	ApplyFixtures(t, "br.PastTime", w, fixt)
}
