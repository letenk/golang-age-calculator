package main

import (
	"fmt"
	"time"
)

func calculateAgeContent(day int, month int, year int) string {
	var ageContent string

	if day == 0 && month == 0 && year == 0 {
		ageContent = fmt.Sprintf("%d Day", 0)
	}

	if day != 0 {
		ageContent = fmt.Sprintf("%d Day", day)
	}
	if month != 0 {
		ageContent = fmt.Sprintf("%d Month %d Day", month, day)
	}
	if year != 0 {
		ageContent = fmt.Sprintf("%d Year %d Month %d Day", year, month, day)
	}

	return ageContent
}

func diffDate(from, to time.Time) (inverted bool, years, months, days, hours, minutes, seconds, nanoseconds int) {
	if from.Location() != to.Location() {
		to = to.In(to.Location())
	}

	inverted = false
	if from.After(to) {
		inverted = true
		from, to = to, from
	}

	y1, M1, d1 := from.Date()
	y2, M2, d2 := to.Date()

	h1, m1, s1 := from.Clock()
	h2, m2, s2 := to.Clock()

	ns1, ns2 := from.Nanosecond(), to.Nanosecond()

	years = y2 - y1
	months = int(M2 - M1)
	days = d2 - d1

	hours = h2 - h1
	minutes = m2 - m1
	seconds = s2 - s1
	nanoseconds = ns2 - ns1

	if nanoseconds < 0 {
		nanoseconds += 1e9
		seconds--
	}
	if seconds < 0 {
		seconds += 60
		minutes--
	}
	if minutes < 0 {
		minutes += 60
		hours--
	}
	if hours < 0 {
		hours += 24
		days--
	}

	if days < 0 {
		t := time.Date(y2, M2, 0, 0, 0, 0, 0, time.UTC)
		days += t.Day()
		months--
	}

	if months < 0 {
		months += 12
		years--
	}
	return
}

func main() {
	// "2023-09-07"
	birthdayStr := "2022-03-09"
	birthday, err := time.Parse("2006-01-02", birthdayStr)

	if err != nil {
		fmt.Println("Wrong format:", err)
		return
	}

	now := time.Now()

	age := now.Sub(birthday)
	ageDay := (uint)(age.Hours() / 24)

	_, ageInYears, ageInMonths, ageInDay, _, _, _, _ := diffDate(now, birthday)

	ageContent := calculateAgeContent(ageInDay, ageInMonths, ageInYears)

	fmt.Printf("Age in day: %d\n", ageDay)
	fmt.Printf("Your age now: %s\n", ageContent)
}
