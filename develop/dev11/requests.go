package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

const TEMPLATE_TOP string = `<html>
<style>
	table, th, td {
		border: 1px solid black;
		text-align: center;
		padding: 5px;
	}
</style>
<div style="text-align: center">
	<form action="/events_for_day">
		<button>Show day</button>
	</form><form action="/events_for_week">
		<button>Show week</button>
	</form><form action="/events_for_month">
		<button>Show month</button>
	</form>
</div>`

const DAY_TEMPLATE_WITH_COMMENT string = `<table><tr>
<th>%02d:%02d - %02d:%02d %s</th>
</tr>
<tr>
<td>%s</td>
</tr></table>
<form action="/update_event_form" method="post">
	<input id="uuid" type="hidden" name="uuid" value="%s">
	<button>Update event</button>
</form>`

const DAY_TEMPLATE_NO_COMMENT string = `<table><tr>
<th>%02d:%02d - %02d:%02d %s</th>
</tr>
</table>
<form action="/update_event_form" method="post">
	<input id="uuid" type="hidden" name="uuid" value="%s">
	<button>Update event</button>
</form>`

const WEEK_TEMPLATE string = `<table>
<tr><th>Monday</th><th>Tuesday</th><th>Wednesday</th><th>Thursday</th><th>Friday</th><th>Saturday</th><th>Sunday</th></tr>
<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>
<tr></tr>
</table>`

const CREATE_EVENT_TEMPLATE string = `<h2>Add new event:</h2><br><form action="/create_event" method="post">
<label for="start-time">Event start time: </label>
<input id="start-time" type="time" name="start-time" value="12:00" /><br>
<label for="end-time">Event end time: </label>
<input id="end-time" type="time" name="end-time" value="12:00" /><br>
<label for="event-name">Event name: </label>
<input id="event-name" type="text" name="event-name" value="New event" /><br>
<label for="start-time">Comment: </label>
<input id="comment" type="text" name="comment"/>
<input type="submit" value="Submit">
</form></html>`

const UPDATE_EVENT_TEMPLATE string = `<h2>Update this event:</h2><br><form action="/update_event" method="post">
<label for="start-time">Event start time: </label>
<input id="start-time" type="time" name="start-time" value="12:00" /><br>
<label for="end-time">Event end time: </label>
<input id="end-time" type="time" name="end-time" value="12:00" /><br>
<label for="event-name">Event name: </label>
<input id="event-name" type="text" name="event-name" value="New event" /><br>
<label for="start-time">Comment: </label>
<input id="comment" type="text" name="comment"/>
<input id="uuid" type="hidden" name="uuid" value="%s">
<input type="submit" value="Submit">
</form></html>`

func (c *calendar) events_for_day(r *http.Request, w http.ResponseWriter) {
	move := r.FormValue("move")
	if move != "" {
		changeDay(c, move)
	} else {
		selectDay(c, r)
	}
	data := TEMPLATE_TOP
	day := &c.years[CURRENT_YEAR-c.firstYear].months[CURRENT_MONTH].days[CURRENT_DAY]
	data += fmt.Sprintf(`<h1 style="text-align: center;">%d/%d/%d - %s</h1>
						<form action="/events_for_day" method="get"><button name="move" value="previous">Previous day</button>
						<button name="move" value="next">Next day</button></form>`, CURRENT_DAY+1, CURRENT_MONTH+1, CURRENT_YEAR, day.weekday.String())
	dayEvents := day.events
	if dayEvents == nil {
		data += `<h2>No events</h2>`
	}
	for dayEvents != nil {
		var new string
		if dayEvents.comment != "" {
			new = fmt.Sprintf(DAY_TEMPLATE_WITH_COMMENT, dayEvents.startTime.Hour(), dayEvents.startTime.Minute(),
				dayEvents.endTime.Hour(), dayEvents.endTime.Minute(), dayEvents.name, dayEvents.comment, dayEvents.uuid)
		} else {
			new = fmt.Sprintf(DAY_TEMPLATE_NO_COMMENT, dayEvents.startTime.Hour(), dayEvents.startTime.Minute(),
				dayEvents.endTime.Hour(), dayEvents.endTime.Minute(), dayEvents.name, dayEvents.uuid)
		}
		data = string(append([]byte(data), []byte(new)...))
		dayEvents = dayEvents.next
	}
	data = string(append([]byte(data), []byte(CREATE_EVENT_TEMPLATE)...))
	tmpl, _ := template.New("response").Parse(data)
	tmpl.Execute(w, nil)
}

