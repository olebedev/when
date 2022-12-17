package zh_test

import (
	"github.com/olebedev/when/rules/zh"
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
)

func TestTraditionHour(t *testing.T) {
	// current is Monday
	fixt := []Fixture{
		{"午 时123", 0, "午 时", 11 * time.Hour},
		{"子时", 0, "子时", 23 * time.Hour},
		{"午时太阳正好", 0, "午时", 11 * time.Hour},
		{"我们在酉时喝一杯吧", 9, "酉时", 17 * time.Hour},
		{"午时三刻问斩", 0, "午时三刻", 11*time.Hour + 45*time.Minute},
		{"午时四刻吃饭", 0, "午时四刻", 12 * time.Hour},
		{"戌时1刻", 0, "戌时1刻", 19*time.Hour + 15*time.Minute},
	}

	w := when.New(nil)

	w.Add(zh.TraditionHour(rules.Override))

	ApplyFixtures(t, "zh.TraditionHour", w, fixt)
}
