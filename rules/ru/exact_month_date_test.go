package ru_test

import (
	"github.com/olebedev/when/rules/ru"
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
)

func TestExactMonthDate(t *testing.T) {
	w := when.New(nil)
	w.Add(ru.ExactMonthDate(rules.Override))

	fixtok := []Fixture{
		{"3 марта", 0, "3 марта", 1368 * time.Hour},
		{"1 сентября", 0, "1 сентября", 5736 * time.Hour},
		{"1 сент", 0, "1 сен", 5736 * time.Hour},
		{"1 сент.", 0, "1 сен", 5736 * time.Hour},
		{"21 окт", 0, "21 окт", 6936 * time.Hour},
		{"июня", 0, "июня", 3648 * time.Hour},
	}

	ApplyFixtures(t, "ru.ExactMonthDate", w, fixtok)
}
