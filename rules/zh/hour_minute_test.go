package zh_test

import (
	"github.com/olebedev/when/rules/zh"
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
)

func TestHourMinute(t *testing.T) {
	// current is Monday
	fixt := []Fixture{
		{"上午 11:30", 0, "上午 11:30", 11*time.Hour + 30*time.Minute},
		{"下午 3:30", 0, "下午 3:30", 15*time.Hour + 30*time.Minute},
		{"下午 3点半", 0, "下午 3点半", 15*time.Hour + 30*time.Minute},
		{"凌晨 3点半", 0, "凌晨 3点半", 3*time.Hour + 30*time.Minute},
		{"晚上8:00", 0, "晚上8:00", 20*time.Hour + 0*time.Minute},
		{"晚上9:32", 0, "晚上9:32", 21*time.Hour + 32*time.Minute},
		{"晚 上 8:00", 0, "晚 上 8:00", 20*time.Hour + 0*time.Minute},
		{"晚上 8 点干啥去", 0, "晚上 8 点", 20*time.Hour + 0*time.Minute},
		{"他俩凌晨 3点去散步太可怕了", 6, "凌晨 3点", 3*time.Hour + 0*time.Minute},
		{"早晨八点一刻", 0, "早晨八点一刻", 8*time.Hour + 15*time.Minute},
		{"早上八点半", 0, "早上八点半", 8*time.Hour + 30*time.Minute},
		{"今晚八点", 0, "今晚八点", 20 * time.Hour},
		{"今晚八点半", 0, "今晚八点半", 20*time.Hour + 30*time.Minute},
	}

	w := when.New(nil)

	w.Add(zh.HourMinute(rules.Override))

	ApplyFixtures(t, "zh.HourMinute", w, fixt)
}
