package main

import (
	"time"

	"github.com/google/uuid"
)

func (d *day) insertEvent(name, comment, start, end string) {
	var new event
	new.uuid = uuid.NewString()
	new.startTime, _ = time.Parse("15:04", start)
	new.endTime, _ = time.Parse("15:04", end)
	new.name = name
	new.comment = comment
	if d.events == nil {
		d.events = &new
	} else if new.startTime.Before(d.events.startTime) {
		new.next = d.events
		d.events = &new
	} else {
		p := d.events
		for p.next != nil && !new.startTime.Before(p.next.startTime) {
			p = p.next
		}
		new.next = p.next
		p.next = &new
	}
}

func (d *day) updateEvent(name, comment, start, end, uuid string) {
	p := d.events
	for p != nil && p.uuid != uuid {
		p = p.next
	}
	if p == nil {
		return
	}
	p.name = name
	p.comment = comment
	p.startTime, _ = time.Parse("15:04", start)
	p.endTime, _ = time.Parse("15:04", end)
}
