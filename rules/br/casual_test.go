package br_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/br"
)

func TestCasualDate(t *testing.T) {
	fixt := []Fixture{
		{"O prazo final é agora, ok", 17, "agora", 0},
		{"O prazo final é hoje", 17, "hoje", 0},
		{"O prazo final é esta noite", 17, "esta noite", 23 * time.Hour},
		{"O prazo final é amanhã à noite", 17, "amanhã ", time.Hour * 24},
		{"O prazo foi ontem à noite", 12, "ontem ", -(time.Hour * 24)},
	}

	w := when.New(nil)
	w.Add(br.CasualDate(rules.Skip))

	ApplyFixtures(t, "br.CasualDate", w, fixt)
}

func TestCasualTime(t *testing.T) {
	fixt := []Fixture{
		{"O prazo foi esta manhã ", 12, "esta manhã", 8 * time.Hour},
		{"O prazo final foi ao meio-dia ", 18, "ao meio-dia", 12 * time.Hour},
		{"O prazo final foi esta tarde ", 18, "esta tarde", 15 * time.Hour},
		{"O prazo foi nesta noite ", 12, "nesta noite", 18 * time.Hour},
	}

	w := when.New(nil)
	w.Add(br.CasualTime(rules.Skip))

	ApplyFixtures(t, "br.CasualTime", w, fixt)
}

func TestCasualDateCasualTime(t *testing.T) {
	fixt := []Fixture{
		{"O prazo final é amanhã de tarde", 17, "amanhã de tarde", (15 + 24) * time.Hour},
	}

	w := when.New(nil)
	w.Add(
		br.CasualDate(rules.Skip),
		br.CasualTime(rules.Override),
	)

	ApplyFixtures(t, "br.CasualDate|br.CasualTime", w, fixt)
}
