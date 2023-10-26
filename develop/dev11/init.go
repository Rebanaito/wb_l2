package main

import (
	"time"
)

func makeYear(num uint16) (y year) {
	y.yearNum = num
	y.months = []month{{"January", makeMonth(num, time.January, 31)},
		{"February", makeMonth(num, time.February, 28+leapYear(num))},
		{"March", makeMonth(num, time.March, 31)},
		{"April", makeMonth(num, time.April, 30)},
		{"May", makeMonth(num, time.May, 31)},
		{"June", makeMonth(num, time.June, 30)},
		{"July", makeMonth(num, time.July, 31)},
		{"August", makeMonth(num, time.August, 31)},
		{"Septermber", makeMonth(num, time.September, 30)},
		{"October", makeMonth(num, time.October, 31)},
		{"November", makeMonth(num, time.November, 30)},
		{"December", makeMonth(num, time.December, 31)},
	}
	return
}

func makeMonth(year uint16, month time.Month, dayCount int) []day {
	startsWith := int(time.Date(int(year), month, 1, 0, 0, 0, 0, time.UTC).Weekday())
	days := make([]day, dayCount)
	for i := range days {
		days[i].weekday = time.Weekday((startsWith + i) % 7)
	}
	return days
}

func leapYear(num uint16) int {
	if num%4 == 0 {
		return 1
	}
	return 0
}

func initEnv() (c calendar, noEnv bool) {
	return calendar{}, true
}

func initYears() (c calendar) {
	c.years = make([]year, YEARS_MAX)
	currentYear := time.Now().Year()
	c.firstYear = uint16(currentYear)
	for i := 0; i < YEARS_MAX; i++ {
		c.years[i] = makeYear(uint16(currentYear + i))
	}
	return
}
