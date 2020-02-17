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
			CreationDate: time.Date(2020, 2, 16, 0, 0, 0, 0, time.Local),
		},
	}
}

func copyBikramSambatEventsToEventDetail() []*EventDetail {
	var events []*EventDetail

	for _, bEvent := range allImportantEvents() {
		for i := 0; i < len(bEvent.EventDates); i++ {
			eventDetail := &EventDetail{
				Day:          bEvent.EventDates[i].Day,
				Month:        bEvent.EventDates[i].Month,
				Year:         bEvent.EventDates[i].Year,
				Title:        bEvent.Title,
				Info:         bEvent.Info,
				CreationDate: bEvent.CreationDate,
			}
			events = append(events, eventDetail)
		}
	}

	return events
}
