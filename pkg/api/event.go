package api

import (
	"fmt"
	"net/http"
	"quote/pkg/constants"
	"quote/pkg/model"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) events(w http.ResponseWriter, r *http.Request) {
	searchText := mux.Vars(r)["searchText"]

	filteredEvents := s.findEvents(searchText)
	displayEvents(filteredEvents, w)
}

func (s *Server) findEvents(searchText string) []model.EventDetail {
	eventDetailList, err := s.eventDetailService.GetEventDetailByTitleOrInfo(searchText)
	if err != nil {
		fmt.Println("error =", err)
	}

	return eventDetailList
}

func displayEvents(filteredEvents []model.EventDetail, w http.ResponseWriter) {
	fmt.Fprintf(w, "<h1>Events:</h1>")

	fmt.Fprintf(w, fmt.Sprintf("<table border='2'>"))

	fmt.Fprintf(w, fmt.Sprintf("<tr>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>SN</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>ID</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Title</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Info</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Links</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Date</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Creattion Date</th>"))
	fmt.Fprintf(w, fmt.Sprintf("</tr>"))
	for i, event := range filteredEvents {
		fmt.Fprintf(w, fmt.Sprintf("<tr>"))
		fmt.Fprintf(w, fmt.Sprintf("<td>%d.</td>", i+1))
		fmt.Fprintf(w, fmt.Sprintf("<td>%d</td>", event.ID))
		fmt.Fprintf(w, fmt.Sprintf("<td>%s</td>", event.Title))

		// Split info by .
		fmt.Fprintf(w, fmt.Sprintf("<td>"))
		fmt.Fprintf(w, fmt.Sprintf("<table>"))
		for _, info := range strings.Split(event.Info, ".") {
			fmt.Fprintf(w, fmt.Sprintf("<tr>"))
			fmt.Fprintf(w, fmt.Sprintf("<td>%s</td>", info))
			fmt.Fprintf(w, fmt.Sprintf("</tr>"))
		}
		fmt.Fprintf(w, fmt.Sprintf("</table>"))
		fmt.Fprintf(w, fmt.Sprintf("</td>"))

		// Display URL in different table under td
		fmt.Fprintf(w, fmt.Sprintf("<td>"))
		fmt.Fprintf(w, fmt.Sprintf("<table>"))

		for i, url := range event.Links {
			var youtubeLink string
			if strings.Contains(strings.ToLower(url), "youtube") {
				youtubeLink = "click me to watch on youtube"
			}
			if !strings.Contains(strings.ToLower(url), "no-link") {
				fmt.Fprintf(w, fmt.Sprintf("<tr><td><a href='%s'>Links%d. %s </a></td></tr>", url, i+1, youtubeLink))
			}
		}
		fmt.Fprintf(w, fmt.Sprintf("</td>"))
		fmt.Fprintf(w, fmt.Sprintf("</table>"))

		eventDate := time.Date(event.Year, time.Month(event.Month), event.Day, 0, 0, 0, 0, time.Local)

		fmt.Fprintf(w, fmt.Sprintf("<td>%s</td>", eventDate.Format(constants.DATE_FORMAT_EVENT_DATE_DISPLAY)))
		fmt.Fprintf(w, fmt.Sprintf("<td>%v</td>", event.CreationDate.Format(constants.DATE_FORMAT_INFO)))

		fmt.Fprintf(w, fmt.Sprintf("</tr>"))
	}
	fmt.Fprintf(w, fmt.Sprintf("</table>"))
}
