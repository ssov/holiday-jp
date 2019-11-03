package holiday

import (
	"testing"
	"time"
)

func BenchmarkIsHoliday(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsHoliday(time.Date(1955, time.January, 1, 0, 0, 0, 0, jst))
	}
}

func TestIsHoliday(t *testing.T) {
	tt := []struct {
		in  time.Time
		out bool
	}{
		{
			in:  time.Date(1954, time.December, 12, 31, 23, 59, 999999, jst),
			out: false,
		},
		{
			in:  time.Date(1955, time.January, 1, 0, 0, 0, 0, jst),
			out: true,
		},
		{
			in:  time.Date(1955, time.January, 1, 23, 59, 59, 999999, jst),
			out: true,
		},
		{
			in:  time.Date(1955, time.January, 2, 0, 0, 0, 0, jst),
			out: false,
		},

		{
			in:  time.Date(1954, time.December, 31, 14, 59, 59, 999999, time.UTC),
			out: false,
		},
		{
			in:  time.Date(1954, time.December, 31, 15, 0, 0, 0, time.UTC),
			out: true,
		},
		{
			in:  time.Date(1955, time.January, 1, 14, 59, 59, 999999, time.UTC),
			out: true,
		},
		{
			in:  time.Date(1955, time.January, 1, 15, 0, 0, 0, time.UTC),
			out: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.in.String(), func(t *testing.T) {
			if got := IsHoliday(tc.in); tc.out != got {
				t.Errorf("got %t; want %t", got, tc.out)
			}
		})
	}
}
