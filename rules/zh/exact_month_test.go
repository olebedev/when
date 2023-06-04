package zh_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/zh"
)

func TestExactMonthDate(t *testing.T) {
	// current is Monday
	fixt := []Fixture{
		{"4月1日", 0, "4月1日", (18 * 24) * time.Hour},
		{"4月2日", 0, "4月2日", (19 * 24) * time.Hour},
		{"4月 2日", 0, "4月 2日", (19 * 24) * time.Hour},
		{"4 月 2 日", 0, "4 月 2 日", (19 * 24) * time.Hour},
		{"四月一日", 0, "四月一日", (18 * 24) * time.Hour},
		{"四月1日", 0, "四月1日", (18 * 24) * time.Hour},
		{"四月", 0, "四月", (18 * 24) * time.Hour},
		{"十一月一日", 0, "十一月一日", 5568 * time.Hour},
		{"四月三十日", 0, "四月三十日", 1128 * time.Hour},
		{"4月30日", 0, "4月30日", 1128 * time.Hour},
		{"5月1号", 0, "5月1号", 1152 * time.Hour},
		{"5/1", 0, "5/1", 1152 * time.Hour},
		{"5月1日", 0, "5月1日", 1152 * time.Hour},
		{"五月", 0, "五月", 1152 * time.Hour},
		{"12号", 0, "12号", (-2 * 24) * time.Hour},
	}

	w := when.New(nil)

	w.Add(zh.ExactMonthDate(rules.Override))

	ApplyFixtures(t, "zh.ExactMonthDate", w, fixt)
}

func TestExactMonthDateNil(t *testing.T) {
	fixt := []Fixture{
		{"41", 0, "", (18 * 24) * time.Hour},
	}

	w := when.New(nil)

	w.Add(zh.ExactMonthDate(rules.Override))
	ApplyFixturesNil(t, "zh.ExactMonthDate", w, fixt)

}

/*
	(([1-9](?:月|-|/|\.|))|1[0-2])\s*(月|-|/|\.|)\s*([1-9]|1[0-9]|2[0-9]|3[0-1])\s*(日|号)?(?:\W|$)
*/
