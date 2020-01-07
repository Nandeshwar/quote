package event

import (
	"fmt"
	"time"

	"github.com/gookit/color"
)

type EventDetail struct {
	Day   int
	Month int
	Year  int
	Title string
	Info  string
	URL   string
}

func allEvents() []EventDetail {
	var allEvents []EventDetail

	kripaluJiJagadGurutam := EventDetail{
		Day:   14,
		Month: 1,
		Year:  1957,
		Title: "Kripalu Ji Maharaj - Jagadgurutam title day",
		Info: `    He was formally installed as the fifth Jagadguru (world teacher).
   He was 34 years old when given the title on 14 January 1957 by the Kashi Vidvat Parishat, a group of Hindu scholars.
   The Kashi Vidvat Parishat conferred on him the titles Bhaktiyog-Ras-Avtar and Jagadguruttama.
   Followers claim that he is the "fifth original Jagadguru" in the series of Jagadgurus after 
   Śrīpāda Śaṅkarācārya (A.D. 788-820),
   Śrīpāda Rāmānujācārya (1017-1137),
   Śrī Nimbārkācārya and, 
   Śrīpāda Madhvācārya (1239-1319)`,
		URL: "https://en.wikipedia.org/wiki/Kripalu_Maharaj",
	}
	allEvents = append(allEvents, kripaluJiJagadGurutam)

	kripaluJiAppearance := EventDetail{
		Day:   5,
		Month: 10,
		Year:  1922,
		Title: "Kripalu Ji Maharaj - appearance day",
		Info: `    Kripalu (Sanskrit: जगद्गुरु श्री कृपालु जी महाराज, IAST: Kṛpālu) (5 October 1922 – 15 November 2013) 
    He was a Hindu spiritual leader  from Allahabad (Prayag) - Mangarh, Pratapgarh, India.
    He was the founder of Jagadguru Kripalu Parishat (JKP), 
    a worldwide Hindu non-profit organization with five main ashrams; four in India and one in the United States.
    JKP Radha Madhav Dham is one of the largest Hindu Temple complexes in the Western Hemisphere, and the largest in North America.
     `,
		URL: "https://en.wikipedia.org/wiki/Kripalu_Maharaj",
	}
	allEvents = append(allEvents, kripaluJiAppearance)

	return allEvents
}

func TodayEvents() []EventDetail {
	_, month, day := time.Now().Date()

	var todayEvents []EventDetail

	for _, event := range allEvents() {
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
	fmt.Printf("\n%s\n ", blue(e.URL))
}
