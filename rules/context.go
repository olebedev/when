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
		t = time.Date(t.Year(), time.Month(*c.Month), t.Day(),
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
