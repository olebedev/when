package en_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/en"
	"github.com/stretchr/testify/require"
)

var null = time.Date(0, 0, 0, 0, 0, 0, 0, time.Local)

func TestWhenCasualDate(t *testing.T) {
	w := when.New(nil)
	w.Add(en.NewCasualDate())

	_, index, text, err := w.When("The Deadline is now, ok", null)

	require.Nil(t, err)
	require.Equal(t, 15, index)
	require.Equal(t, " now", text)
}

func TestCasualDateToday(t *testing.T) {

	text := "The Deadline is today"
	w := when.New(nil)
	w.Add(en.NewCasualDate())

	ti, index, text, err := w.When(text, null)

	require.Nil(t, err)
	require.Equal(t, ti.Hour(), 18)
	require.Equal(t, 15, index)
	require.Equal(t, " today", text)
	require.Equal(t, time.Hour*18, ti.Sub(null))
}

func TestCasualDateTonight(t *testing.T) {

	text := "The Deadline is tonight"
	w := when.New(nil)
	w.Add(en.NewCasualDate())

	ti, index, text, err := w.When(text, null)

	require.Nil(t, err)
	require.Equal(t, ti.Hour(), 23)
	require.Equal(t, 15, index)
	require.Equal(t, " tonight", text)
	require.Equal(t, time.Hour*23, ti.Sub(null))
}

func TestCasualDateTomorrow(t *testing.T) {

	text := "The Deadline is tomorrow everning"
	w := when.New(nil)
	w.Add(en.NewCasualDate())

	ti, index, text, err := w.When(text, null)

	require.Nil(t, err)
	require.Equal(t, ti.Hour(), 0)
	require.Equal(t, 15, index)
	require.Equal(t, " tomorrow", text)
	require.Equal(t, time.Hour*24, ti.Sub(null))
}

func TestCasualDateYesterday(t *testing.T) {

	text := "The Deadline is yesterday everning"
	w := when.New(nil)
	w.Add(en.NewCasualDate())

	ti, index, text, err := w.When(text, time.Now())

	require.Nil(t, err)
	require.Equal(t, ti.Hour(), time.Now().Hour())
	require.Equal(t, 15, index)
	require.Equal(t, " yesterday", text)
	require.Equal(t, time.Now().Day()-1, ti.Day())
}
