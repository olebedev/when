package en_test

import (
	"github.com/olebedev/when/rules/common"
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/en"
	"github.com/stretchr/testify/assert"
)

func TestCasualDate(t *testing.T) {
	fixt := []Fixture{
		{"The Deadline is now, ok", 16, "now", 0},
		{"The Deadline is today", 16, "today", 0},
		{"The Deadline is tonight", 16, "tonight", 23 * time.Hour},
		{"The Deadline is tomorrow evening", 16, "tomorrow", time.Hour * 24},
		{"The Deadline is yesterday evening", 16, "yesterday", -(time.Hour * 24)},
	}

	w := when.New(nil)
	w.Add(en.CasualDate(rules.Skip))

	ApplyFixtures(t, "en.CasualDate", w, fixt)
}

func TestTodayTomorrowYesterday(t *testing.T) {
	var w = when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)
	now := time.Now()
	r, err := w.Parse("today", now)
	assert.EqualValues(t, now.Year(), r.Time.Year())
	assert.EqualValues(t, now.Month(), r.Time.Month())
	assert.EqualValues(t, now.Day(), r.Time.Day())
	assert.EqualValues(t, 0, r.Time.Minute())
	assert.EqualValues(t, 0, r.Time.Hour())
	assert.EqualValues(t, 0, r.Time.Second())
	r, err = w.Parse("yesterday", now)
	if err != nil {
		t.Fail()
	}
	assert.EqualValues(t, now.Year(), r.Time.Year())
	assert.EqualValues(t, now.Month(), r.Time.Month())
	assert.EqualValues(t, now.Day()-1, r.Time.Day())
	assert.EqualValues(t, 0, r.Time.Minute())
	assert.EqualValues(t, 0, r.Time.Hour())
	assert.EqualValues(t, 0, r.Time.Second())

	r, err = w.Parse("tomorrow", now)
	if err != nil {
		t.Fail()
	}
	assert.EqualValues(t, now.Year(), r.Time.Year())
	assert.EqualValues(t, now.Month(), r.Time.Month())
	assert.EqualValues(t, now.Day()+1, r.Time.Day())
	assert.EqualValues(t, 0, r.Time.Minute())
	assert.EqualValues(t, 0, r.Time.Hour())
	assert.EqualValues(t, 0, r.Time.Second())

}

func TestCasualTime(t *testing.T) {
	fixt := []Fixture{
		{"The Deadline was this morning ", 17, "this morning", 8 * time.Hour},
		{"The Deadline was this noon ", 17, "this noon", 12 * time.Hour},
		{"The Deadline was this afternoon ", 17, "this afternoon", 15 * time.Hour},
		{"The Deadline was this evening ", 17, "this evening", 18 * time.Hour},
	}

	w := when.New(nil)
	w.Add(en.CasualTime(rules.Skip))

	ApplyFixtures(t, "en.CasualTime", w, fixt)
}

func TestCasualDateCasualTime(t *testing.T) {
	fixt := []Fixture{
		{"The Deadline is tomorrow this afternoon ", 16, "tomorrow this afternoon", (15 + 24) * time.Hour},
	}

	w := when.New(nil)
	w.Add(
		en.CasualDate(rules.Skip),
		en.CasualTime(rules.Override),
	)

	ApplyFixtures(t, "en.CasualDate|en.CasualTime", w, fixt)
}
