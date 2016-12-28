package en_test

// func TestDeadline(t *testing.T) {
// 	fixt := []Fixture{
// 		{"within half an hour", 0, "within half an hour", time.Hour / 2},
// 		{"within 1 hour", 0, "within 1 hour", time.Hour},
// 		{"in 5 minutes", 0, "in 5 minutes", time.Minute * 5},
// 		{"In 5 minutes I will go home", 0, "In 5 minutes", time.Minute * 5},
// 		{"we have to do something within 10 days.", 23, " within 10 days", 10 * 24 * time.Hour},
// 		{"we have to do something in five days.", 23, " in five days", 5 * 24 * time.Hour},
// 		{"we have to do something in 5 days.", 23, " in 5 days", 5 * 24 * time.Hour},
// 		{"In 5 seconds A car need to move", 0, "In 5 seconds", 5 * time.Second},
// 		{"within two weeks", 0, "within two weeks", 14 * 24 * time.Hour},
// 		{"within a month", 0, "within a month", 31 * 24 * time.Hour},
// 		{"within a few months", 0, "within a few months", 91 * 24 * time.Hour},
// 		{"within one year", 0, "within one year", 366 * 24 * time.Hour},
// 		{"in a week", 0, "in a week", 7 * 24 * time.Hour},
// 	}
//
// 	w := when.New(nil)
// 	w.Add(en.Deadline())
//
// 	for _, f := range fixt {
// 		ti, index, txt, err := w.Parse(f.Text, null)
// 		require.Nil(t, err)
// 		require.Equal(t, f.Index, index)
// 		require.Equal(t, f.Phrase, txt)
// 		require.Equal(t, f.Diff, ti.Sub(null))
// 	}
// }
