package main

import (
	"fmt"
	"html/template"
	"net/http"
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
}

func (c *calendar) events_for_month(r *http.Request, w http.ResponseWriter) {

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
				CURRENT_MONTH = 1
				CURRENT_YEAR++
			}
		}
	}
}
