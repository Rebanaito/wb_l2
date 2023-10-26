package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

const YEARS_MAX int = 100
const DISPLAY_DAY uint16 = 1
const DISPLAY_WEEK uint16 = 2
const DISPLAY_MONTH uint16 = 3

var CURRENT_DISPLAY int = int(DISPLAY_MONTH)
var CURRENT_YEAR uint16 = uint16(time.Now().Year())
var CURRENT_MONTH int = int(time.Now().Month()) - 1
var CURRENT_DAY int = time.Now().Day() - 1

type event struct {
	uuid      string
	name      string
	comment   string
	startTime time.Time
	endTime   time.Time
	next      *event
}

type day struct {
	weekday time.Weekday
	events  *event
}

type month struct {
	name string
	days []day
}

type year struct {
	yearNum uint16
	months  []month
}

type calendar struct {
	firstYear uint16
	years     []year
}

func main() {
	calendar, noEnv := initEnv()
	if noEnv {
		calendar = initYears()
	}
	runServer(calendar)
}

func runServer(calendar calendar) {
	mu := &sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, calendar, mu)
	})

	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	kill := make(chan os.Signal, 1)
	signal.Notify(kill, os.Interrupt)
	<-kill

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}

func handler(w http.ResponseWriter, r *http.Request, calendar calendar, mu *sync.Mutex) {
	if r.Method == "GET" {
		switch r.URL.Path {
		case "/events_for_day":
			calendar.events_for_day(r, w)
		case "/events_for_week":
			calendar.events_for_week(r, w)
		case "/", "/events_for_month":
			calendar.events_for_month(r, w)
		}
	} else if r.Method == "POST" {
		switch r.URL.Path {
		case "/create_event":
			calendar.create_event(w, r)
		case "/update_event":
			calendar.update_event(w, r)
		case "/delete_event":
			calendar.delete_event(w, r.FormValue("event_uid"))
		case "/update_event_form":
			calendar.update_event_form(w, r)
		}
	}
}
