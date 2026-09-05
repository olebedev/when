package rules

import "time"

type Context struct {
	Text string

	// accumulator of relative values
	Duration time.Duration

	// Aboslute values
	Year, Month, Weekday, Day, Hour, Minute, Second *int

	Location *time.Location
}

func (c *Context) Time(t time.Time) (time.Time, error) {
	if t.IsZero() {
		t = time.Now()
	}

	if c.Duration != 0 {
		t = t.Add(c.Duration)
	}

	if c.Year != nil {
		t = time.Date(*c.Year, t.Month(), t.Day(), t.Hour(),
			t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	if c.Month != nil {
		month := time.Month(*c.Month)

		// time.Date would normalize a day the target month lacks into the next
		// one; clamp instead.
		day := t.Day()
		if last := daysIn(month, t.Year()); day > last {
			day = last
		}

		t = time.Date(t.Year(), month, day,
			t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	if c.Weekday != nil {
		diff := int(time.Weekday(*c.Weekday) - t.Weekday())
		t = time.Date(t.Year(), t.Month(), t.Day()+diff, t.Hour(),
			t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	if c.Day != nil {
		t = time.Date(t.Year(), t.Month(), *c.Day, t.Hour(),
			t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	if c.Hour != nil {
		t = time.Date(t.Year(), t.Month(), t.Day(), *c.Hour,
			t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	if c.Minute != nil {
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(),
			*c.Minute, t.Second(), t.Nanosecond(), t.Location())
	}

	if c.Second != nil {
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(),
			t.Minute(), *c.Second, t.Nanosecond(), t.Location())
	}

	if c.Location != nil {
		t = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(),
			t.Minute(), t.Second(), t.Nanosecond(), c.Location)
	}

	return t, nil
}

// daysIn returns the number of days in the given month of the given year.
func daysIn(m time.Month, year int) int {
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}
