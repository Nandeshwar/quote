package api

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"quote/pkg/constants"
	"quote/pkg/model"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func (s *Server) events(w http.ResponseWriter, r *http.Request) {
	searchText := mux.Vars(r)["searchText"]

	filteredEvents, err := s.findEvents(searchText)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"searchText": searchText,
			"error":      err,
		}).Errorf("error searching events")
		w.WriteHeader(http.StatusInternalServerError)
	}
	displayEvents(filteredEvents, w)
	w.WriteHeader(http.StatusOK)
}

func (s *Server) findEvents(searchText string) ([]model.EventDetail, error) {
	eventDetailList, err := s.eventDetailService.GetEventDetailByTitleOrInfo(searchText)
	if err != nil {
		return nil, err
	}

	return eventDetailList, nil
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

// swagger:operation GET /api/quote/v1/eventsByMonth/{month} EventDetail eventDetail
// ---
// description: get events by month
// consumes:
// - "application/json"
// parameters:
// - name: month
//   description: get events by month
//   in: path
//   required: true
//   default: Jan
//   type: string
// Responses:
//   '200':
//     description: Ok
//   '400':
//     description: Bad request
//   '404':
//     description: Not found
//   '500':
//     description: Internal server error
func (s *Server) getEventByMonth(w http.ResponseWriter, r *http.Request) {
	month := mux.Vars(r)["month"]

	monthInt, err := strconv.Atoi(month)
	if err != nil {
		switch strings.ToLower(month) {
		case "jan", "january":
			monthInt = 1

		case "feb", "february":
			monthInt = 2

		case "mar", "march":

			monthInt = 3
		case "apr", "april":
			monthInt = 4

		case "may":
			monthInt = 5

		case "jun":
			monthInt = 6

		case "jul", "july":
			monthInt = 7

		case "aug", "august":
			monthInt = 8

		case "oct", "october":
			monthInt = 10

		case "nov", "november":
			monthInt = 11

		case "dec", "december":
			monthInt = 12
		default:
			monthInt = 0
		}
	}
	if monthInt == 0 {
		logrus.WithFields(logrus.Fields{
			"month": month,
			"error": "invalid month",
		}).Errorf("expected month either 1-12 or Jan-Dec or January-December")
		w.WriteHeader(http.StatusBadRequest)
	}

	events, err := s.eventDetailService.GetEventDetailByMonth(monthInt)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"month": month,
			"error": err,
		}).Errorf("error fetching records for the month")
		w.WriteHeader(http.StatusInternalServerError)
	}

	jsonBytes, err := json.Marshal(events)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"month": month,
			"error": err,
		}).Errorf("error converting to json")
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
