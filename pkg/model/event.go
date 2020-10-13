package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/gookit/color"
)

//go:generate gofp -destination fpEventDetail.go -pkg model -type "EventDetail"
type EventDetail struct {
	ID           int64
	Day          int
	Month        int
	Year         int
	Title        string
	Info         string
	URL          string
	Links        []string
	Type         string // Value can be same|different: different - event occurs different day in each year. Example Ram birthday is different each year as per Bikram Sambat
	CreationDate time.Time
	UpdatedAt    time.Time
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
