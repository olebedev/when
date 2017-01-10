package ru_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/ru"
)

func TestCasualDate(t *testing.T) {
	fixt := []Fixture{
		{"Это нужно сделать прямо сейчас", 44, "сейчас", 0},
		{"Это нужно сделать сегодня", 33, "сегодня", 0},
		{"Это нужно сделать завтра вечером", 33, "завтра", time.Hour * 24},
		{"Это нужно было сделать вчера вечером", 42, "вчера", -(time.Hour * 24)},
	}

	w := when.New(nil)
	w.Add(ru.CasualDate(rules.Skip))

	ApplyFixtures(t, "ru.CasualDate", w, fixt)
}

func TestCasualTime(t *testing.T) {
	fixt := []Fixture{
		{"Это нужно было сделать этим утром ", 42, "этим утром", 8 * time.Hour},
		{"Это нужно сделать до обеда", 33, "до обеда", 12 * time.Hour},
		{"Это нужно сделать после обеда", 33, "после обеда", 15 * time.Hour},
		{"Это нужно сделать к вечеру", 33, "к вечеру", 18 * time.Hour},
		{"вечером", 0, "вечером", 18 * time.Hour},
		{"вечером", 0, "вечером", 18 * time.Hour},
	}

	w := when.New(nil)
	w.Add(ru.CasualTime(rules.Skip))

	ApplyFixtures(t, "ru.CasualTime", w, fixt)
}

func TestCasualDateCasualTime(t *testing.T) {
	fixt := []Fixture{
		{"Это нужно сделать завтра после обеда", 33, "завтра после обеда", (15 + 24) * time.Hour},
		{"Это нужно сделать завтра утром", 33, "завтра утром", (8 + 24) * time.Hour},
		{"Это нужно было сделать вчера утром", 42, "вчера утром", (8 - 24) * time.Hour},
		{"Это нужно было сделать вчера после обеда", 42, "вчера после обеда", (15 - 24) * time.Hour},
		{"помыть окна до вечера", 22, "до вечера", 18 * time.Hour},
		{"помыть окна до обеда", 22, "до обеда", 12 * time.Hour},
		{"сделать это к вечеру", 22, "к вечеру", 18 * time.Hour},
		{"помыть окна завтра утром", 22, "завтра утром", 32 * time.Hour},
		{"написать письмо во вторник после обеда", 50, "после обеда", 15 * time.Hour},
		{"написать письмо до утра ", 30, "до утра", 8 * time.Hour},
		{"к вечеру", 0, "к вечеру", 18 * time.Hour},
	}

	w := when.New(nil)
	w.Add(
		ru.CasualDate(rules.Skip),
		ru.CasualTime(rules.OverWrite),
	)

	ApplyFixtures(t, "ru.CasualDate|ru.CasualTime", w, fixt)
}
