package zh_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
	"github.com/olebedev/when/rules/zh"
)

func TestCasualDate(t *testing.T) {
	fixt := []Fixture{
		{"后天中午", 0, "后天", (2 * 24) * time.Hour},
		{"大后天中午", 0, "大后天", (3 * 24) * time.Hour},
		{"昨天", 0, "昨天", (-1 * 24) * time.Hour},
		{"前天", 0, "前天", (-2 * 24) * time.Hour},
		{"大前天", 0, "大前天", (-3 * 24) * time.Hour},
		{"下月", 0, "下月", (31 * 24) * time.Hour},
		{"下个月", 0, "下个月", (31 * 24) * time.Hour},
		{"下下月", 0, "下下月", (31*24 + 30*24) * time.Hour},
		{"下下个月", 0, "下下个月", (31*24 + 30*24) * time.Hour},
		{"明年", 0, "明年", (365 * 24) * time.Hour},
		{"后年", 0, "后年", now.AddDate(2, 0, 0).Sub(now)},
		{"下月6号", 0, "下月6号", 552 * time.Hour},
	}

	w := when.New(nil)

	w.Add(zh.CasualDate(rules.Override))

	ApplyFixtures(t, "zh.TestCasualDate", w, fixt)
}

/*
	(([1-9](?:月|-|/|\.|))|1[0-2])\s*(月|-|/|\.|)\s*([1-9]|1[0-9]|2[0-9]|3[0-1])\s*(日|号)?(?:\W|$)
*/
