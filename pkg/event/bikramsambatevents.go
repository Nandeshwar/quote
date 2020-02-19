package event

import "time"

type EventDate struct {
	Day   int
	Month int
	Year  int
}
type Events struct {
	EventDates   []EventDate
	Title        string
	Info         string
	URL          string
	CreationDate time.Time
}

func allImportantEvents() []Events {
	return []Events{
		{
			EventDates: []EventDate{
				{Day: 11, Month: 8, Year: 2020},
				{Day: 30, Month: 8, Year: 2021},
				{Day: 18, Month: 8, Year: 2022},
				{Day: 6, Month: 9, Year: 2023},
				{Day: 26, Month: 8, Year: 2024},
				{Day: 15, Month: 8, Year: 2025},
			},
			Title:        "Krishna appearance day",
			Info:         "Krishna Janmasthmi",
			URL:          "https://www.calendardate.com/krishna_janmashtami_2025.htm",
			CreationDate: time.Date(2020, 2, 16, 0, 0, 0, 0, time.Local),
		},

		{
			EventDates: []EventDate{
				{Day: 26, Month: 8, Year: 2020},
				{Day: 13, Month: 9, Year: 2021},
				{Day: 3, Month: 9, Year: 2022},
				{Day: 22, Month: 9, Year: 2023},
				{Day: 11, Month: 8, Year: 2024},
				{Day: 31, Month: 8, Year: 2025},
			},
			Title:        "Radha ji appearance day",
			Info:         "Radha Janmasthmi",
			URL:          "https://www.calendardate.com/krishna_janmashtami_2025.htm",
			CreationDate: time.Date(2020, 2, 16, 0, 0, 0, 0, time.Local),
		},

		{
			EventDates: []EventDate{
				{Day: 02, Month: 4, Year: 2020},
				{Day: 21, Month: 4, Year: 2021},
				{Day: 10, Month: 4, Year: 2022},
				{Day: 30, Month: 3, Year: 2023},
				{Day: 16, Month: 4, Year: 2024},
				{Day: 5, Month: 4, Year: 2025},
			},
			Title:        "Shree Ram Ji appearance day",
			Info:         "Ram, Bharat, Laxman, Satrughan Appearance day",
			URL:          "https://www.calendardate.com/rama_navami_2020.htm",
			CreationDate: time.Date(2020, 2, 19, 8, 40, 0, 0, time.Local),
		},
	}
}

func copyBikramSambatEventsToEventDetail() []*EventDetail {
	var events []*EventDetail

	for _, bEvent := range allImportantEvents() {
		for _, eventDate := range bEvent.EventDates {
			eventDetail := &EventDetail{
				Day:          eventDate.Day,
				Month:        eventDate.Month,
				Year:         eventDate.Year,
				Title:        bEvent.Title,
				Info:         bEvent.Info,
				URL:          bEvent.URL,
				CreationDate: bEvent.CreationDate,
			}
			events = append(events, eventDetail)
		}
	}

	return events
}
