package en_test

import "time"

var null = time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC)

type Fixture struct {
	Text   string
	Index  int
	Phrase string
	Diff   time.Duration
}
