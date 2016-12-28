package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/en"
	"github.com/stretchr/testify/require"
)

func TestCasualDate(t *testing.T) {
	fixt := []Fixture{
		{"The Deadline is now, ok", 16, "now", 0},
		{"The Deadline is today", 16, "today", 0},
		{"The Deadline is tonight", 16, "tonight", 23 * time.Hour},
		{"The Deadline is tomorrow everning", 16, "tomorrow", time.Hour * 24},
		{"The Deadline is yesterday everning", 16, "yesterday", -(time.Hour * 24)},
	}

	w := when.New(nil)
	w.Add(en.CasualDate())

	for i, f := range fixt {
		ti, index, txt, err := w.Parse(f.Text, null)
		require.Nil(t, err, "#%d", i)
		require.Equal(t, f.Index, index, "#%d", i)
		require.Equal(t, f.Phrase, txt, "#%d", i)
		require.Equal(t, f.Diff, ti.Sub(null), "#%d", i)
	}
}

// func TestCasualTime(t *testing.T) {
// 	fixt := []Fixture{
// 		{"The Deadline was this morning ", 16, " this morning", 8 * time.Hour},
// 	}

// 	w := when.New(nil)
// 	w.Add(en.CasualDate())

// 	for _, f := range fixt {
// 		ti, index, txt, err := w.Parse(f.Text, null)
// 		require.Nil(t, err)
// 		require.Equal(t, f.Index, index)
// 		require.Equal(t, f.Phrase, txt)
// 		require.Equal(t, f.Diff, ti.Sub(null))
// 	}
// }

// /// //// // / / / /
// func TestCasualTime2(t *testing.T) {
// 	text := "The Deadline was this morning "
// 	w := when.New(nil)
// 	w.Add(en.CasualTime())
//
// 	ti, index, txt, err := w.Parse(text, null)
//
// 	require.Nil(t, err)
// 	require.Equal(t, 16, index)
// 	require.Equal(t, " this morning", txt)
// 	require.Equal(t, ti.Hour(), 8)
// 	require.Equal(t, 8*time.Hour, ti.Sub(null))
// }
//
// func TestCasualTimeNoon(t *testing.T) {
// 	text := "The Deadline was this noon "
// 	w := when.New(nil)
// 	w.Add(en.CasualTime())
//
// 	ti, index, txt, err := w.Parse(text, null)
//
// 	require.Nil(t, err)
// 	require.Equal(t, ti.Hour(), 12)
// 	require.Equal(t, 16, index)
// 	require.Equal(t, " this noon", txt)
// 	require.Equal(t, 12*time.Hour, ti.Sub(null))
// }
//
// func TestCasualTimeAfternoon(t *testing.T) {
// 	text := "The Deadline was this afternoon "
// 	w := when.New(nil)
// 	w.Add(en.CasualTime())
//
// 	ti, index, txt, err := w.Parse(text, null)
//
// 	require.Nil(t, err)
// 	require.Equal(t, ti.Hour(), 15)
// 	require.Equal(t, 16, index)
// 	require.Equal(t, " this afternoon", txt)
// 	require.Equal(t, 15*time.Hour, ti.Sub(null))
// }
//
// func TestCasualTimeEvening(t *testing.T) {
// 	text := "The Deadline was this evening "
// 	w := when.New(nil)
// 	w.Add(en.CasualTime())
//
// 	ti, index, txt, err := w.Parse(text, null)
//
// 	require.Nil(t, err)
// 	require.Equal(t, ti.Hour(), 18)
// 	require.Equal(t, 16, index)
// 	require.Equal(t, " this evening", txt)
// 	require.Equal(t, 18*time.Hour, ti.Sub(null))
// }
//
// func TestCasualDateTimeAfternoon(t *testing.T) {
// 	text := "The Deadline is tomorrow afternoon "
// 	w := when.New(nil)
// 	w.Add(
// 		en.CasualTime(),
// 		en.CasualDate(),
// 	)
//
// 	ti, index, txt, err := w.Parse(text, null)
//
// 	require.Nil(t, err)
// 	require.Equal(t, ti.Hour(), 15)
// 	require.Equal(t, 15, index)
// 	require.Equal(t, " tomorrow afternoon", txt)
// 	require.Equal(t, (15+24)*time.Hour, ti.Sub(null))
//
// }
