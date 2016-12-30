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
		{"Это нужно сделать вчера вечером", 33, "вчера", -(time.Hour * 24)},
	}

	w := when.New(nil)
	w.Add(ru.CasualDate(rules.Skip))

	ApplyFixtures(t, "ru.CasualDate", w, fixt)
}

// func TestCasualTime(t *testing.T) {
// 	fixt := []Fixture{
// 		{"The Deadline was this morning ", 17, "this morning", 8 * time.Hour},
// 		{"The Deadline was this noon ", 17, "this noon", 12 * time.Hour},
// 		{"The Deadline was this afternoon ", 17, "this afternoon", 15 * time.Hour},
// 		{"The Deadline was this evening ", 17, "this evening", 18 * time.Hour},
// 	}
//
// 	w := when.New(nil)
// 	w.Add(en.CasualTime(rules.Skip))
//
// 	ApplyFixtures(t, "en.CasualTime", w, fixt)
// }
//
// func TestCasualDateCasualTime(t *testing.T) {
// 	fixt := []Fixture{
// 		{"Это нужно сделать tomorrow this afternoon ", 16, "tomorrow this afternoon", (15 + 24) * time.Hour},
// 	}
//
// 	w := when.New(nil)
// 	w.Add(
// 		en.CasualDate(rules.Skip),
// 		en.CasualTime(rules.OverWrite),
// 	)
//
// 	ApplyFixtures(t, "en.CasualDate|en.CasualTime", w, fixt)
// }
