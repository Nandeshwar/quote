package event

import (
	"fmt"
	"strings"
	"time"

	"github.com/gookit/color"
)

//go:generate gofp -destination fp.go -pkg event -type "EventDetail" -mapfun "true"
type EventDetail struct {
	Day          int
	Month        int
	Year         int
	Title        string
	Info         string
	URL          string
	Type         string // Value can be same|different: different - event occurs different day in each year
	CreationDate time.Time
}

func AllEvents() []*EventDetail {
	var allEvents []*EventDetail
	allEvents = append(allEvents, KripaluJiMaharajEvents()...)
	allEvents = append(allEvents, PrabhuyEvents()...)
	allEvents = append(allEvents, MixEvents()...)
	allEvents = append(allEvents, copyBikramSambatEventsToEventDetail()...)
	return allEvents
}

func TodayEvents() []*EventDetail {
	year, month, day := time.Now().Date()

	findTodayEvent := func(event *EventDetail) bool {
		if event.Type == "different" {
			if event.Year == year && event.Month == int(month) && event.Day == day {
				return true
			}
		} else if event.Month == int(month) && event.Day == day {
			return true
		}
		return false
	}

	todayEvents := FilterEventDetailPtr(findTodayEvent, AllEvents())

	return todayEvents
}

func (e EventDetail) DisplayEvent() {
	blue := color.FgBlue.Render
	fmt.Println(e.Title)
	fmt.Println(e.Info)
	fmt.Println(e.URL)
	fmt.Printf("%d-%d-%d\n", e.Year, e.Month, e.Day)
	for i, url := range strings.Split(e.URL, ";") {
		fmt.Printf("\n%d. %s ", i+1, blue(url))
	}
}
