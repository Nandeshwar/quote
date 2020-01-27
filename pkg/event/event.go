package event

import (
	"fmt"
	"strings"
	"time"

	"github.com/gookit/color"
)

//go:generate gofp -destination fp.go -pkg event -type "EventDetail" -map-function "true"
type EventDetail struct {
	Day          int
	Month        int
	Year         int
	Title        string
	Info         string
	URL          string
	CreationDate time.Time
}

func AllEvents() []*EventDetail {
	var allEvents []*EventDetail
	allEvents = append(allEvents, KripaluJiMaharajEvents()...)
	allEvents = append(allEvents, PrabhuyEvents()...)
	allEvents = append(allEvents, MixEvents()...)
	return allEvents
}

func TodayEvents() []*EventDetail {
	_, month, day := time.Now().Date()

	var todayEvents []*EventDetail

	for _, event := range AllEvents() {
		if event.Month == int(month) && event.Day == day {
			todayEvents = append(todayEvents, event)
		}
	}
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
