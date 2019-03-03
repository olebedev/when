package br_test

import (
	"testing"
	"time"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/br"
	"github.com/stretchr/testify/require"
)

var null = time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC)

type Fixture struct {
	Text   string
	Index  int
	Phrase string
	Diff   time.Duration
}

func ApplyFixtures(t *testing.T, name string, w *when.Parser, fixt []Fixture) {
	for i, f := range fixt {
		res, err := w.Parse(f.Text, null)
		require.Nil(t, err, "[%s] err #%d", name, i)
		require.NotNil(t, res, "[%s] res #%d", name, i)
		require.Equal(t, f.Index, res.Index, "[%s] index #%d", name, i)
		require.Equal(t, f.Phrase, res.Text, "[%s] text #%d", name, i)
		require.Equal(t, f.Diff, res.Time.Sub(null), "[%s] diff #%d", name, i)
	}
}

func ApplyFixturesNil(t *testing.T, name string, w *when.Parser, fixt []Fixture) {
	for i, f := range fixt {
		res, err := w.Parse(f.Text, null)
		require.Nil(t, err, "[%s] err #%d", name, i)
		require.Nil(t, res, "[%s] res #%d", name, i)
	}
}

func ApplyFixturesErr(t *testing.T, name string, w *when.Parser, fixt []Fixture) {
	for i, f := range fixt {
		_, err := w.Parse(f.Text, null)
		require.NotNil(t, err, "[%s] err #%d", name, i)
		require.Equal(t, f.Phrase, err.Error(), "[%s] err text #%d", name, i)
	}
}

func TestAll(t *testing.T) {
	w := when.New(nil)
	w.Add(br.All...)

	// complex cases
	fixt := []Fixture{
		{"hoje de noite às 11:10 pm", 0, "hoje de noite às 11:10 pm", (23 * time.Hour) + (10 * time.Minute)},
		{"na tarde de sexta", 3, "tarde de sexta", ((2 * 24) + 15) * time.Hour},
		{"na próxima terça às 14:00", 3, "próxima terça às 14:00", ((6 * 24) + 14) * time.Hour},
		{"na próxima terça às 2p", 3, "próxima terça às 2p", ((6 * 24) + 14) * time.Hour},
		{"na próxima quarta-feira às 2:25 p.m.", 3, "próxima quarta-feira às 2:25 p.m.", (((7 * 24) + 14) * time.Hour) + (25 * time.Minute)},
		{"11 am última terça", 0, "11 am última terça", -13 * time.Hour},
	}

	ApplyFixtures(t, "br.All...", w, fixt)
}
