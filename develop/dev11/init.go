package main

import (
	"time"
)

func makeYear(num uint16) (y year) {
	y.yearNum = num
	y.months = []month{{"January", make([]day, 31)},
		{"February", make([]day, 28+leapYear(num))},
		{"March", make([]day, 31)},
		{"April", make([]day, 30)},
		{"May", make([]day, 31)},
		{"June", make([]day, 30)},
		{"July", make([]day, 31)},
		{"August", make([]day, 31)},
		{"Septermber", make([]day, 30)},
		{"October", make([]day, 31)},
		{"November", make([]day, 30)},
		{"December", make([]day, 31)},
	}
	return
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