func (c *calendar) events_for_week(r *http.Request, w http.ResponseWriter) {
	move := r.FormValue("move")
	if move != "" {
		changeWeek(c, move)
	}
	data := TEMPLATE_TOP
	days := findWeek(c)
	data += fmt.Sprintf(`<h1 style="text-align: center;">Week #%d %s - %d %s %d</h1>
						<form action="/events_for_week" method="get"><button name="move" value="previous">Previous week</button>
						<button name="move" value="next">Next week</button></form><table><tr>`, days[0].t.Day(), days[0].t.Month().String(), days[6].t.Day(), days[6].t.Month().String(), CURRENT_YEAR)
	for i := 0; i < len(days); i++ {
		data += fmt.Sprintf(`<th>%d/%d/%d</th>`, days[i].t.Day(), days[i].t.Month(), days[i].t.Year())
	}
	data += `</tr><tr>`
	for i := 0; i < len(days); i++ {
		data += fmt.Sprintf(`<td>%d events</td>`, getEventLen(days[i].events))
	}
	data += `</tr><tr>`
	for i := 0; i < len(days); i++ {
		data += fmt.Sprintf(`<td><form action="/events_for_day" method="get">
							<input name="current_year" type="hidden" value="%d">
							<input name="current_month" type="hidden" value="%d">
							<input name="current_day" type="hidden" value="%d">
							<input type="submit" value="Open"></form></td>`, days[i].t.Year(), days[i].t.Month()-1, days[i].t.Day()-1)
	}
	data += `</tr></table></html>`
	tmpl, _ := template.New("response").Parse(data)
	tmpl.Execute(w, nil)
}

func (c *calendar) events_for_month(r *http.Request, w http.ResponseWriter) {
	move := r.FormValue("move")
	if move != "" {
		changeMonth(c, move)
	}
	data := TEMPLATE_TOP
	days := c.years[CURRENT_YEAR-c.firstYear].months[CURRENT_MONTH].days
	data += fmt.Sprintf(`<h1 style="text-align: center;">%s %d</h1>
						<form action="/events_for_month" method="get"><button name="move" value="previous">Previous month</button>
						<button name="move" value="next">Next month</button></form><table><tr>`, days[CURRENT_DAY].t.Month().String(), CURRENT_YEAR)
	i := 0
	for i < int(days[0].weekday) {
		data += `<td></td>`
		i++
	}
	for j := range days {
		data += fmt.Sprintf(`<td><form action="/events_for_day" method="get">
							<input name="current_day" type="hidden" value="%d">
							<input type="submit" value="%s %d/%d"></form></td>`, days[j].t.Day()-1, days[j].weekday.String()[:3], days[j].t.Day(), days[j].t.Month())
		i++
		if i > 6 {
			i = 0
			data += `</tr><tr>`
		} else if j == len(days)-1 {
			for i < 7 {
				data += `<td></td>`
				i++
			}
		}
	}
	data += `</tr></table></html>`
	tmpl, _ := template.New("response").Parse(data)
	tmpl.Execute(w, nil)
}

func (c *calendar) create_event(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start-time")
	end := r.FormValue("end-time")
	name := r.FormValue("event-name")
	comment := r.FormValue("comment")
	c.years[CURRENT_YEAR-c.firstYear].months[CURRENT_MONTH].days[CURRENT_DAY].insertEvent(name, comment, start, end)
	c.events_for_day(r, w)
}

func (c *calendar) update_event(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start-time")
	end := r.FormValue("end-time")
	name := r.FormValue("event-name")
	comment := r.FormValue("comment")
	uuid := r.FormValue("uuid")
	c.years[CURRENT_YEAR-c.firstYear].months[CURRENT_MONTH].days[CURRENT_DAY].updateEvent(name, comment, start, end, uuid)
	c.events_for_day(r, w)
}

func (c *calendar) update_event_form(w http.ResponseWriter, r *http.Request) {
	uuid := r.FormValue("uuid")
	data := TEMPLATE_TOP
	data += fmt.Sprintf(`<h1>%d/%d/%d</h1>`, CURRENT_DAY+1, CURRENT_MONTH+1, CURRENT_YEAR)
	dayEvents := c.years[CURRENT_YEAR-c.firstYear].months[CURRENT_MONTH].days[CURRENT_DAY].events
	for dayEvents.uuid != uuid {
		dayEvents = dayEvents.next
	}
	data += fmt.Sprintf(`<table><tr>
		<th>%02d:%02d - %02d:%02d %s</th>
		</tr>
		<tr>
		<td>%s</td>
		</tr></table>`, dayEvents.startTime.Hour(), dayEvents.startTime.Minute(),
		dayEvents.endTime.Hour(), dayEvents.endTime.Minute(), dayEvents.name, dayEvents.comment)
	data += fmt.Sprintf(UPDATE_EVENT_TEMPLATE, uuid)
	tmpl, _ := template.New("response").Parse(data)
	tmpl.Execute(w, nil)
}

