package main

import (
	"testing"
	"time"
)

func TestDiff(t *testing.T) {

	testCases := []struct {
		testName                         string // name of the test case
		birthday                         time.Time
		timeDateNow                      time.Time
		year, month, day, hour, min, sec int // expected outputs
	}{
		// Case total month 28
		{
			"total day 28 - 0 year, 1 month, 16 day",
			time.Date(2023, 02, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 03, 30, 0, 0, 0, 0, time.UTC),
			0, 1, 16, 0, 0, 0,
		},
		{
			"total day 28 - 0 year, 1 month, 17 day",
			time.Date(2023, 02, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 03, 31, 0, 0, 0, 0, time.UTC),
			0, 1, 17, 0, 0, 0,
		},
		{
			"total day 28 - 0 year, 1 month, 18 day",
			time.Date(2023, 02, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 04, 1, 0, 0, 0, 0, time.UTC),
			0, 1, 18, 0, 0, 0,
		},
		{
			"total day 28 - 0 year, 1 month, 19 day",
			time.Date(2023, 02, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 04, 2, 0, 0, 0, 0, time.UTC),
			0, 1, 19, 0, 0, 0,
		},
		{
			"total day 28 - 0 year, 2 month, 16 day",
			time.Date(2023, 02, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 04, 30, 0, 0, 0, 0, time.UTC),
			0, 2, 16, 0, 0, 0,
		},
		{
			"total day 28 - 0 year, 2 month, 17 day",
			time.Date(2023, 02, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 05, 1, 0, 0, 0, 0, time.UTC),
			0, 2, 17, 0, 0, 0,
		},
		{
			"total day 28 - 0 year, 1 month, 18 day",
			time.Date(2023, 02, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 05, 2, 0, 0, 0, 0, time.UTC),
			0, 2, 18, 0, 0, 0,
		},

		// Case total month 29
		{
			"total day 29 - 0 year, 1 month, 16 day",
			time.Date(2020, 02, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2020, 03, 30, 0, 0, 0, 0, time.UTC),
			0, 1, 16, 0, 0, 0,
		},
		{
			"total day 29 - 0 year, 1 month, 17 day",
			time.Date(2020, 02, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2020, 03, 31, 0, 0, 0, 0, time.UTC),
			0, 1, 17, 0, 0, 0,
		},
		{
			"total day 29 - 0 year, 1 month, 18 day",
			time.Date(2020, 02, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2020, 04, 1, 0, 0, 0, 0, time.UTC),
			0, 1, 18, 0, 0, 0,
		},
		{
			"total day 29 - 0 year, 1 month, 19 day",
			time.Date(2020, 02, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2020, 04, 2, 0, 0, 0, 0, time.UTC),
			0, 1, 19, 0, 0, 0,
		},

		{
			"total day 29 - 0 year, 2 month, 16 day",
			time.Date(2020, 02, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2020, 04, 30, 0, 0, 0, 0, time.UTC),
			0, 2, 16, 0, 0, 0,
		},
		{
			"total day 29 - 0 year, 2 month, 17 day",
			time.Date(2020, 02, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2020, 05, 1, 0, 0, 0, 0, time.UTC),
			0, 2, 17, 0, 0, 0,
		},
		{
			"total day 29 - 0 year, 2 month, 18 day",
			time.Date(2020, 02, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2020, 05, 2, 0, 0, 0, 0, time.UTC),
			0, 2, 18, 0, 0, 0,
		},

		// Case total month 30
		{
			"total day 30 - 0 year, 1 month, 16 day",
			time.Date(2023, 9, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 10, 30, 0, 0, 0, 0, time.UTC),
			0, 1, 16, 0, 0, 0,
		},
		{
			"total day 30 - 0 year, 1 month, 17 day",
			time.Date(2023, 9, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 10, 31, 0, 0, 0, 0, time.UTC),
			0, 1, 17, 0, 0, 0,
		},
		{
			"total day 30 - 0 year, 1 month, 18 day",
			time.Date(2023, 9, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 11, 1, 0, 0, 0, 0, time.UTC),
			0, 1, 18, 0, 0, 0,
		},
		{
			"total day 30 - 0 year, 1 month, 19 day",
			time.Date(2023, 9, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 11, 2, 0, 0, 0, 0, time.UTC),
			0, 1, 19, 0, 0, 0,
		},
		{
			"total day 30 - 0 year, 2 month, 16 day",
			time.Date(2023, 9, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 11, 30, 0, 0, 0, 0, time.UTC),
			0, 2, 16, 0, 0, 0,
		},
		{
			"total day 30 - 0 year, 2 month, 17 day",
			time.Date(2023, 9, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
			0, 2, 17, 0, 0, 0,
		},
		{
			"total day 30 - 0 year, 2 month, 18 day",
			time.Date(2023, 9, 14, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 12, 2, 0, 0, 0, 0, time.UTC),
			0, 2, 18, 0, 0, 0,
		},

		// Case total month 31
		{
			"total day 31 - 0 year, 1 month, 16 day",
			time.Date(2023, 10, 13, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 11, 29, 0, 0, 0, 0, time.UTC),
			0, 1, 16, 0, 0, 0,
		},
		{
			"total day 31 - 0 year, 1 month, 17 day",
			time.Date(2023, 10, 13, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 11, 30, 0, 0, 0, 0, time.UTC),
			0, 1, 17, 0, 0, 0,
		},
		{
			"total day 31 - 0 year, 1 month, 18 day",
			time.Date(2023, 10, 13, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 12, 1, 0, 0, 0, 0, time.UTC),
			0, 1, 18, 0, 0, 0,
		},
		{
			"total day 31 - 0 year, 1 month, 19 day",
			time.Date(2023, 10, 13, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 12, 2, 0, 0, 0, 0, time.UTC),
			0, 1, 19, 0, 0, 0,
		},
		{
			"total day 31 - 0 year, 2 month, 16 day",
			time.Date(2023, 10, 13, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 12, 29, 0, 0, 0, 0, time.UTC),
			0, 2, 16, 0, 0, 0,
		},
		{
			"total day 31 - 0 year, 2 month, 17 day",
			time.Date(2023, 10, 13, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 12, 30, 0, 0, 0, 0, time.UTC),
			0, 2, 17, 0, 0, 0,
		},
		{
			"total day 31 - 0 year, 2 month, 18 day",
			time.Date(2023, 10, 13, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC),
			0, 2, 18, 0, 0, 0,
		},
		{
			"total day 31 - 0 year, 2 month, 19 day",
			time.Date(2023, 10, 13, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			0, 2, 19, 0, 0, 0,
		},
		{
			"total day 31 - 0 year, 2 month, 20 day",
			time.Date(2023, 10, 13, 0, 0, 0, 0, time.UTC),
			time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC),
			0, 2, 20, 0, 0, 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			_, year, month, day, hour, min, sec, _ := diffDate(tc.birthday, tc.timeDateNow)
			if year != tc.year || month != tc.month || day != tc.day || hour != tc.hour || min != tc.min || sec != tc.sec {
				t.Errorf("[%s] Expected: %d years, %d months, %d days, %d hours, %d mins and %d seconds, got: %d years, %d months, %d days, %d hours, %d mins and %d seconds",
					tc.testName,
					tc.year, tc.month, tc.day, tc.hour, tc.min, tc.sec,
					year, month, day, hour, min, sec,
				)
			}
		})
	}
}
