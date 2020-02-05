package api

import (
	"fmt"
	"net/http"
	"quote/pkg/event"
	"strings"
)

func findEvents(searchText string) []*event.EventDetail {
	allEvents := event.AllEvents()

	var filteredEvents []*event.EventDetail

	filterBySearch := func(event *event.EventDetail) bool {

		if strings.Contains(strings.ToLower(event.Info), searchText) ||
			strings.Contains(strings.ToLower(event.Title), searchText) ||
			strings.Contains(strings.ToLower(event.URL), searchText) {
			return true
		}
		return false
	}

	if searchText != "" {
		searchText = strings.ToLower(searchText)
		filteredEvents = event.FilterEventDetailPtr(filterBySearch, allEvents)
	} else {
		filteredEvents = allEvents
	}
	return filteredEvents
}

func displayEvents(filteredEvents []*event.EventDetail, w http.ResponseWriter) {
	fmt.Fprintf(w, "<h1>Events:</h1>")

	fmt.Fprintf(w, fmt.Sprintf("<table border='2'>"))

	fmt.Fprintf(w, fmt.Sprintf("<tr>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Number</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Title</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Info</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Link</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Date</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Creattion Date</th>"))
	fmt.Fprintf(w, fmt.Sprintf("</tr>"))
	for i, event := range filteredEvents {
		fmt.Fprintf(w, fmt.Sprintf("<tr>"))
		fmt.Fprintf(w, fmt.Sprintf("<td>%d.</td>", i+1))
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
		for i, url := range strings.Split(event.URL, ";") {
			fmt.Fprintf(w, fmt.Sprintf("<tr><td><a href='%s'>Link%d </a></td></tr>", url, i+1))
		}
		fmt.Fprintf(w, fmt.Sprintf("</td>"))
		fmt.Fprintf(w, fmt.Sprintf("</table>"))

		fmt.Fprintf(w, fmt.Sprintf("<td>%d-%d-%d</td>", event.Year, event.Month, event.Day))
		fmt.Fprintf(w, fmt.Sprintf("<td>%v</td>", event.CreationDate))

		fmt.Fprintf(w, fmt.Sprintf("</tr>"))
		fmt.Fprintf(w, fmt.Sprintf("</br>"))

	}
	fmt.Fprintf(w, fmt.Sprintf("</table>"))
}
