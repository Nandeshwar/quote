package api

import (
	"fmt"
	"net/http"
	info2 "quote/pkg/info"
	"strings"
)

func findInfo(searchText string) []info2.Info {
	allInfo := info2.GetAllInfo()

	var filteredInfo []info2.Info

	filterBySearch := func(info info2.Info) bool {

		if strings.Contains(strings.ToLower(info.Info), searchText) ||
			strings.Contains(strings.ToLower(info.Title), searchText) {
			return true
		}
		return false
	}

	if searchText != "" {
		searchText = strings.ToLower(searchText)
		filteredInfo = info2.Filter(filterBySearch, allInfo)
	} else {
		filteredInfo = allInfo
	}
	return filteredInfo
}

func displayInfo(filteredInfo []info2.Info, w http.ResponseWriter) {
	fmt.Fprintf(w, "<title>Info</title>")

	fmt.Fprintf(w, fmt.Sprintf("<table border='2'>"))

	fmt.Fprintf(w, fmt.Sprintf("<tr>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Event Number</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Title</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Info</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Link</th>"))
	fmt.Fprintf(w, fmt.Sprintf("<th>Info Added</th>"))
	fmt.Fprintf(w, fmt.Sprintf("</tr>"))
	for i, info := range filteredInfo {
		fmt.Fprintf(w, fmt.Sprintf("<tr>"))
		fmt.Fprintf(w, fmt.Sprintf("<td>%d.</td>", i+1))
		fmt.Fprintf(w, fmt.Sprintf("<td>%s</td>", info.Title))

		// Split info by .
		fmt.Fprintf(w, fmt.Sprintf("<td>"))
		fmt.Fprintf(w, fmt.Sprintf("<table>"))
		for _, info := range strings.Split(info.Info, ".") {
			fmt.Fprintf(w, fmt.Sprintf("<tr>"))
			fmt.Fprintf(w, fmt.Sprintf("<td>%s</td>", info))
			fmt.Fprintf(w, fmt.Sprintf("</tr>"))
		}
		fmt.Fprintf(w, fmt.Sprintf("</table>"))
		fmt.Fprintf(w, fmt.Sprintf("</td>"))

		// Display URL in different table under td
		fmt.Fprintf(w, fmt.Sprintf("<td>"))
		fmt.Fprintf(w, fmt.Sprintf("<table>"))
		for i, url := range info.Link {
			fmt.Fprintf(w, fmt.Sprintf("<tr><td><a href='%s'>Link%d </a></td></tr>", url, i+1))
		}
		fmt.Fprintf(w, fmt.Sprintf("</td>"))
		fmt.Fprintf(w, fmt.Sprintf("</table>"))

		fmt.Fprintf(w, fmt.Sprintf("<td>%v</td>", info.CreationDate))
		fmt.Fprintf(w, fmt.Sprintf("</tr>"))
		fmt.Fprintf(w, fmt.Sprintf("</br>"))

	}
	fmt.Fprintf(w, fmt.Sprintf("</table>"))
}