func (c *calendar) delete_event(w http.ResponseWriter, formValue string) {

}

func changeDay(cal *calendar, move string) {
	if move == "previous" {
		CURRENT_DAY--
		if CURRENT_DAY < 0 {
			CURRENT_MONTH--
			if CURRENT_MONTH < 0 {
				CURRENT_MONTH = 11
				CURRENT_YEAR--
			}
			CURRENT_DAY = len(cal.years[CURRENT_YEAR-cal.firstYear].months[CURRENT_MONTH].days) - 1
		}
	} else if move == "next" {
		CURRENT_DAY++
		if CURRENT_DAY >= len(cal.years[CURRENT_YEAR-cal.firstYear].months[CURRENT_MONTH].days) {
			CURRENT_DAY = 1
			CURRENT_MONTH++
			if CURRENT_MONTH > 11 {
				CURRENT_MONTH = 0
				CURRENT_YEAR++
			}
		}
	}
}

func changeWeek(cal *calendar, move string) {
	if move == "previous" {
		CURRENT_DAY -= 7
		if CURRENT_DAY < 0 {
			CURRENT_MONTH--
			if CURRENT_MONTH < 0 {
				CURRENT_MONTH = 11
				CURRENT_YEAR--
			}
			CURRENT_DAY = len(cal.years[CURRENT_YEAR-cal.firstYear].months[CURRENT_MONTH].days) - 1 + CURRENT_DAY
		}
	} else if move == "next" {
		CURRENT_DAY += 7
		if CURRENT_DAY >= len(cal.years[CURRENT_YEAR-cal.firstYear].months[CURRENT_MONTH].days) {
			CURRENT_DAY %= len(cal.years[CURRENT_YEAR-cal.firstYear].months[CURRENT_MONTH].days)
			CURRENT_MONTH++
			if CURRENT_MONTH > 11 {
				CURRENT_MONTH = 0
				CURRENT_YEAR++
			}
		}
	}
}

func changeMonth(cal *calendar, move string) {
	if move == "previous" {
		CURRENT_MONTH--
		if CURRENT_MONTH < 0 {
			CURRENT_MONTH = 11
			CURRENT_YEAR--
		}
	} else if move == "next" {
		CURRENT_MONTH++
		if CURRENT_MONTH > 11 {
			CURRENT_MONTH = 0
			CURRENT_YEAR++
		}
	}
}

func findWeek(c *calendar) []day {
	days := make([]day, 7)
	currentDay := &c.years[CURRENT_YEAR-c.firstYear].months[CURRENT_MONTH].days[CURRENT_DAY]
	for i := 0; i < 7; i++ {
		days[i] = getDay(c, i-int(currentDay.weekday))
	}
	return days
}

func getDay(cal *calendar, dir int) day {
	cd := CURRENT_DAY
	cm := CURRENT_MONTH
	cy := CURRENT_YEAR
	cd += dir
	if cd < 0 {
		cm--
		if cm < 0 {
			cm = 11
			cy--
		}
		cd = len(cal.years[cy-cal.firstYear].months[cm].days) - 1 + cd
	} else if cd >= len(cal.years[cy-cal.firstYear].months[cm].days) {
		cd %= len(cal.years[cy-cal.firstYear].months[cm].days)
		cm++
		if cm > 11 {
			cm = 0
			cy++
		}
	}
	return cal.years[cy-cal.firstYear].months[cm].days[cd]
}

func getEventLen(e *event) (len int) {
	for e != nil {
		len++
		e = e.next
	}
	return
}

func selectDay(c *calendar, r *http.Request) {
	y := r.FormValue("current_year")
	m := r.FormValue("current_month")
	d := r.FormValue("current_day")
	if y == "" || m == "" || d == "" {
		return
	}
	year, _ := strconv.Atoi(y)
	month, _ := strconv.Atoi(m)
	day, _ := strconv.Atoi(d)
	CURRENT_YEAR = uint16(year)
	CURRENT_MONTH = month
	CURRENT_DAY = day
}
