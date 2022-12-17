package zh_test

import (
	"github.com/olebedev/when/rules/zh"
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules"
)

func TestWeekday(t *testing.T) {
	// current is Monday
	fixt := []Fixture{
		{"和你下周一吃饭", 6, "下周一", 7 * 24 * time.Hour},
		{"下星期三", 0, "下星期三", 9 * 24 * time.Hour},
		{"和小西本周三一起打羽毛球", 9, "本周三", 2 * 24 * time.Hour},
		{"这周三", 0, "这周三", 2 * 24 * time.Hour},
		{"这礼拜四浇花", 0, "这礼拜四", 3 * 24 * time.Hour},
		{"这星期 4", 0, "这星期 4", 3 * 24 * time.Hour},
		{"和李星期这星期 4喝茶", 12, "这星期 4", 3 * 24 * time.Hour},
		{"周日", 0, "周日", 6 * 24 * time.Hour},
		{"下周日", 0, "下周日", (6 + 7) * 24 * time.Hour},
		{"2下周天", 1, "下周天", (6 + 7) * 24 * time.Hour},
		{"上周三", 0, "上周三", -5 * 24 * time.Hour},
		{"下个周三", 0, "下个周三", (7 + 2) * 24 * time.Hour},
		{"1下个礼拜 3", 1, "下个礼拜 3", (7 + 2) * 24 * time.Hour},
		{"下下礼拜 3", 0, "下下礼拜 3", (7 + 7 + 2) * 24 * time.Hour},
	}

	w := when.New(nil)

	w.Add(zh.Weekday(rules.Override))

	ApplyFixtures(t, "zh.Weekday", w, fixt)
}
