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
</tr></table>`

const DAY_TEMPLATE_NO_COMMENT string = `<table><tr>
<th>%02d:%02d - %02d:%02d %s</th>
</tr>
</table>`

const WEEK_TEMPLATE string = `<table>
<tr><th>Monday</th><th>Tuesday</th><th>Wednesday</th><th>Thursday</th><th>Friday</th><th>Saturday</th><th>Sunday</th></tr>
<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>
<tr></tr>
</table>`

const CREATE_EVENT_TEMPLATE string = `<form action="/create_event" method="post">
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

func (c calendar) events_for_day(w http.ResponseWriter) {
	data := TEMPLATE_TOP
	data += fmt.Sprintf(`<h1>%d/%d/%d</h1>`, CURRENT_DAY+1, CURRENT_MONTH+1, CURRENT_YEAR)
	dayEvents := c.years[CURRENT_YEAR-c.firstYear].months[CURRENT_MONTH].days[CURRENT_DAY].events
	if dayEvents == nil {
		data += `<h2>No events</h2>`
	}
	for dayEvents != nil {
		var new string
		if dayEvents.comment != "" {
			new = fmt.Sprintf(DAY_TEMPLATE_WITH_COMMENT, dayEvents.startTime.Hour(), dayEvents.startTime.Minute(),
				dayEvents.endTime.Hour(), dayEvents.endTime.Minute(), dayEvents.name, dayEvents.comment)
		} else {
			new = fmt.Sprintf(DAY_TEMPLATE_NO_COMMENT, dayEvents.startTime.Hour(), dayEvents.startTime.Minute(),
				dayEvents.endTime.Hour(), dayEvents.endTime.Minute(), dayEvents.name)
		}
		data = string(append([]byte(data), []byte(new)...))
		dayEvents = dayEvents.next
	}
	data = string(append([]byte(data), []byte(CREATE_EVENT_TEMPLATE)...))
	tmpl, _ := template.New("response").Parse(data)
	tmpl.Execute(w, nil)
}

func (c calendar) events_for_week(w http.ResponseWriter) {
}

func (c calendar) events_for_month(w http.ResponseWriter) {

}

func (c calendar) create_event(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start-time")
	end := r.FormValue("end-time")
	name := r.FormValue("event-name")
	comment := r.FormValue("comment")
	c.years[CURRENT_YEAR-c.firstYear].months[CURRENT_MONTH].days[CURRENT_DAY].insertEvent(name, comment, start, end)
	c.events_for_day(w)
}

func (c calendar) update_event(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start-time")
	end := r.FormValue("end-time")
	name := r.FormValue("event-name")
	comment := r.FormValue("comment")
	uuid := r.FormValue("uuid")
	c.years[CURRENT_YEAR-c.firstYear].months[CURRENT_MONTH].days[CURRENT_DAY].updateEvent(name, comment, start, end, uuid)
	c.events_for_day(w)
}

func (c calendar) delete_event(w http.ResponseWriter, formValue string) {

}
