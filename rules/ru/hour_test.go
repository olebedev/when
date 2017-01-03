package ru_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/ru"
)

func TestHour(t *testing.T) {
	fixt := []Fixture{
		{"5вечера", 0, "5вечера", 17 * time.Hour},
		{"в 5 вечера", 3, "5 вечера", 17 * time.Hour},
		{"нужно к 5 часам вечера", 14, "5 часам вечера", 17 * time.Hour},
		{"в три часа дня", 3, "три часа дня", 15 * time.Hour},
		{"в час дня", 3, "час дня", 13 * time.Hour},
		{"в одиннадцать часов утра", 3, "одиннадцать часов утра", 11 * time.Hour},
		{"в семь вечера", 3, "семь вечера", 19 * time.Hour},
	}

	w := when.New(nil)
	w.Add(ru.Hour(rules.OverWrite))

	ApplyFixtures(t, "en.Hour", w, fixt)
}
