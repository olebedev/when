package br_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/br"
)

func TestDeadline(t *testing.T) {
	fixt := []Fixture{
		{"dentro de meia hora", 0, "dentro de meia hora", time.Hour / 2},
		{"dentro de 1 hora", 0, "dentro de 1 hora", time.Hour},
		{"em 5 minutos", 0, "em 5 minutos", time.Minute * 5},
		{"Em 5 minutos eu irei para casa", 0, "Em 5 minutos", time.Minute * 5},
		{"nós precisamos fazer algo dentro de 10 dias.", 27, "dentro de 10 dias", 10 * 24 * time.Hour},
		{"nós temos que fazer algo em cinco dias.", 26, "em cinco dias", 5 * 24 * time.Hour},
		{"nós temos que fazer algo em 5 dias.", 26, "em 5 dias", 5 * 24 * time.Hour},
		{"Em 5 segundos, um carro precisa se mover", 0, "Em 5 segundos", 5 * time.Second},
		{"dentro de duas semanas", 0, "dentro de duas semanas", 14 * 24 * time.Hour},
		{"dentro de um mês", 0, "dentro de um mês", 31 * 24 * time.Hour},
		{"dentro de alguns meses", 0, "dentro de alguns meses", 91 * 24 * time.Hour},
		{"dentro de poucos meses", 0, "dentro de poucos meses", 91 * 24 * time.Hour},
		{"dentro de um ano", 0, "dentro de um ano", 366 * 24 * time.Hour},
		{"em uma semana", 0, "em uma semana", 7 * 24 * time.Hour},
	}

	w := when.New(nil)
	w.Add(br.Deadline(rules.Skip))

	ApplyFixtures(t, "br.Deadline", w, fixt)
}
