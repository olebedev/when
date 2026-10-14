package en

import (
	"testing"
	"time"

	"github.com/olebedev/when/rules"
	"github.com/stretchr/testify/assert"
)

func TestTimezone(t *testing.T) {
	rule := Timezone(rules.Override)
	ref := time.Date(2026, 10, 14, 12, 0, 0, 0, time.UTC)

	tests := []struct {
		text         string
		expectedZone string
		matched      bool
	}{
		{"10 PM EST", "EST", true},
		{"10:30am PDT", "PDT", true},
		{"tomorrow at 5pm UTC", "UTC", true},
		{"12:00 +02:00", "+02:00", true},
		{"invalid", "", false},
	}

	for _, tc := range tests {
		m := rule.Find(tc.text)
		if !tc.matched {
			assert.Nil(t, m)
			continue
		}
		assert.NotNil(t, m)
		ctx := &rules.Context{}
		ok, err := m.Apply(ctx, nil, ref)
		assert.NoError(t, err)
		assert.True(t, ok)
		assert.NotNil(t, ctx.Location)
	}
}
