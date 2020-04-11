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

		{
			EventDates: []EventDate{
				{Day: 02, Month: 5, Year: 2020},
				{Day: 21, Month: 5, Year: 2021},
				{Day: 10, Month: 5, Year: 2022},
				{Day: 29, Month: 4, Year: 2023},
				{Day: 17, Month: 5, Year: 2024},
				{Day: 6, Month: 5, Year: 2025},
			},
			Title:        "Sita Ji appearance day",
			Info:         "Sita Navami, Sita Jayanti",
			URL:          "https://www.prokerala.com/festivals/sita-navami.html",
			CreationDate: time.Date(2020, 3, 2, 14, 13, 0, 0, time.Local),
		},

		{
			EventDates: []EventDate{
				{Day: 18, Month: 12, Year: 2020},
				{Day: 8, Month: 12, Year: 2021},
				{Day: 28, Month: 11, Year: 2022},
				{Day: 16, Month: 12, Year: 2023},
				{Day: 5, Month: 12, Year: 2024},
				{Day: 25, Month: 12, Year: 2025},
			},
			Title:        "Sita Ram Vivah",
			Info:         "Vivah Panchami",
			URL:          "https://www.drikpanchang.com/festivals/vivah-panchami/vivah-panchami-date-time.html?year=2021",
			CreationDate: time.Date(2020, 3, 6, 15, 19, 0, 0, time.Local),
		},

		{
			EventDates: []EventDate{
				{Day: 8, Month: 4, Year: 2020},
				{Day: 27, Month: 4, Year: 2021},
				{Day: 16, Month: 4, Year: 2022},
				{Day: 6, Month: 4, Year: 2023},
				{Day: 23, Month: 4, Year: 2024},
				{Day: 12, Month: 4, Year: 2025},
			},
			Title:        "Hanuman Jayanti",
			Info:         "Hanuman ji appearance day",
			URL:          "https://www.drikpanchang.com/festivals/hanuman-jayanti/hanuman-jayanti-date-time.html?year=2021",
			CreationDate: time.Date(2020, 3, 6, 15, 19, 0, 0, time.Local),
		},

		{
			EventDates: []EventDate{
				{Day: 21, Month: 2, Year: 2020},
				{Day: 11, Month: 3, Year: 2021},
				{Day: 28, Month: 2, Year: 2022},
				{Day: 18, Month: 2, Year: 2023},
				{Day: 8, Month: 3, Year: 2024},
				{Day: 25, Month: 25, Year: 2025},
			},
			Title:        "Maha Shiva Ratri",
			Info:         "Maha Shivaratri is a Hindu festival celebrated annually in honour of Lord Shiva. The name also refers to the night when Shiva performs the heavenly dance",
			URL:          "https://www.calendardate.com/maha_shivaratri_2020.htm;https://en.wikipedia.org/wiki/Maha_Shivaratri",
			CreationDate: time.Date(2020, 4, 11, 13, 42, 0, 0, time.Local),
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
				Type:         "different",
				CreationDate: bEvent.CreationDate,
			}
			events = append(events, eventDetail)
		}
	}

	return events
}
